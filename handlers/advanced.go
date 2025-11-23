package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddGroupAlias adds a group alias
func (h *Handler) AddGroupAlias(w http.ResponseWriter, r *http.Request) {
	var req struct {
		GroupAliasID   string `json:"group_alias_id"`
		GroupAliasName string `json:"group_alias_name"`
		CallerIDGroup  string `json:"caller_id_group"`
		Active         string `json:"active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Active == "" {
		req.Active = "Y"
	}

	query := `
		INSERT INTO vicidial_group_aliases (group_alias_id, group_alias_name, caller_id_group, active)
		VALUES (?, ?, ?, ?)
	`

	_, err := h.DB.Exec(query, req.GroupAliasID, req.GroupAliasName, req.CallerIDGroup, req.Active)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add group alias: "+err.Error())
		return
	}

	respondWithSuccess(w, "Group alias added successfully", map[string]string{
		"group_alias_id": req.GroupAliasID,
	})
}

// UpdateLogEntry updates a call log entry
func (h *Handler) UpdateLogEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryID := vars["entry_id"]

	var req struct {
		Status   string `json:"status"`
		Comments string `json:"comments"`
		UserGroup string `json:"user_group"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_log SET status = ?, comments = ?, user_group = ?
		WHERE uniqueid = ?
	`

	result, err := h.DB.Exec(query, req.Status, req.Comments, req.UserGroup, entryID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update log entry: "+err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithError(w, http.StatusNotFound, "Log entry not found")
		return
	}

	respondWithSuccess(w, "Log entry updated successfully", nil)
}

// UpdateCIDGroupEntry updates a caller ID group entry
func (h *Handler) UpdateCIDGroupEntry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entryID := vars["entry_id"]

	var req struct {
		CallerIDNumber string `json:"caller_id_number"`
		CallerIDName   string `json:"caller_id_name"`
		Active         string `json:"active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_inbound_group_cid SET
			caller_id_number = ?, caller_id_name = ?, active = ?
		WHERE cid_id = ?
	`

	_, err := h.DB.Exec(query, req.CallerIDNumber, req.CallerIDName, req.Active, entryID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update CID group entry: "+err.Error())
		return
	}

	respondWithSuccess(w, "CID group entry updated successfully", nil)
}

// UpdateAltURL updates an alternate URL
func (h *Handler) UpdateAltURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	urlID := vars["url_id"]

	var req struct {
		URL        string `json:"url"`
		URLType    string `json:"url_type"`
		URLRank    int    `json:"url_rank"`
		Active     string `json:"active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_url_multi SET url = ?, url_type = ?, url_rank = ?, active = ?
		WHERE url_id = ?
	`

	_, err := h.DB.Exec(query, req.URL, req.URLType, req.URLRank, req.Active, urlID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update alternate URL: "+err.Error())
		return
	}

	respondWithSuccess(w, "Alternate URL updated successfully", nil)
}

// UpdatePresets updates system presets
func (h *Handler) UpdatePresets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	presetID := vars["preset_id"]

	var req struct {
		PresetName  string `json:"preset_name"`
		PresetValue string `json:"preset_value"`
		MenuID      string `json:"menu_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_lists_fields SET field_default = ?
		WHERE field_id = ? AND field_name = ?
	`

	_, err := h.DB.Exec(query, req.PresetValue, presetID, req.PresetName)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update preset: "+err.Error())
		return
	}

	respondWithSuccess(w, "Preset updated successfully", nil)
}

// CallidInfo retrieves call information by call ID
func (h *Handler) CallidInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	callID := vars["call_id"]

	// Check vicidial_log first
	query := `
		SELECT uniqueid, lead_id, list_id, campaign_id, call_date,
			   start_epoch, end_epoch, length_in_sec, status, phone_number,
			   user, comments, processed, user_group, term_reason
		FROM vicidial_log WHERE uniqueid = ?
	`

	var callInfo struct {
		UniqueID      string `json:"uniqueid"`
		LeadID        int    `json:"lead_id"`
		ListID        int    `json:"list_id"`
		CampaignID    string `json:"campaign_id"`
		CallDate      string `json:"call_date"`
		StartEpoch    int64  `json:"start_epoch"`
		EndEpoch      int64  `json:"end_epoch"`
		Length        int    `json:"length_in_sec"`
		Status        string `json:"status"`
		PhoneNumber   string `json:"phone_number"`
		User          string `json:"user"`
		Comments      string `json:"comments"`
		Processed     string `json:"processed"`
		UserGroup     string `json:"user_group"`
		TermReason    string `json:"term_reason"`
	}

	err := h.DB.QueryRow(query, callID).Scan(
		&callInfo.UniqueID, &callInfo.LeadID, &callInfo.ListID, &callInfo.CampaignID,
		&callInfo.CallDate, &callInfo.StartEpoch, &callInfo.EndEpoch, &callInfo.Length,
		&callInfo.Status, &callInfo.PhoneNumber, &callInfo.User, &callInfo.Comments,
		&callInfo.Processed, &callInfo.UserGroup, &callInfo.TermReason,
	)

	if err == sql.ErrNoRows {
		// Try vicidial_closer_log
		query = `
			SELECT closecallid as uniqueid, lead_id, list_id, campaign_id, call_date,
				   start_epoch, end_epoch, length_in_sec, status, phone_number,
				   user, comments, queue_seconds, user_group, term_reason
			FROM vicidial_closer_log WHERE closecallid = ?
		`
		var queueSeconds int
		err = h.DB.QueryRow(query, callID).Scan(
			&callInfo.UniqueID, &callInfo.LeadID, &callInfo.ListID, &callInfo.CampaignID,
			&callInfo.CallDate, &callInfo.StartEpoch, &callInfo.EndEpoch, &callInfo.Length,
			&callInfo.Status, &callInfo.PhoneNumber, &callInfo.User, &callInfo.Comments,
			&queueSeconds, &callInfo.UserGroup, &callInfo.TermReason,
		)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Call ID not found")
			return
		}
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve call info: "+err.Error())
		return
	}

	respondWithSuccess(w, "Call information retrieved", callInfo)
}

// CCCLeadInfo retrieves cross-campaign calling lead information
func (h *Handler) CCCLeadInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leadID := vars["lead_id"]

	query := `
		SELECT l.lead_id, l.list_id, l.phone_number, l.first_name, l.last_name,
			   l.status, l.called_count,
			   COUNT(DISTINCT vl.uniqueid) as total_calls,
			   COUNT(DISTINCT vl.campaign_id) as campaigns_called,
			   MAX(vl.call_date) as last_call_date
		FROM vicidial_list l
		LEFT JOIN vicidial_log vl ON l.lead_id = vl.lead_id
		WHERE l.lead_id = ?
		GROUP BY l.lead_id, l.list_id, l.phone_number, l.first_name, l.last_name, l.status, l.called_count
	`

	var cccInfo struct {
		LeadID          int    `json:"lead_id"`
		ListID          int    `json:"list_id"`
		PhoneNumber     string `json:"phone_number"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		Status          string `json:"status"`
		CalledCount     int    `json:"called_count"`
		TotalCalls      int    `json:"total_calls"`
		CampaignsCalled int    `json:"campaigns_called"`
		LastCallDate    string `json:"last_call_date"`
	}

	err := h.DB.QueryRow(query, leadID).Scan(
		&cccInfo.LeadID, &cccInfo.ListID, &cccInfo.PhoneNumber, &cccInfo.FirstName,
		&cccInfo.LastName, &cccInfo.Status, &cccInfo.CalledCount, &cccInfo.TotalCalls,
		&cccInfo.CampaignsCalled, &cccInfo.LastCallDate,
	)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Lead not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve CCC lead info: "+err.Error())
		return
	}

	respondWithSuccess(w, "CCC lead information retrieved", cccInfo)
}
