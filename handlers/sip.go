package handlers

import (
	"database/sql"
	"net/http"
	"time"
)

// SIPLog represents a carrier log entry
type SIPLog struct {
	UniqueID         string    `json:"uniqueid"`
	CallDate         time.Time `json:"call_date"`
	ServerIP         string    `json:"server_ip"`
	LeadID           int       `json:"lead_id"`
	HangupCause      int       `json:"hangup_cause"`
	DialStatus       string    `json:"dialstatus"`
	Channel          string    `json:"channel"`
	DialTime         int       `json:"dial_time"`
	AnsweredTime     int       `json:"answered_time"`
	SIPHangupCause   int       `json:"sip_hangup_cause"`
	SIPHangupReason  string    `json:"sip_hangup_reason"`
	CallerCode       string    `json:"caller_code"`
}

// GetSIPLog fetches data from carrier log
func (h *Handler) GetSIPLog(w http.ResponseWriter, r *http.Request) {
	// Query parameters
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	leadID := r.URL.Query().Get("lead_id")
	serverIP := r.URL.Query().Get("server_ip")
	dialStatus := r.URL.Query().Get("dialstatus")
	sipHangupCause := r.URL.Query().Get("sip_hangup_cause")
	limit := r.URL.Query().Get("limit")

	if limit == "" {
		limit = "100"
	}

	query := `
		SELECT uniqueid, call_date, server_ip, lead_id, hangup_cause,
			   dialstatus, channel, dial_time, answered_time,
			   sip_hangup_cause, sip_hangup_reason, caller_code
		FROM vicidial_carrier_log
		WHERE 1=1
	`
	args := []interface{}{}

	if startDate != "" {
		query += " AND call_date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND call_date <= ?"
		args = append(args, endDate)
	}
	if leadID != "" {
		query += " AND lead_id = ?"
		args = append(args, leadID)
	}
	if serverIP != "" {
		query += " AND server_ip = ?"
		args = append(args, serverIP)
	}
	if dialStatus != "" {
		query += " AND dialstatus = ?"
		args = append(args, dialStatus)
	}
	if sipHangupCause != "" {
		query += " AND sip_hangup_cause = ?"
		args = append(args, sipHangupCause)
	}

	query += " ORDER BY call_date DESC LIMIT ?"
	args = append(args, limit)

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve carrier logs: "+err.Error())
		return
	}
	defer rows.Close()

	logs := []SIPLog{}
	for rows.Next() {
		var log SIPLog
		var callDate sql.NullTime
		var leadID sql.NullInt64

		err := rows.Scan(
			&log.UniqueID, &callDate, &log.ServerIP, &leadID,
			&log.HangupCause, &log.DialStatus, &log.Channel,
			&log.DialTime, &log.AnsweredTime, &log.SIPHangupCause,
			&log.SIPHangupReason, &log.CallerCode,
		)

		if err != nil {
			continue
		}

		if callDate.Valid {
			log.CallDate = callDate.Time
		}
		if leadID.Valid {
			log.LeadID = int(leadID.Int64)
		}

		logs = append(logs, log)
	}

	respondWithSuccess(w, "SIP/Carrier logs retrieved successfully", map[string]interface{}{
		"count": len(logs),
		"logs":  logs,
	})
}

// GetSIPEventLog fetches SIP event logs
func (h *Handler) GetSIPEventLog(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	sipCallID := r.URL.Query().Get("sip_call_id")
	sipEvent := r.URL.Query().Get("sip_event")
	limit := r.URL.Query().Get("limit")

	if limit == "" {
		limit = "100"
	}

	query := `
		SELECT sip_event_id, sip_call_id, sip_event, event_date,
			   server_ip, caller_id_number, caller_id_name, extension
		FROM vicidial_sip_event_log
		WHERE 1=1
	`
	args := []interface{}{}

	if startDate != "" {
		query += " AND event_date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND event_date <= ?"
		args = append(args, endDate)
	}
	if sipCallID != "" {
		query += " AND sip_call_id = ?"
		args = append(args, sipCallID)
	}
	if sipEvent != "" {
		query += " AND sip_event = ?"
		args = append(args, sipEvent)
	}

	query += " ORDER BY event_date DESC LIMIT ?"
	args = append(args, limit)

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve SIP event logs: "+err.Error())
		return
	}
	defer rows.Close()

	type SIPEventLog struct {
		SIPEventID      int       `json:"sip_event_id"`
		SIPCallID       string    `json:"sip_call_id"`
		SIPEvent        string    `json:"sip_event"`
		EventDate       time.Time `json:"event_date"`
		ServerIP        string    `json:"server_ip"`
		CallerIDNumber  string    `json:"caller_id_number"`
		CallerIDName    string    `json:"caller_id_name"`
		Extension       string    `json:"extension"`
	}

	logs := []SIPEventLog{}
	for rows.Next() {
		var log SIPEventLog
		rows.Scan(&log.SIPEventID, &log.SIPCallID, &log.SIPEvent, &log.EventDate,
			&log.ServerIP, &log.CallerIDNumber, &log.CallerIDName, &log.Extension)
		logs = append(logs, log)
	}

	respondWithSuccess(w, "SIP event logs retrieved successfully", map[string]interface{}{
		"count": len(logs),
		"logs":  logs,
	})
}

// GetLiveSIPChannels retrieves live SIP channels
func (h *Handler) GetLiveSIPChannels(w http.ResponseWriter, r *http.Request) {
	serverIP := r.URL.Query().Get("server_ip")
	channelGroup := r.URL.Query().Get("channel_group")

	query := `
		SELECT channel, server_ip, channel_group, extension, context,
			   caller_id_number, caller_id_name, application, app_data
		FROM live_sip_channels
		WHERE 1=1
	`
	args := []interface{}{}

	if serverIP != "" {
		query += " AND server_ip = ?"
		args = append(args, serverIP)
	}
	if channelGroup != "" {
		query += " AND channel_group = ?"
		args = append(args, channelGroup)
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve live SIP channels: "+err.Error())
		return
	}
	defer rows.Close()

	type LiveSIPChannel struct {
		Channel        string `json:"channel"`
		ServerIP       string `json:"server_ip"`
		ChannelGroup   string `json:"channel_group"`
		Extension      string `json:"extension"`
		Context        string `json:"context"`
		CallerIDNumber string `json:"caller_id_number"`
		CallerIDName   string `json:"caller_id_name"`
		Application    string `json:"application"`
		AppData        string `json:"app_data"`
	}

	channels := []LiveSIPChannel{}
	for rows.Next() {
		var ch LiveSIPChannel
		rows.Scan(&ch.Channel, &ch.ServerIP, &ch.ChannelGroup, &ch.Extension,
			&ch.Context, &ch.CallerIDNumber, &ch.CallerIDName, &ch.Application, &ch.AppData)
		channels = append(channels, ch)
	}

	respondWithSuccess(w, "Live SIP channels retrieved successfully", map[string]interface{}{
		"count":    len(channels),
		"channels": channels,
	})
}
