package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vicidb/non-agent-api/models"
)

// RecordingLookup searches for call recordings
func (h *Handler) RecordingLookup(w http.ResponseWriter, r *http.Request) {
	leadID := r.URL.Query().Get("lead_id")
	user := r.URL.Query().Get("user")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	query := `
		SELECT recording_id, channel, server_ip, extension, start_time,
			   end_time, length_in_sec, filename, location, lead_id, user, vicidial_id
		FROM recording_log WHERE 1=1
	`
	args := []interface{}{}

	if leadID != "" {
		query += " AND lead_id = ?"
		args = append(args, leadID)
	}
	if user != "" {
		query += " AND user = ?"
		args = append(args, user)
	}
	if startDate != "" {
		query += " AND start_time >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND start_time <= ?"
		args = append(args, endDate)
	}

	query += " ORDER BY start_time DESC LIMIT 100"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to search recordings: "+err.Error())
		return
	}
	defer rows.Close()

	recordings := []models.Recording{}
	for rows.Next() {
		var rec models.Recording
		rows.Scan(&rec.RecordingID, &rec.Channel, &rec.ServerIP, &rec.Extension,
			&rec.StartTime, &rec.EndTime, &rec.Length, &rec.Filename,
			&rec.Location, &rec.LeadID, &rec.User, &rec.VicidialID)
		recordings = append(recordings, rec)
	}

	respondWithSuccess(w, "Recordings retrieved", recordings)
}

