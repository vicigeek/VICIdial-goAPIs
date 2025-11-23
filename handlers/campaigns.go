package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vicidb/non-agent-api/models"
)

// UpdateCampaign updates campaign settings
func (h *Handler) UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignID := vars["campaign_id"]

	var campaign models.Campaign
	if err := json.NewDecoder(r.Body).Decode(&campaign); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_campaigns SET
			campaign_name = ?, active = ?, dial_status = ?,
			lead_order = ?, dial_method = ?, auto_dial_level = ?,
			local_call_time = ?, dial_prefix = ?
		WHERE campaign_id = ?
	`

	_, err := h.DB.Exec(query,
		campaign.CampaignName, campaign.Active, campaign.DialStatus,
		campaign.LeadOrder, campaign.DialMethod, campaign.AutoDialLevel,
		campaign.LocalCallTime, campaign.DialPrefix, campaignID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update campaign: "+err.Error())
		return
	}

	respondWithSuccess(w, "Campaign updated successfully", map[string]string{"campaign_id": campaignID})
}

// CampaignsList retrieves all campaigns
func (h *Handler) CampaignsList(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("active")

	query := "SELECT campaign_id, campaign_name, active, dial_status, dial_method, auto_dial_level FROM vicidial_campaigns"
	args := []interface{}{}

	if active != "" {
		query += " WHERE active = ?"
		args = append(args, active)
	}

	query += " ORDER BY campaign_name"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve campaigns: "+err.Error())
		return
	}
	defer rows.Close()

	campaigns := []models.Campaign{}
	for rows.Next() {
		var camp models.Campaign
		rows.Scan(&camp.CampaignID, &camp.CampaignName, &camp.Active,
			&camp.DialStatus, &camp.DialMethod, &camp.AutoDialLevel)
		campaigns = append(campaigns, camp)
	}

	respondWithSuccess(w, "Campaigns retrieved", campaigns)
}

// HopperList retrieves leads in campaign hopper
func (h *Handler) HopperList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignID := vars["campaign_id"]

	query := `
		SELECT h.hopper_id, h.lead_id, h.campaign_id, h.status, h.user,
			   h.list_id, h.priority, l.phone_number, l.first_name, l.last_name
		FROM vicidial_hopper h
		LEFT JOIN vicidial_list l ON h.lead_id = l.lead_id
		WHERE h.campaign_id = ?
		ORDER BY h.priority DESC, h.hopper_id
		LIMIT 100
	`

	rows, err := h.DB.Query(query, campaignID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve hopper: "+err.Error())
		return
	}
	defer rows.Close()

	type HopperEntry struct {
		HopperID    int    `json:"hopper_id"`
		LeadID      int    `json:"lead_id"`
		CampaignID  string `json:"campaign_id"`
		Status      string `json:"status"`
		User        string `json:"user"`
		ListID      int    `json:"list_id"`
		Priority    int    `json:"priority"`
		PhoneNumber string `json:"phone_number"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
	}

	hopperEntries := []HopperEntry{}
	for rows.Next() {
		var entry HopperEntry
		rows.Scan(&entry.HopperID, &entry.LeadID, &entry.CampaignID, &entry.Status,
			&entry.User, &entry.ListID, &entry.Priority, &entry.PhoneNumber,
			&entry.FirstName, &entry.LastName)
		hopperEntries = append(hopperEntries, entry)
	}

	respondWithSuccess(w, "Hopper entries retrieved", hopperEntries)
}

