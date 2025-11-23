package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/vicidb/non-agent-api/middleware"
	"github.com/vicidb/non-agent-api/models"
)

// AddLead adds a new lead to the system
func (h *Handler) AddLead(w http.ResponseWriter, r *http.Request) {
	var lead models.Lead
	if err := json.NewDecoder(r.Body).Decode(&lead); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate required fields
	if lead.PhoneNumber == "" || lead.ListID == 0 {
		respondWithError(w, http.StatusBadRequest, "Phone number and list ID are required")
		return
	}

	// Set default values
	if lead.Status == "" {
		lead.Status = "NEW"
	}
	if lead.CountryCode == "" {
		lead.CountryCode = "1"
	}

	query := `
		INSERT INTO vicidial_list (
			list_id, phone_number, first_name, last_name, middle_initial,
			address1, address2, address3, city, state, province, postal_code,
			country_code, gender, date_of_birth, alt_phone, email, security,
			comments, status, entry_date, modify_date, rank, owner
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), ?, ?)
	`

	result, err := h.DB.Exec(query,
		lead.ListID, lead.PhoneNumber, lead.FirstName, lead.LastName, lead.MiddleInitial,
		lead.Address1, lead.Address2, lead.Address3, lead.City, lead.State, lead.Province,
		lead.PostalCode, lead.CountryCode, lead.Gender, lead.DateOfBirth, lead.AltPhone,
		lead.Email, lead.Security, lead.Comments, lead.Status, lead.Rank, lead.Owner,
	)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add lead: "+err.Error())
		return
	}

	leadID, _ := result.LastInsertId()
	lead.LeadID = int(leadID)

	respondWithSuccess(w, "Lead added successfully", lead)
}

// UpdateLead updates an existing lead
func (h *Handler) UpdateLead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leadID, err := strconv.Atoi(vars["lead_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid lead ID")
		return
	}

	var lead models.Lead
	if err := json.NewDecoder(r.Body).Decode(&lead); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_list SET
			first_name = ?, last_name = ?, middle_initial = ?,
			address1 = ?, address2 = ?, address3 = ?,
			city = ?, state = ?, province = ?, postal_code = ?,
			country_code = ?, gender = ?, date_of_birth = ?,
			alt_phone = ?, email = ?, security = ?, comments = ?,
			status = ?, modify_date = NOW(), owner = ?
		WHERE lead_id = ?
	`

	_, err = h.DB.Exec(query,
		lead.FirstName, lead.LastName, lead.MiddleInitial,
		lead.Address1, lead.Address2, lead.Address3,
		lead.City, lead.State, lead.Province, lead.PostalCode,
		lead.CountryCode, lead.Gender, lead.DateOfBirth,
		lead.AltPhone, lead.Email, lead.Security, lead.Comments,
		lead.Status, lead.Owner, leadID,
	)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update lead: "+err.Error())
		return
	}

	respondWithSuccess(w, "Lead updated successfully", map[string]int{"lead_id": leadID})
}

// BatchUpdateLead updates multiple leads
func (h *Handler) BatchUpdateLead(w http.ResponseWriter, r *http.Request) {
	var req struct {
		LeadIDs []int  `json:"lead_ids"`
		Status  string `json:"status"`
		Owner   string `json:"owner"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if len(req.LeadIDs) == 0 {
		respondWithError(w, http.StatusBadRequest, "No lead IDs provided")
		return
	}

	// Build query with placeholders
	query := "UPDATE vicidial_list SET modify_date = NOW()"
	args := []interface{}{}

	if req.Status != "" {
		query += ", status = ?"
		args = append(args, req.Status)
	}
	if req.Owner != "" {
		query += ", owner = ?"
		args = append(args, req.Owner)
	}

	query += " WHERE lead_id IN ("
	for i, id := range req.LeadIDs {
		if i > 0 {
			query += ","
		}
		query += "?"
		args = append(args, id)
	}
	query += ")"

	result, err := h.DB.Exec(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update leads: "+err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	respondWithSuccess(w, "Leads updated successfully", map[string]int64{"updated_count": rowsAffected})
}

