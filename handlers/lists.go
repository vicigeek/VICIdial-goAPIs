package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vicidb/non-agent-api/models"
)

// AddList creates a new list
func (h *Handler) AddList(w http.ResponseWriter, r *http.Request) {
	var list models.List
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if list.ListName == "" {
		respondWithError(w, http.StatusBadRequest, "List name is required")
		return
	}

	if list.Active == "" {
		list.Active = "Y"
	}

	query := `
		INSERT INTO vicidial_lists (list_name, campaign_id, active, list_description, script, web_form)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := h.DB.Exec(query, list.ListName, list.CampaignID, list.Active, list.ListDescription, list.Script, list.WebForm)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create list: "+err.Error())
		return
	}

	listID, _ := result.LastInsertId()
	list.ListID = int(listID)

	respondWithSuccess(w, "List created successfully", list)
}

// UpdateList updates an existing list
func (h *Handler) UpdateList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["list_id"]

	var list models.List
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_lists SET
			list_name = ?, campaign_id = ?, active = ?,
			list_description = ?, script = ?, web_form = ?
		WHERE list_id = ?
	`

	_, err := h.DB.Exec(query, list.ListName, list.CampaignID, list.Active,
		list.ListDescription, list.Script, list.WebForm, listID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update list: "+err.Error())
		return
	}

	respondWithSuccess(w, "List updated successfully", map[string]string{"list_id": listID})
}

// ListInfo retrieves list information
func (h *Handler) ListInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["list_id"]

	query := `
		SELECT list_id, list_name, campaign_id, active, list_description, script, web_form
		FROM vicidial_lists WHERE list_id = ?
	`

	var list models.List
	err := h.DB.QueryRow(query, listID).Scan(
		&list.ListID, &list.ListName, &list.CampaignID, &list.Active,
		&list.ListDescription, &list.Script, &list.WebForm,
	)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "List not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve list: "+err.Error())
		return
	}

	// Get lead count
	var leadCount int
	h.DB.QueryRow("SELECT COUNT(*) FROM vicidial_list WHERE list_id = ?", listID).Scan(&leadCount)

	respondWithSuccess(w, "List retrieved", map[string]interface{}{
		"list":       list,
		"lead_count": leadCount,
	})
}

// ListCustomFields manages custom fields for a list
func (h *Handler) ListCustomFields(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID := vars["list_id"]

	switch r.Method {
	case "GET":
		h.getListCustomFields(w, listID)
	case "POST":
		h.addListCustomField(w, r, listID)
	case "PUT":
		h.updateListCustomField(w, r, listID)
	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *Handler) getListCustomFields(w http.ResponseWriter, listID string) {
	query := `
		SELECT field_id, field_label, field_name, field_type, field_options, field_size, field_max, field_default, field_required
		FROM vicidial_lists_fields WHERE list_id = ? ORDER BY field_rank
	`

	rows, err := h.DB.Query(query, listID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve fields: "+err.Error())
		return
	}
	defer rows.Close()

	type CustomField struct {
		FieldID       int    `json:"field_id"`
		FieldLabel    string `json:"field_label"`
		FieldName     string `json:"field_name"`
		FieldType     string `json:"field_type"`
		FieldOptions  string `json:"field_options"`
		FieldSize     int    `json:"field_size"`
		FieldMax      int    `json:"field_max"`
		FieldDefault  string `json:"field_default"`
		FieldRequired string `json:"field_required"`
	}

	fields := []CustomField{}
	for rows.Next() {
		var field CustomField
		rows.Scan(&field.FieldID, &field.FieldLabel, &field.FieldName, &field.FieldType,
			&field.FieldOptions, &field.FieldSize, &field.FieldMax, &field.FieldDefault, &field.FieldRequired)
		fields = append(fields, field)
	}

	respondWithSuccess(w, "Custom fields retrieved", fields)
}

func (h *Handler) addListCustomField(w http.ResponseWriter, r *http.Request, listID string) {
	type FieldRequest struct {
		FieldLabel    string `json:"field_label"`
		FieldName     string `json:"field_name"`
		FieldType     string `json:"field_type"`
		FieldOptions  string `json:"field_options"`
		FieldSize     int    `json:"field_size"`
		FieldMax      int    `json:"field_max"`
		FieldDefault  string `json:"field_default"`
		FieldRequired string `json:"field_required"`
	}

	var field FieldRequest
	if err := json.NewDecoder(r.Body).Decode(&field); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		INSERT INTO vicidial_lists_fields (list_id, field_label, field_name, field_type, field_options, field_size, field_max, field_default, field_required)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := h.DB.Exec(query, listID, field.FieldLabel, field.FieldName, field.FieldType,
		field.FieldOptions, field.FieldSize, field.FieldMax, field.FieldDefault, field.FieldRequired)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add field: "+err.Error())
		return
	}

	fieldID, _ := result.LastInsertId()
	respondWithSuccess(w, "Field added successfully", map[string]int64{"field_id": fieldID})
}

func (h *Handler) updateListCustomField(w http.ResponseWriter, r *http.Request, listID string) {
	type FieldUpdate struct {
		FieldID       int    `json:"field_id"`
		FieldLabel    string `json:"field_label"`
		FieldType     string `json:"field_type"`
		FieldOptions  string `json:"field_options"`
		FieldSize     int    `json:"field_size"`
		FieldMax      int    `json:"field_max"`
		FieldDefault  string `json:"field_default"`
		FieldRequired string `json:"field_required"`
	}

	var field FieldUpdate
	if err := json.NewDecoder(r.Body).Decode(&field); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_lists_fields SET
			field_label = ?, field_type = ?, field_options = ?,
			field_size = ?, field_max = ?, field_default = ?, field_required = ?
		WHERE field_id = ? AND list_id = ?
	`

	_, err := h.DB.Exec(query, field.FieldLabel, field.FieldType, field.FieldOptions,
		field.FieldSize, field.FieldMax, field.FieldDefault, field.FieldRequired,
		field.FieldID, listID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update field: "+err.Error())
		return
	}

	respondWithSuccess(w, "Field updated successfully", nil)
}
