package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AddDNCPhone adds a phone number to the DNC list
func (h *Handler) AddDNCPhone(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhoneNumber string `json:"phone_number"`
		CampaignID  string `json:"campaign_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.PhoneNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Phone number is required")
		return
	}

	if req.CampaignID == "" {
		req.CampaignID = "---ALL---" // Default campaign ID for global DNC
	}

	query := `
		INSERT IGNORE INTO vicidial_dnc (phone_number, campaign_id, entry_date)
		VALUES (?, ?, NOW())
	`

	result, err := h.DB.Exec(query, req.PhoneNumber, req.CampaignID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add DNC entry: "+err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithSuccess(w, "Phone number already in DNC list", nil)
		return
	}

	respondWithSuccess(w, "Phone number added to DNC list", map[string]string{
		"phone_number": req.PhoneNumber,
		"campaign_id":  req.CampaignID,
	})
}

// DeleteDNCPhone removes a phone number from the DNC list
func (h *Handler) DeleteDNCPhone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phoneNumber := vars["phone"]
	campaignID := r.URL.Query().Get("campaign_id")

	if phoneNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Phone number is required")
		return
	}

	var query string
	var args []interface{}

	if campaignID != "" {
		query = "DELETE FROM vicidial_dnc WHERE phone_number = ? AND campaign_id = ?"
		args = []interface{}{phoneNumber, campaignID}
	} else {
		query = "DELETE FROM vicidial_dnc WHERE phone_number = ?"
		args = []interface{}{phoneNumber}
	}

	result, err := h.DB.Exec(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete DNC entry: "+err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	respondWithSuccess(w, "DNC entry deleted", map[string]int64{
		"rows_deleted": rowsAffected,
	})
}

// AddFPGPhone adds a phone number to a filter phone group
func (h *Handler) AddFPGPhone(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PhoneNumber    string `json:"phone_number"`
		FilterPhoneGroupID string `json:"filter_phone_group_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.PhoneNumber == "" || req.FilterPhoneGroupID == "" {
		respondWithError(w, http.StatusBadRequest, "Phone number and filter group ID are required")
		return
	}

	query := `
		INSERT IGNORE INTO vicidial_filter_phone_groups (phone_number, filter_phone_group_id, entry_date)
		VALUES (?, ?, NOW())
	`

	result, err := h.DB.Exec(query, req.PhoneNumber, req.FilterPhoneGroupID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add filter group entry: "+err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respondWithSuccess(w, "Phone number already in filter group", nil)
		return
	}

	respondWithSuccess(w, "Phone number added to filter group", map[string]string{
		"phone_number":           req.PhoneNumber,
		"filter_phone_group_id": req.FilterPhoneGroupID,
	})
}

// DeleteFPGPhone removes a phone number from a filter phone group
func (h *Handler) DeleteFPGPhone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phoneNumber := vars["phone"]
	filterGroupID := r.URL.Query().Get("filter_phone_group_id")

	if phoneNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Phone number is required")
		return
	}

	var query string
	var args []interface{}

	if filterGroupID != "" {
		query = "DELETE FROM vicidial_filter_phone_groups WHERE phone_number = ? AND filter_phone_group_id = ?"
		args = []interface{}{phoneNumber, filterGroupID}
	} else {
		query = "DELETE FROM vicidial_filter_phone_groups WHERE phone_number = ?"
		args = []interface{}{phoneNumber}
	}

	result, err := h.DB.Exec(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete filter group entry: "+err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	respondWithSuccess(w, "Filter group entry deleted", map[string]int64{
		"rows_deleted": rowsAffected,
	})
}
