package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vicidb/non-agent-api/models"
)

// AddPhone adds a new phone extension
func (h *Handler) AddPhone(w http.ResponseWriter, r *http.Request) {
	var phone models.Phone
	if err := json.NewDecoder(r.Body).Decode(&phone); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if phone.Extension == "" {
		respondWithError(w, http.StatusBadRequest, "Extension is required")
		return
	}

	if phone.Active == "" {
		phone.Active = "Y"
	}

	query := `
		INSERT INTO phones (extension, dialplan_number, voicemail_ext, phone_ip,
							computer_ip, server_ip, login, pass, status, active,
							phone_type, fullname, company, outbound_cid)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := h.DB.Exec(query, phone.Extension, phone.Dialplan, phone.VoicemailExt,
		phone.PhoneIP, phone.ComputerIP, phone.ServerIP, phone.Login, phone.Pass,
		phone.Status, phone.Active, phone.PhoneType, phone.FullName,
		phone.CompanyName, phone.OutboundCID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add phone: "+err.Error())
		return
	}

	respondWithSuccess(w, "Phone added successfully", map[string]string{"extension": phone.Extension})
}

// UpdatePhone updates an existing phone
func (h *Handler) UpdatePhone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phoneID := vars["phone_id"]

	var phone models.Phone
	if err := json.NewDecoder(r.Body).Decode(&phone); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE phones SET
			dialplan_number = ?, voicemail_ext = ?, phone_ip = ?,
			computer_ip = ?, server_ip = ?, status = ?, active = ?,
			phone_type = ?, fullname = ?, company = ?, outbound_cid = ?
		WHERE extension = ?
	`

	_, err := h.DB.Exec(query, phone.Dialplan, phone.VoicemailExt, phone.PhoneIP,
		phone.ComputerIP, phone.ServerIP, phone.Status, phone.Active,
		phone.PhoneType, phone.FullName, phone.CompanyName, phone.OutboundCID, phoneID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update phone: "+err.Error())
		return
	}

	respondWithSuccess(w, "Phone updated successfully", map[string]string{"extension": phoneID})
}

// AddPhoneAlias adds a phone alias
func (h *Handler) AddPhoneAlias(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AliasID    string `json:"alias_id"`
		AliasName  string `json:"alias_name"`
		Extension  string `json:"extension"`
		Active     string `json:"active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `INSERT INTO phone_aliases (alias_id, alias_name, logins_list, active) VALUES (?, ?, ?, ?)`
	_, err := h.DB.Exec(query, req.AliasID, req.AliasName, req.Extension, req.Active)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add phone alias: "+err.Error())
		return
	}

	respondWithSuccess(w, "Phone alias added successfully", map[string]string{"alias_id": req.AliasID})
}

// UpdatePhoneAlias updates a phone alias
func (h *Handler) UpdatePhoneAlias(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aliasID := vars["alias_id"]

	var req struct {
		AliasName  string `json:"alias_name"`
		Extension  string `json:"extension"`
		Active     string `json:"active"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `UPDATE phone_aliases SET alias_name = ?, logins_list = ?, active = ? WHERE alias_id = ?`
	_, err := h.DB.Exec(query, req.AliasName, req.Extension, req.Active, aliasID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update phone alias: "+err.Error())
		return
	}

	respondWithSuccess(w, "Phone alias updated successfully", map[string]string{"alias_id": aliasID})
}

// AddDID adds a new DID
func (h *Handler) AddDID(w http.ResponseWriter, r *http.Request) {
	var did models.DID
	if err := json.NewDecoder(r.Body).Decode(&did); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if did.DIDPattern == "" {
		respondWithError(w, http.StatusBadRequest, "DID pattern is required")
		return
	}

	if did.Active == "" {
		did.Active = "Y"
	}

	query := `
		INSERT INTO vicidial_inbound_dids (did_pattern, did_description, did_route,
										   record_call, extension, exten, voicemail_ext,
										   filter_inbound_group, group_id, user, active)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := h.DB.Exec(query, did.DIDPattern, did.DIDDescription, did.DIDRoute,
		did.RecordCall, did.Extension, did.Exten, did.VoicemailExt,
		did.FilterInboundGroup, did.Group, did.User, did.Active)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to add DID: "+err.Error())
		return
	}

	didID, _ := result.LastInsertId()
	respondWithSuccess(w, "DID added successfully", map[string]int64{"did_id": didID})
}

// UpdateDID updates an existing DID
func (h *Handler) UpdateDID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	didID := vars["did_id"]

	var did models.DID
	if err := json.NewDecoder(r.Body).Decode(&did); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_inbound_dids SET
			did_pattern = ?, did_description = ?, did_route = ?,
			record_call = ?, extension = ?, exten = ?, voicemail_ext = ?,
			filter_inbound_group = ?, group_id = ?, user = ?, active = ?
		WHERE did_id = ?
	`

	_, err := h.DB.Exec(query, did.DIDPattern, did.DIDDescription, did.DIDRoute,
		did.RecordCall, did.Extension, did.Exten, did.VoicemailExt,
		did.FilterInboundGroup, did.Group, did.User, did.Active, didID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update DID: "+err.Error())
		return
	}

	respondWithSuccess(w, "DID updated successfully", map[string]string{"did_id": didID})
}

// CopyDID duplicates a DID configuration
func (h *Handler) CopyDID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceDID := vars["did_id"]

	var req struct {
		NewDIDPattern string `json:"new_did_pattern"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		INSERT INTO vicidial_inbound_dids (did_pattern, did_description, did_route,
										   record_call, extension, exten, voicemail_ext,
										   filter_inbound_group, group_id, user, active)
		SELECT ?, did_description, did_route, record_call, extension, exten, voicemail_ext,
			   filter_inbound_group, group_id, user, active
		FROM vicidial_inbound_dids WHERE did_id = ?
	`

	result, err := h.DB.Exec(query, req.NewDIDPattern, sourceDID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to copy DID: "+err.Error())
		return
	}

	newDIDID, _ := result.LastInsertId()
	respondWithSuccess(w, "DID copied successfully", map[string]int64{"new_did_id": newDIDID})
}
