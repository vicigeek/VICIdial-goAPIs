package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vicidb/non-agent-api/models"
)

// AddUser creates a new user
func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if user.User == "" || user.Pass == "" {
		respondWithError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	if user.Active == "" {
		user.Active = "Y"
	}

	query := `
		INSERT INTO vicidial_users (user, pass, full_name, user_level, user_group, phone_login, phone_pass, active, email)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := h.DB.Exec(query, user.User, user.Pass, user.FullName, user.UserLevel,
		user.UserGroup, user.PhoneLogin, user.PhonePass, user.Active, user.Email)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}

	respondWithSuccess(w, "User created successfully", map[string]string{"user": user.User})
}

// UpdateUser updates an existing user
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_users SET
			full_name = ?, user_level = ?, user_group = ?,
			phone_login = ?, phone_pass = ?, active = ?, email = ?
		WHERE user = ?
	`

	_, err := h.DB.Exec(query, user.FullName, user.UserLevel, user.UserGroup,
		user.PhoneLogin, user.PhonePass, user.Active, user.Email, userID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update user: "+err.Error())
		return
	}

	respondWithSuccess(w, "User updated successfully", map[string]string{"user": userID})
}

// CopyUser duplicates a user configuration
func (h *Handler) CopyUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceUser := vars["user_id"]

	var req struct {
		NewUser string `json:"new_user"`
		NewPass string `json:"new_pass"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Copy user settings
	query := `
		INSERT INTO vicidial_users (user, pass, full_name, user_level, user_group, phone_login, phone_pass, active, email)
		SELECT ?, ?, full_name, user_level, user_group, phone_login, phone_pass, active, email
		FROM vicidial_users WHERE user = ?
	`

	_, err := h.DB.Exec(query, req.NewUser, req.NewPass, sourceUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to copy user: "+err.Error())
		return
	}

	respondWithSuccess(w, "User copied successfully", map[string]string{"new_user": req.NewUser})
}

// UserDetails retrieves detailed user information
func (h *Handler) UserDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	query := `
		SELECT user, full_name, user_level, user_group, phone_login, active, email,
			   custom_one, custom_two, custom_three, custom_four, custom_five
		FROM vicidial_users WHERE user = ?
	`

	var user models.User
	err := h.DB.QueryRow(query, userID).Scan(
		&user.User, &user.FullName, &user.UserLevel, &user.UserGroup,
		&user.PhoneLogin, &user.Active, &user.Email, &user.CustomOne,
		&user.CustomTwo, &user.CustomThree, &user.CustomFour, &user.CustomFive,
	)

	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
		return
	}

	respondWithSuccess(w, "User details retrieved", user)
}

// LoggedInAgents retrieves currently logged-in agents
func (h *Handler) LoggedInAgents(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT user, server_ip, extension, status, campaign_id, last_update_time
		FROM vicidial_live_agents
		ORDER BY last_update_time DESC
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve agents: "+err.Error())
		return
	}
	defer rows.Close()

	type LiveAgent struct {
		User           string    `json:"user"`
		ServerIP       string    `json:"server_ip"`
		Extension      string    `json:"extension"`
		Status         string    `json:"status"`
		CampaignID     string    `json:"campaign_id"`
		LastUpdateTime time.Time `json:"last_update_time"`
	}

	agents := []LiveAgent{}
	for rows.Next() {
		var agent LiveAgent
		rows.Scan(&agent.User, &agent.ServerIP, &agent.Extension, &agent.Status,
			&agent.CampaignID, &agent.LastUpdateTime)
		agents = append(agents, agent)
	}

	respondWithSuccess(w, "Logged-in agents retrieved", agents)
}

// AgentStatus retrieves real-time agent status
func (h *Handler) AgentStatus(w http.ResponseWriter, r *http.Request) {
	campaignID := r.URL.Query().Get("campaign_id")
	_ = r.URL.Query().Get("user_group") // TODO: implement user_group filtering

	query := `
		SELECT user, status, server_ip, extension, campaign_id, last_call_time, pause_code, calls_today
		FROM vicidial_live_agents WHERE 1=1
	`
	args := []interface{}{}

	if campaignID != "" {
		query += " AND campaign_id = ?"
		args = append(args, campaignID)
	}

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve status: "+err.Error())
		return
	}
	defer rows.Close()

	type AgentStatusInfo struct {
		User         string    `json:"user"`
		Status       string    `json:"status"`
		ServerIP     string    `json:"server_ip"`
		Extension    string    `json:"extension"`
		CampaignID   string    `json:"campaign_id"`
		LastCallTime time.Time `json:"last_call_time"`
		PauseCode    string    `json:"pause_code"`
		CallsToday   int       `json:"calls_today"`
	}

	agents := []AgentStatusInfo{}
	for rows.Next() {
		var agent AgentStatusInfo
		rows.Scan(&agent.User, &agent.Status, &agent.ServerIP, &agent.Extension,
			&agent.CampaignID, &agent.LastCallTime, &agent.PauseCode, &agent.CallsToday)
		agents = append(agents, agent)
	}

	respondWithSuccess(w, "Agent status retrieved", agents)
}

// AgentIngroupInfo retrieves agent inbound group information
func (h *Handler) AgentIngroupInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	agentID := vars["agent_id"]

	query := `
		SELECT group_id, user, group_rank, group_web_vars
		FROM vicidial_inbound_group_agents
		WHERE user = ?
		ORDER BY group_rank
	`

	rows, err := h.DB.Query(query, agentID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve ingroups: "+err.Error())
		return
	}
	defer rows.Close()

	type IngroupAssignment struct {
		GroupID      string `json:"group_id"`
		User         string `json:"user"`
		GroupRank    int    `json:"group_rank"`
		GroupWebVars string `json:"group_web_vars"`
	}

	ingroups := []IngroupAssignment{}
	for rows.Next() {
		var ig IngroupAssignment
		rows.Scan(&ig.GroupID, &ig.User, &ig.GroupRank, &ig.GroupWebVars)
		ingroups = append(ingroups, ig)
	}

	respondWithSuccess(w, "Agent ingroups retrieved", ingroups)
}

// AgentCampaigns retrieves campaigns assigned to an agent
func (h *Handler) AgentCampaigns(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	agentID := vars["agent_id"]

	query := `
		SELECT campaign_id, campaign_rank
		FROM vicidial_campaign_agents
		WHERE user = ?
		ORDER BY campaign_rank
	`

	rows, err := h.DB.Query(query, agentID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve campaigns: "+err.Error())
		return
	}
	defer rows.Close()

	type CampaignAssignment struct {
		CampaignID   string `json:"campaign_id"`
		CampaignRank int    `json:"campaign_rank"`
	}

	campaigns := []CampaignAssignment{}
	for rows.Next() {
		var camp CampaignAssignment
		rows.Scan(&camp.CampaignID, &camp.CampaignRank)
		campaigns = append(campaigns, camp)
	}

	respondWithSuccess(w, "Agent campaigns retrieved", campaigns)
}

// UpdateRemoteAgent updates remote agent settings
func (h *Handler) UpdateRemoteAgent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	agentID := vars["agent_id"]

	var req struct {
		RemoteAgentID string `json:"remote_agent_id"`
		Status        string `json:"status"`
		ServerIP      string `json:"server_ip"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	query := `
		UPDATE vicidial_remote_agents SET
			status = ?, server_ip = ?
		WHERE user = ? AND remote_agent_id = ?
	`

	_, err := h.DB.Exec(query, req.Status, req.ServerIP, agentID, req.RemoteAgentID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to update remote agent: "+err.Error())
		return
	}

	respondWithSuccess(w, "Remote agent updated successfully", nil)
}
