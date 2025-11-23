package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SendTestCall places a test call using campaign settings
// This mimics the VICIdial admin.php test_call pattern (lines 17173-17293)
func (h *Handler) SendTestCall(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CampaignID  string `json:"campaign_id"`  // Campaign ID (required)
		PhoneNumber string `json:"phone_number"` // Phone number to dial (required)
		PhoneCode   string `json:"phone_code"`   // Phone code/country code (optional, default: "1")
		User        string `json:"user"`         // User placing the call (optional, default: "API")
		VdadExten   string `json:"vdad_exten"`   // Optional explicit VDAD/routing extension (overrides campaign setting)
		ServerIP    string `json:"server_ip"`    // Optional explicit server_ip (overrides default active server selection)
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate required fields
	if req.CampaignID == "" {
		respondWithError(w, http.StatusBadRequest, "campaign_id is required")
		return
	}
	if req.PhoneNumber == "" {
		respondWithError(w, http.StatusBadRequest, "phone_number is required")
		return
	}
	if len(req.PhoneNumber) < 6 {
		respondWithError(w, http.StatusBadRequest, "phone_number too small, must be at least 6 digits")
		return
	}

	// Set defaults
	if req.PhoneCode == "" {
		req.PhoneCode = "1"
	}
	if req.User == "" {
		req.User = "API"
	}

	now := time.Now()
	sqlDate := now.Format("2006-01-02 15:04:05")
	cidDate := now.Format("010215") // mddhhmmss format for caller ID

	// Get campaign settings
	campaignQuery := `
		SELECT campaign_name, dial_prefix, campaign_cid, dial_timeout,
		       omit_phone_code, campaign_vdad_exten, manual_dial_list_id,
		       ext_context, active, user_group
		FROM vicidial_campaigns
		WHERE campaign_id = ?
	`
	var campaignName, dialPrefix, campaignCID, extContext, active, userGroup string
	var omitPhoneCode, campaignVdadExten sql.NullString
	var dialTimeout, manualDialListID int

	err := h.DB.QueryRow(campaignQuery, req.CampaignID).Scan(
		&campaignName, &dialPrefix, &campaignCID, &dialTimeout,
		&omitPhoneCode, &campaignVdadExten, &manualDialListID,
		&extContext, &active, &userGroup)

	if err != nil {
		if strings.Contains(err.Error(), "Unknown column 'ext_context'") {
			// Older schemas may not have ext_context; fallback without it
			fallbackQuery := `
				SELECT campaign_name, dial_prefix, campaign_cid, dial_timeout,
				       omit_phone_code, campaign_vdad_exten, manual_dial_list_id,
				       active, user_group
				FROM vicidial_campaigns
				WHERE campaign_id = ?
			`
			err = h.DB.QueryRow(fallbackQuery, req.CampaignID).Scan(
				&campaignName, &dialPrefix, &campaignCID, &dialTimeout,
				&omitPhoneCode, &campaignVdadExten, &manualDialListID,
				&active, &userGroup)
			extContext = ""
		}
	}

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Campaign not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve campaign: "+err.Error())
		return
	}

	// Get server details (prefer ext_context + answer_transfer_agent)
	var (
		serverIP            string
		serverExtContext    sql.NullString
		answerTransferAgent sql.NullString
		asteriskVersion     string
		routingPrefix       string
		serverID            string
	)

	req.ServerIP = strings.TrimSpace(req.ServerIP)
	serverInfoQuery := `
		SELECT server_ip, ext_context, answer_transfer_agent, asterisk_version, routing_prefix, server_id
		FROM servers
		WHERE %s
		LIMIT 1
	`

	serverWhere := "active = 'Y'"
	args := []interface{}{}
	if req.ServerIP != "" {
		serverWhere = "server_ip = ?"
		args = append(args, req.ServerIP)
	}

	err = h.DB.QueryRow(fmt.Sprintf(serverInfoQuery, serverWhere), args...).Scan(&serverIP, &serverExtContext, &answerTransferAgent, &asteriskVersion, &routingPrefix, &serverID)
	if err != nil {
		// Fallback for schemas without ext_context/answer_transfer_agent
		fallbackQuery := `
			SELECT server_ip, asterisk_version, routing_prefix, server_id
			FROM servers
			WHERE %s
			LIMIT 1
		`
		err = h.DB.QueryRow(fmt.Sprintf(fallbackQuery, serverWhere), args...).Scan(&serverIP, &asteriskVersion, &routingPrefix, &serverID)
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve server: "+err.Error())
		return
	}

	// Insert test lead into vicidial_list
	leadQuery := `
		INSERT INTO vicidial_list SET
		phone_code=?, phone_number=?, list_id=?, status='CTCALL',
		user='VDAD', called_since_last_reset='Y', entry_date=?,
		last_local_call_time=?, called_count='1', first_name='Test',
		last_name='Call', address1='Test Call', address2='2', address3='3',
		city='Springfield', state='IL', vendor_lead_code='999999',
		comments=?, rank='99', owner='Test Outbound Call'
	`
	comments := fmt.Sprintf("%s test call placed %s", req.CampaignID, sqlDate)

	result, err := h.DB.Exec(leadQuery, req.PhoneCode, req.PhoneNumber,
		manualDialListID, sqlDate, sqlDate, comments)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create test lead: "+err.Error())
		return
	}

	leadID, _ := result.LastInsertId()

	// Get list CID overrides
	var campaignCIDOverride, cidGroupID sql.NullString
	listQuery := "SELECT campaign_cid_override, cid_group_id FROM vicidial_lists WHERE list_id = ?"
	h.DB.QueryRow(listQuery, manualDialListID).Scan(&campaignCIDOverride, &cidGroupID)

	// Defaults if server info not fully available
	if asteriskVersion == "" {
		asteriskVersion = "11"
	}
	if serverID == "" {
		serverID = serverIP
	}

	// Build dial string using campaign settings
	localOutPrefix := "9"
	localDialTimeout := 60
	if dialTimeout > 4 {
		localDialTimeout = dialTimeout
	}
	localDialTimeoutMS := localDialTimeout * 1000

	if len(dialPrefix) > 0 && !strings.Contains(dialPrefix, "x") {
		localOutPrefix = dialPrefix
	} else if strings.Contains(dialPrefix, "x") {
		localOutPrefix = ""
	}

	// Determine VDAD dial extension (request override > campaign setting > server answer_transfer_agent > default)
	req.VdadExten = strings.TrimSpace(req.VdadExten)
	extContext = strings.TrimSpace(extContext)

	vdadDialExten := "8368"
	if answerTransferAgent.Valid && len(strings.TrimSpace(answerTransferAgent.String)) > 0 {
		vdadDialExten = strings.TrimSpace(answerTransferAgent.String)
	}
	if len(req.VdadExten) > 0 {
		vdadDialExten = req.VdadExten
	} else if campaignVdadExten.Valid && len(strings.TrimSpace(campaignVdadExten.String)) > 0 {
		vdadDialExten = strings.TrimSpace(campaignVdadExten.String)
	}

	// Add routing prefix for Asterisk 12+
	majorVersionStr := strings.Split(asteriskVersion, ".")[0]
	majorVersion, _ := strconv.Atoi(majorVersionStr)
	if majorVersion >= 12 && len(routingPrefix) > 0 {
		vdadDialExten = routingPrefix + vdadDialExten
	}

	// Build dial string
	var nDialString string
	if omitPhoneCode.Valid && strings.ToUpper(omitPhoneCode.String) == "Y" {
		nDialString = fmt.Sprintf("%s%s", localOutPrefix, req.PhoneNumber)
	} else {
		nDialString = fmt.Sprintf("%s%s%s", localOutPrefix, req.PhoneCode, req.PhoneNumber)
	}

	// Set default context (prefer server ext_context)
	if serverExtContext.Valid && len(strings.TrimSpace(serverExtContext.String)) > 0 {
		extContext = strings.TrimSpace(serverExtContext.String)
	}
	if len(extContext) < 1 {
		extContext = "default"
	}

	// Generate unique caller ID: VmddhhmmssLLLLLLLLLLL
	padLeadID := fmt.Sprintf("%010d", leadID)
	if len(padLeadID) > 10 {
		padLeadID = padLeadID[len(padLeadID)-10:]
	}
	vQueryCID := fmt.Sprintf("V%s%s", cidDate, padLeadID)

	// Determine caller ID to use
	ccid := campaignCID
	if campaignCIDOverride.Valid && len(campaignCIDOverride.String) > 6 {
		ccid = campaignCIDOverride.String
	}

	var cidString string
	if len(ccid) > 6 {
		cidString = fmt.Sprintf("\"%s\" <%s>", vQueryCID, ccid)
	} else {
		cidString = vQueryCID
	}

	// Build channel string
	localDEF := "Local/"
	localAMP := "@"
	channel := fmt.Sprintf("%s%s%s%s", localDEF, nDialString, localAMP, extContext)

	// Insert into vicidial_manager table
	managerQuery := `
		INSERT INTO vicidial_manager
		(uniqueid, entry_date, status, response, server_ip, channel, action,
		 callerid, cmd_line_b, cmd_line_c, cmd_line_d, cmd_line_e, cmd_line_f,
		 cmd_line_g, cmd_line_k)
		VALUES ('', ?, 'NEW', 'N', ?, '', 'Originate', ?, ?, ?, ?, ?, ?, ?, ?)
	`

	cmdLineB := fmt.Sprintf("Exten: %s", vdadDialExten)
	cmdLineC := fmt.Sprintf("Context: %s", extContext)
	cmdLineD := fmt.Sprintf("Channel: %s", channel)
	cmdLineE := "Priority: 1"
	cmdLineF := fmt.Sprintf("Callerid: %s", cidString)
	cmdLineG := fmt.Sprintf("Timeout: %d", localDialTimeoutMS)
	cmdLineK := fmt.Sprintf("VDACnote: %s|%d|%s|%s|OUT|MAIN|99",
		req.CampaignID, leadID, req.PhoneCode, req.PhoneNumber)

	result, err = h.DB.Exec(managerQuery, sqlDate, serverIP, vQueryCID,
		cmdLineB, cmdLineC, cmdLineD, cmdLineE, cmdLineF, cmdLineG, cmdLineK)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to queue test call: "+err.Error())
		return
	}

	managerID, _ := result.LastInsertId()

	// Insert into vicidial_auto_calls
	autoCallQuery := `
		INSERT INTO vicidial_auto_calls
		(server_ip, campaign_id, status, lead_id, callerid, phone_code,
		 phone_number, call_time, call_type, alt_dial, queue_priority)
		VALUES (?, ?, 'SENT', ?, ?, ?, ?, ?, 'OUT', 'MAIN', '99')
	`
	_, err = h.DB.Exec(autoCallQuery, serverIP, req.CampaignID, leadID,
		vQueryCID, req.PhoneCode, req.PhoneNumber, sqlDate)
	if err != nil {
		fmt.Printf("Warning: Failed to log auto call: %v\n", err)
	}

	// Insert into vicidial_dial_log
	dialLogQuery := `
		INSERT INTO vicidial_dial_log SET
		caller_code=?, lead_id=?, server_ip=?, call_date=?, extension=?,
		channel=?, timeout=?, outbound_cid=?, context=?
	`
	_, err = h.DB.Exec(dialLogQuery, vQueryCID, leadID, serverIP, sqlDate,
		vdadDialExten, channel, localDialTimeoutMS, cidString, extContext)
	if err != nil {
		fmt.Printf("Warning: Failed to log dial: %v\n", err)
	}

	// Insert into vicidial_dial_cid_log
	cidLogQuery := `
		INSERT INTO vicidial_dial_cid_log SET
		caller_code=?, call_date=?, call_type='MANUAL', call_alt='MAIN',
		outbound_cid=?, outbound_cid_type='CAMPAIGN_TEST'
	`
	_, err = h.DB.Exec(cidLogQuery, vQueryCID, sqlDate, ccid)
	if err != nil {
		fmt.Printf("Warning: Failed to log CID: %v\n", err)
	}

	// Insert into vicidial_user_dial_log
	userDialLogQuery := `
		INSERT INTO vicidial_user_dial_log SET
		caller_code=?, user=?, call_date=?, call_type='M',
		notes='API test call CAMPAIGN_TEST'
	`
	_, err = h.DB.Exec(userDialLogQuery, vQueryCID, req.User, sqlDate)
	if err != nil {
		fmt.Printf("Warning: Failed to log user dial: %v\n", err)
	}

	// Insert into user_call_log
	userCallLogQuery := `
		INSERT INTO user_call_log
		(user, call_date, call_type, server_ip, phone_number, number_dialed,
		 lead_id, callerid, group_alias_id, preset_name)
		VALUES (?, ?, 'API', ?, ?, ?, ?, ?, '', '')
	`
	_, err = h.DB.Exec(userCallLogQuery, req.User, sqlDate, serverIP,
		req.PhoneNumber, nDialString, leadID, ccid)
	if err != nil {
		fmt.Printf("Warning: Failed to log user call: %v\n", err)
	}

	respondWithSuccess(w, "Test call placed successfully", map[string]interface{}{
		"caller_code":   vQueryCID,
		"manager_id":    managerID,
		"lead_id":       leadID,
		"campaign_id":   req.CampaignID,
		"campaign_name": campaignName,
		"phone_number":  req.PhoneNumber,
		"phone_code":    req.PhoneCode,
		"server_ip":     serverIP,
		"server_id":     serverID,
		"channel":       channel,
		"extension":     vdadDialExten,
		"dial_string":   nDialString,
		"caller_id":     cidString,
		"call_date":     sqlDate,
	})
}