// DIDLogExport exports DID call logs
func (h *Handler) DIDLogExport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	didPattern := r.URL.Query().Get("did_pattern")

	query := `
		SELECT uniqueid, server_ip, channel, caller_id_number, caller_id_name,
			   extension, call_date, did_id, did_route
		FROM vicidial_did_log WHERE 1=1
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
	if didPattern != "" {
		query += " AND did_id LIKE ?"
		args = append(args, "%"+didPattern+"%")
	}

	query += " ORDER BY call_date DESC LIMIT 1000"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to export DID logs: "+err.Error())
		return
	}
	defer rows.Close()

	type DIDLog struct {
		UniqueID         string    `json:"uniqueid"`
		ServerIP         string    `json:"server_ip"`
		Channel          string    `json:"channel"`
		CallerIDNumber   string    `json:"caller_id_number"`
		CallerIDName     string    `json:"caller_id_name"`
		Extension        string    `json:"extension"`
		CallDate         time.Time `json:"call_date"`
		DIDID            string    `json:"did_id"`
		DIDRoute         string    `json:"did_route"`
	}

	logs := []DIDLog{}
	for rows.Next() {
		var log DIDLog
		rows.Scan(&log.UniqueID, &log.ServerIP, &log.Channel, &log.CallerIDNumber,
			&log.CallerIDName, &log.Extension, &log.CallDate, &log.DIDID, &log.DIDRoute)
		logs = append(logs, log)
	}

	respondWithSuccess(w, "DID logs exported", logs)
}

// PhoneNumberLog retrieves call history for a phone number
func (h *Handler) PhoneNumberLog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phoneNumber := vars["phone"]

	query := `
		SELECT uniqueid, lead_id, list_id, campaign_id, call_date, start_epoch,
			   end_epoch, length_in_sec, status, phone_code, phone_number, user, comments
		FROM vicidial_log WHERE phone_number = ?
		ORDER BY call_date DESC LIMIT 50
	`

	rows, err := h.DB.Query(query, phoneNumber)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve phone logs: "+err.Error())
		return
	}
	defer rows.Close()

	logs := []models.CallLog{}
	for rows.Next() {
		var log models.CallLog
		rows.Scan(&log.UniqueID, &log.LeadID, &log.ListID, &log.CampaignID,
			&log.CallDate, &log.StartEpoch, &log.EndEpoch, &log.Length,
			&log.Status, &log.PhoneCode, &log.PhoneNumber, &log.User, &log.Comments)
		logs = append(logs, log)
	}

	respondWithSuccess(w, "Phone number history retrieved", logs)
}

// AgentStatsExport exports agent statistics
func (h *Handler) AgentStatsExport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	user := r.URL.Query().Get("user")
	campaignID := r.URL.Query().Get("campaign_id")

	query := `
		SELECT user, event_time, campaign_id, pause_epoch, pause_sec,
			   wait_epoch, wait_sec, talk_epoch, talk_sec, dispo_epoch, dispo_sec,
			   status, calls
		FROM vicidial_agent_log WHERE 1=1
	`
	args := []interface{}{}

	if startDate != "" {
		query += " AND event_time >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND event_time <= ?"
		args = append(args, endDate)
	}
	if user != "" {
		query += " AND user = ?"
		args = append(args, user)
	}
	if campaignID != "" {
		query += " AND campaign_id = ?"
		args = append(args, campaignID)
	}

	query += " ORDER BY event_time DESC LIMIT 1000"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to export agent stats: "+err.Error())
		return
	}
	defer rows.Close()

	type AgentStat struct {
		User        string    `json:"user"`
		EventTime   time.Time `json:"event_time"`
		CampaignID  string    `json:"campaign_id"`
		PauseEpoch  int64     `json:"pause_epoch"`
		PauseSec    int       `json:"pause_sec"`
		WaitEpoch   int64     `json:"wait_epoch"`
		WaitSec     int       `json:"wait_sec"`
		TalkEpoch   int64     `json:"talk_epoch"`
		TalkSec     int       `json:"talk_sec"`
		DispoEpoch  int64     `json:"dispo_epoch"`
		DispoSec    int       `json:"dispo_sec"`
		Status      string    `json:"status"`
		Calls       int       `json:"calls"`
	}

	stats := []AgentStat{}
	for rows.Next() {
		var stat AgentStat
		rows.Scan(&stat.User, &stat.EventTime, &stat.CampaignID, &stat.PauseEpoch,
			&stat.PauseSec, &stat.WaitEpoch, &stat.WaitSec, &stat.TalkEpoch,
			&stat.TalkSec, &stat.DispoEpoch, &stat.DispoSec, &stat.Status, &stat.Calls)
		stats = append(stats, stat)
	}

	respondWithSuccess(w, "Agent statistics exported", stats)
}

// CallStatusStats retrieves call status statistics
func (h *Handler) CallStatusStats(w http.ResponseWriter, r *http.Request) {
	campaignID := r.URL.Query().Get("campaign_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	query := `
		SELECT status, COUNT(*) as count, AVG(length_in_sec) as avg_length
		FROM vicidial_log WHERE 1=1
	`
	args := []interface{}{}

	if campaignID != "" {
		query += " AND campaign_id = ?"
		args = append(args, campaignID)
	}
	if startDate != "" {
		query += " AND call_date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND call_date <= ?"
		args = append(args, endDate)
	}

	query += " GROUP BY status ORDER BY count DESC"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve status stats: "+err.Error())
		return
	}
	defer rows.Close()

	type StatusStat struct {
		Status    string  `json:"status"`
		Count     int     `json:"count"`
		AvgLength float64 `json:"avg_length"`
	}

	stats := []StatusStat{}
	for rows.Next() {
		var stat StatusStat
		rows.Scan(&stat.Status, &stat.Count, &stat.AvgLength)
		stats = append(stats, stat)
	}

	respondWithSuccess(w, "Status statistics retrieved", stats)
}

// CallDispoReport retrieves call disposition report
func (h *Handler) CallDispoReport(w http.ResponseWriter, r *http.Request) {
	campaignID := r.URL.Query().Get("campaign_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	query := `
		SELECT status, user, COUNT(*) as count
		FROM vicidial_log WHERE 1=1
	`
	args := []interface{}{}

	if campaignID != "" {
		query += " AND campaign_id = ?"
		args = append(args, campaignID)
	}
	if startDate != "" {
		query += " AND call_date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND call_date <= ?"
		args = append(args, endDate)
	}

	query += " GROUP BY status, user ORDER BY count DESC LIMIT 500"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve dispo report: "+err.Error())
		return
	}
	defer rows.Close()

	type DispoStat struct {
		Status string `json:"status"`
		User   string `json:"user"`
		Count  int    `json:"count"`
	}

	stats := []DispoStat{}
	for rows.Next() {
		var stat DispoStat
		rows.Scan(&stat.Status, &stat.User, &stat.Count)
		stats = append(stats, stat)
	}

	respondWithSuccess(w, "Disposition report retrieved", stats)
}

// BlindMonitor initiates blind monitoring of an agent call
func (h *Handler) BlindMonitor(w http.ResponseWriter, r *http.Request) {
	var req struct {
		User      string `json:"user"`
		Extension string `json:"extension"`
		ServerIP  string `json:"server_ip"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Get the agent's current channel
	query := "SELECT channel FROM vicidial_live_agents WHERE user = ? AND server_ip = ?"
	var channel string
	err := h.DB.QueryRow(query, req.User, req.ServerIP).Scan(&channel)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Agent not found or not in call")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get agent info: "+err.Error())
		return
	}

	respondWithSuccess(w, "Monitor request processed", map[string]string{
		"user":      req.User,
		"channel":   channel,
		"extension": req.Extension,
		"message":   "Blind monitor initiated - use Asterisk manager to complete connection",
	})
}