// HopperBulkInsert inserts multiple leads into hopper
func (h *Handler) HopperBulkInsert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	campaignID := vars["campaign_id"]

	var req struct {
		LeadIDs  []int  `json:"lead_ids"`
		Priority int    `json:"priority"`
		Source   string `json:"source"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(req.LeadIDs) == 0 {
		respondWithError(w, http.StatusBadRequest, "No lead IDs provided")
		return
	}

	if req.Source == "" {
		req.Source = "API"
	}

	tx, err := h.DB.Begin()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to start transaction")
		return
	}

	insertStmt := `
		INSERT IGNORE INTO vicidial_hopper (lead_id, campaign_id, status, priority, source, list_id)
		SELECT lead_id, ?, 'READY', ?, ?, list_id FROM vicidial_list WHERE lead_id = ?
	`

	insertCount := 0
	for _, leadID := range req.LeadIDs {
		result, err := tx.Exec(insertStmt, campaignID, req.Priority, req.Source, leadID)
		if err != nil {
			tx.Rollback()
			respondWithError(w, http.StatusInternalServerError, "Failed to insert lead: "+err.Error())
			return
		}
		rows, _ := result.RowsAffected()
		insertCount += int(rows)
	}

	tx.Commit()
	respondWithSuccess(w, "Leads inserted to hopper", map[string]int{
		"requested": len(req.LeadIDs),
		"inserted":  insertCount,
	})
}

// GetCampaignsWithLists retrieves all campaigns with their associated lists in JSON format
func (h *Handler) GetCampaignsWithLists(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("active")
	campaignID := r.URL.Query().Get("campaign_id")

	// Query to get campaigns
	campaignQuery := `
		SELECT campaign_id, campaign_name, active, dial_status, dial_method,
			   auto_dial_level, lead_order, local_call_time
		FROM vicidial_campaigns
		WHERE 1=1
	`
	campaignArgs := []interface{}{}

	if active != "" {
		campaignQuery += " AND active = ?"
		campaignArgs = append(campaignArgs, active)
	}
	if campaignID != "" {
		campaignQuery += " AND campaign_id = ?"
		campaignArgs = append(campaignArgs, campaignID)
	}

	campaignQuery += " ORDER BY campaign_name"

	campaignRows, err := h.DB.Query(campaignQuery, campaignArgs...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve campaigns: "+err.Error())
		return
	}
	defer campaignRows.Close()

	type CampaignWithLists struct {
		CampaignID    string `json:"campaign_id"`
		CampaignName  string `json:"campaign_name"`
		Active        string `json:"active"`
		DialStatus    string `json:"dial_status"`
		DialMethod    string `json:"dial_method"`
		AutoDialLevel string `json:"auto_dial_level"`
		LeadOrder     string `json:"lead_order"`
		LocalCallTime string `json:"local_call_time"`
		Lists         []struct {
			ListID          int    `json:"list_id"`
			ListName        string `json:"list_name"`
			Active          string `json:"active"`
			ListDescription string `json:"list_description"`
			LeadCount       int    `json:"lead_count"`
		} `json:"lists"`
	}

	campaigns := []CampaignWithLists{}

	for campaignRows.Next() {
		var camp CampaignWithLists
		err := campaignRows.Scan(
			&camp.CampaignID, &camp.CampaignName, &camp.Active,
			&camp.DialStatus, &camp.DialMethod, &camp.AutoDialLevel,
			&camp.LeadOrder, &camp.LocalCallTime,
		)
		if err != nil {
			continue
		}

		// Get lists for this campaign
		listQuery := `
			SELECT l.list_id, l.list_name, l.active, l.list_description,
				   COUNT(vl.lead_id) as lead_count
			FROM vicidial_lists l
			LEFT JOIN vicidial_list vl ON l.list_id = vl.list_id
			WHERE l.campaign_id = ?
			GROUP BY l.list_id, l.list_name, l.active, l.list_description
			ORDER BY l.list_name
		`

		listRows, err := h.DB.Query(listQuery, camp.CampaignID)
		if err != nil {
			continue
		}

		camp.Lists = []struct {
			ListID          int    `json:"list_id"`
			ListName        string `json:"list_name"`
			Active          string `json:"active"`
			ListDescription string `json:"list_description"`
			LeadCount       int    `json:"lead_count"`
		}{}

		for listRows.Next() {
			var list struct {
				ListID          int    `json:"list_id"`
				ListName        string `json:"list_name"`
				Active          string `json:"active"`
				ListDescription string `json:"list_description"`
				LeadCount       int    `json:"lead_count"`
			}
			listRows.Scan(&list.ListID, &list.ListName, &list.Active, &list.ListDescription, &list.LeadCount)
			camp.Lists = append(camp.Lists, list)
		}
		listRows.Close()

		campaigns = append(campaigns, camp)
	}

	respondWithSuccess(w, "Campaigns with lists retrieved successfully", map[string]interface{}{
		"count":     len(campaigns),
		"campaigns": campaigns,
	})
}