// GetTestCallStatus retrieves the status of a test call
func (h *Handler) GetTestCallStatus(w http.ResponseWriter, r *http.Request) {
	callerCode := r.URL.Query().Get("caller_code")

	if callerCode == "" {
		respondWithError(w, http.StatusBadRequest, "caller_code is required")
		return
	}

	// Check vicidial_manager for the call status
	query := `
		SELECT man_id, entry_date, status, response, action, cmd_line_b, cmd_line_c
		FROM vicidial_manager
		WHERE callerid LIKE ?
		ORDER BY entry_date DESC
		LIMIT 1
	`

	var manID int
	var entryDate time.Time
	var status, response, action, channel, context string

	err := h.DB.QueryRow(query, "%"+callerCode+"%").Scan(
		&manID, &entryDate, &status, &response, &action, &channel, &context)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Test call not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve call status: "+err.Error())
		return
	}

	// Extract channel from cmd_line_b
	channelClean := strings.TrimPrefix(channel, "Channel: ")

	// Check dial log for more details
	dialQuery := `
		SELECT call_date, extension, channel
		FROM vicidial_dial_log
		WHERE caller_code = ?
		ORDER BY call_date DESC
		LIMIT 1
	`

	var callDate time.Time
	var extension, dialChannel string
	h.DB.QueryRow(dialQuery, callerCode).Scan(&callDate, &extension, &dialChannel)

	respondWithSuccess(w, "Test call status retrieved", map[string]interface{}{
		"caller_code": callerCode,
		"manager_id":  manID,
		"entry_date":  entryDate,
		"status":      status,
		"response":    response,
		"action":      action,
		"channel":     channelClean,
		"context":     strings.TrimPrefix(context, "Context: "),
		"extension":   extension,
		"call_date":   callDate,
	})
}