// LeadSearch searches for leads
func (h *Handler) LeadSearch(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phone_number")
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	email := r.URL.Query().Get("email")
	listID := r.URL.Query().Get("list_id")
	status := r.URL.Query().Get("status")

	query := "SELECT lead_id, list_id, phone_number, first_name, last_name, email, status, entry_date FROM vicidial_list WHERE 1=1"
	args := []interface{}{}

	if phoneNumber != "" {
		query += " AND phone_number LIKE ?"
		args = append(args, "%"+phoneNumber+"%")
	}
	if firstName != "" {
		query += " AND first_name LIKE ?"
		args = append(args, "%"+firstName+"%")
	}
	if lastName != "" {
		query += " AND last_name LIKE ?"
		args = append(args, "%"+lastName+"%")
	}
	if email != "" {
		query += " AND email LIKE ?"
		args = append(args, "%"+email+"%")
	}
	if listID != "" {
		query += " AND list_id = ?"
		args = append(args, listID)
	}
	if status != "" {
		query += " AND status = ?"
		args = append(args, status)
	}

	query += " LIMIT 100"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to search leads: "+err.Error())
		return
	}
	defer rows.Close()

	leads := []models.Lead{}
	for rows.Next() {
		var lead models.Lead
		err := rows.Scan(&lead.LeadID, &lead.ListID, &lead.PhoneNumber, &lead.FirstName, &lead.LastName, &lead.Email, &lead.Status, &lead.EntryDate)
		if err != nil {
			continue
		}
		leads = append(leads, lead)
	}

	respondWithSuccess(w, "Leads retrieved successfully", leads)
}

// LeadAllInfo retrieves all information for a lead
func (h *Handler) LeadAllInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leadID, err := strconv.Atoi(vars["lead_id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid lead ID")
		return
	}

	query := `
		SELECT lead_id, list_id, phone_number, first_name, last_name, middle_initial,
			   address1, address2, address3, city, state, province, postal_code,
			   country_code, gender, date_of_birth, alt_phone, email, security,
			   comments, status, entry_date, modify_date, called_count, rank, owner
		FROM vicidial_list WHERE lead_id = ?
	`

	var lead models.Lead
	err = h.DB.QueryRow(query, leadID).Scan(
		&lead.LeadID, &lead.ListID, &lead.PhoneNumber, &lead.FirstName, &lead.LastName,
		&lead.MiddleInitial, &lead.Address1, &lead.Address2, &lead.Address3, &lead.City,
		&lead.State, &lead.Province, &lead.PostalCode, &lead.CountryCode, &lead.Gender,
		&lead.DateOfBirth, &lead.AltPhone, &lead.Email, &lead.Security, &lead.Comments,
		&lead.Status, &lead.EntryDate, &lead.ModifyDate, &lead.CalledCount, &lead.Rank, &lead.Owner,
	)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Lead not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve lead: "+err.Error())
		return
	}

	respondWithSuccess(w, "Lead information retrieved", lead)
}

// LeadFieldInfo retrieves specific field information
func (h *Handler) LeadFieldInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leadID := vars["lead_id"]
	field := r.URL.Query().Get("field")

	if field == "" {
		respondWithError(w, http.StatusBadRequest, "Field parameter is required")
		return
	}

	// Validate field name to prevent SQL injection
	validFields := map[string]bool{
		"phone_number": true, "first_name": true, "last_name": true, "email": true,
		"status": true, "comments": true, "address1": true, "city": true, "state": true,
	}

	if !validFields[field] {
		respondWithError(w, http.StatusBadRequest, "Invalid field name")
		return
	}

	query := "SELECT " + field + " FROM vicidial_list WHERE lead_id = ?"
	var value string
	err := h.DB.QueryRow(query, leadID).Scan(&value)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Lead not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve field: "+err.Error())
		return
	}

	respondWithSuccess(w, "Field retrieved", map[string]string{field: value})
}