// ListTestCalls lists recent test calls
func (h *Handler) ListTestCalls(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	phoneLogin := r.URL.Query().Get("phone_login")

	if limit == "" {
		limit = "50"
	}

	query := `
		SELECT dl.caller_code, dl.call_date, dl.extension, dl.channel, dl.server_ip,
			   vm.status, vm.response
		FROM vicidial_dial_log dl
		LEFT JOIN vicidial_manager vm ON dl.caller_code = vm.callerid
		WHERE dl.caller_code LIKE 'TC%'
	`
	args := []interface{}{}

	if phoneLogin != "" {
		query += " AND dl.caller_code LIKE ?"
		args = append(args, "%"+phoneLogin)
	}

	query += " ORDER BY dl.call_date DESC LIMIT ?"
	args = append(args, limit)

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve test calls: "+err.Error())
		return
	}
	defer rows.Close()

	type TestCallEntry struct {
		CallerCode string    `json:"caller_code"`
		CallDate   time.Time `json:"call_date"`
		Extension  string    `json:"extension"`
		Channel    string    `json:"channel"`
		ServerIP   string    `json:"server_ip"`
		Status     string    `json:"status"`
		Response   string    `json:"response"`
	}

	testCalls := []TestCallEntry{}
	for rows.Next() {
		var tc TestCallEntry
		var status, response sql.NullString

		rows.Scan(&tc.CallerCode, &tc.CallDate, &tc.Extension, &tc.Channel,
			&tc.ServerIP, &status, &response)

		if status.Valid {
			tc.Status = status.String
		}
		if response.Valid {
			tc.Response = response.String
		}

		testCalls = append(testCalls, tc)
	}

	respondWithSuccess(w, "Test calls retrieved", map[string]interface{}{
		"count": len(testCalls),
		"calls": testCalls,
	})
}