// LeadStatusSearch searches leads by status
func (h *Handler) LeadStatusSearch(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	listID := r.URL.Query().Get("list_id")

	if status == "" {
		respondWithError(w, http.StatusBadRequest, "Status parameter is required")
		return
	}

	query := "SELECT lead_id, list_id, phone_number, first_name, last_name, status FROM vicidial_list WHERE status = ?"
	args := []interface{}{status}

	if listID != "" {
		query += " AND list_id = ?"
		args = append(args, listID)
	}

	query += " LIMIT 100"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to search leads: "+err.Error())
		return
	}
	defer rows.Close()

	leads := []models.Lead{}
	for rows.Next() {
		var lead models.Lead
		rows.Scan(&lead.LeadID, &lead.ListID, &lead.PhoneNumber, &lead.FirstName, &lead.LastName, &lead.Status)
		leads = append(leads, lead)
	}

	respondWithSuccess(w, "Leads retrieved", leads)
}

// LeadCallbackInfo retrieves callback information for a lead
func (h *Handler) LeadCallbackInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leadID := vars["lead_id"]

	query := `
		SELECT callback_id, lead_id, list_id, campaign_id, status,
			   entry_time, callback_time, user, recipient, comments
		FROM vicidial_callbacks WHERE lead_id = ?
		ORDER BY callback_time DESC LIMIT 10
	`

	rows, err := h.DB.Query(query, leadID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve callbacks: "+err.Error())
		return
	}
	defer rows.Close()

	type Callback struct {
		CallbackID   int       `json:"callback_id"`
		LeadID       int       `json:"lead_id"`
		ListID       int       `json:"list_id"`
		CampaignID   string    `json:"campaign_id"`
		Status       string    `json:"status"`
		EntryTime    time.Time `json:"entry_time"`
		CallbackTime time.Time `json:"callback_time"`
		User         string    `json:"user"`
		Recipient    string    `json:"recipient"`
		Comments     string    `json:"comments"`
	}

	callbacks := []Callback{}
	for rows.Next() {
		var cb Callback
		rows.Scan(&cb.CallbackID, &cb.LeadID, &cb.ListID, &cb.CampaignID, &cb.Status,
			&cb.EntryTime, &cb.CallbackTime, &cb.User, &cb.Recipient, &cb.Comments)
		callbacks = append(callbacks, cb)
	}

	respondWithSuccess(w, "Callbacks retrieved", callbacks)
}

// LeadDearchive restores an archived lead
func (h *Handler) LeadDearchive(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	leadID := vars["lead_id"]

	// Check if lead exists in archive
	var archivedLead models.Lead
	query := "SELECT lead_id, list_id, phone_number FROM vicidial_list_archive WHERE lead_id = ?"
	err := h.DB.QueryRow(query, leadID).Scan(&archivedLead.LeadID, &archivedLead.ListID, &archivedLead.PhoneNumber)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Archived lead not found")
		return
	}

	// Move from archive to active table
	tx, err := h.DB.Begin()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to start transaction")
		return
	}

	// Insert into active table
	_, err = tx.Exec("INSERT INTO vicidial_list SELECT * FROM vicidial_list_archive WHERE lead_id = ?", leadID)
	if err != nil {
		tx.Rollback()
		respondWithError(w, http.StatusInternalServerError, "Failed to restore lead")
		return
	}

	// Delete from archive
	_, err = tx.Exec("DELETE FROM vicidial_list_archive WHERE lead_id = ?", leadID)
	if err != nil {
		tx.Rollback()
		respondWithError(w, http.StatusInternalServerError, "Failed to remove from archive")
		return
	}

	tx.Commit()
	respondWithSuccess(w, "Lead restored successfully", map[string]int{"lead_id": archivedLead.LeadID})
}

// CheckPhoneNumber checks if a phone number exists
func (h *Handler) CheckPhoneNumber(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phone_number")
	if phoneNumber == "" {
		respondWithError(w, http.StatusBadRequest, "Phone number is required")
		return
	}

	user := middleware.GetUserFromContext(r.Context())

	query := "SELECT COUNT(*) FROM vicidial_list WHERE phone_number = ?"
	var count int
	err := h.DB.QueryRow(query, phoneNumber).Scan(&count)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to check phone number")
		return
	}

	exists := count > 0
	respondWithSuccess(w, "Phone check complete", map[string]interface{}{
		"phone_number": phoneNumber,
		"exists":       exists,
		"count":        count,
		"checked_by":   user,
	})
}
