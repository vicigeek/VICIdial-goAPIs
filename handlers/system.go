package handlers

import (
	"encoding/json"
	"net/http"
)

// SoundsList retrieves available sound files
func (h *Handler) SoundsList(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT audio_filename, audio_name, user, active
		FROM vicidial_audio_store
		ORDER BY audio_name
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve sounds: "+err.Error())
		return
	}
	defer rows.Close()

	type Sound struct {
		AudioFilename string `json:"audio_filename"`
		AudioName     string `json:"audio_name"`
		User          string `json:"user"`
		Active        string `json:"active"`
	}

	sounds := []Sound{}
	for rows.Next() {
		var sound Sound
		rows.Scan(&sound.AudioFilename, &sound.AudioName, &sound.User, &sound.Active)
		sounds = append(sounds, sound)
	}

	respondWithSuccess(w, "Sound files retrieved", sounds)
}

// MOHList retrieves music on hold files
func (h *Handler) MOHList(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT moh_id, moh_name, moh_format, random
		FROM vicidial_music_on_hold
		ORDER BY moh_name
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve MOH: "+err.Error())
		return
	}
	defer rows.Close()

	type MOH struct {
		MOHID     string `json:"moh_id"`
		MOHName   string `json:"moh_name"`
		MOHFormat string `json:"moh_format"`
		Random    string `json:"random"`
	}

	mohList := []MOH{}
	for rows.Next() {
		var moh MOH
		rows.Scan(&moh.MOHID, &moh.MOHName, &moh.MOHFormat, &moh.Random)
		mohList = append(mohList, moh)
	}

	respondWithSuccess(w, "Music on hold files retrieved", mohList)
}

// VMList retrieves voicemail files
func (h *Handler) VMList(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT voicemail_id, fullname, email, active
		FROM vicidial_voicemail
		ORDER BY voicemail_id
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve voicemail: "+err.Error())
		return
	}
	defer rows.Close()

	type Voicemail struct {
		VoicemailID string `json:"voicemail_id"`
		FullName    string `json:"fullname"`
		Email       string `json:"email"`
		Active      string `json:"active"`
	}

	voicemails := []Voicemail{}
	for rows.Next() {
		var vm Voicemail
		rows.Scan(&vm.VoicemailID, &vm.FullName, &vm.Email, &vm.Active)
		voicemails = append(voicemails, vm)
	}

	respondWithSuccess(w, "Voicemail boxes retrieved", voicemails)
}

// IngroupList retrieves inbound groups
func (h *Handler) IngroupList(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("active")

	query := "SELECT group_id, group_name, group_color, active, web_form FROM vicidial_inbound_groups"
	args := []interface{}{}

	if active != "" {
		query += " WHERE active = ?"
		args = append(args, active)
	}

	query += " ORDER BY group_name"

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve inbound groups: "+err.Error())
		return
	}
	defer rows.Close()

	type InboundGroup struct {
		GroupID    string `json:"group_id"`
		GroupName  string `json:"group_name"`
		GroupColor string `json:"group_color"`
		Active     string `json:"active"`
		WebForm    string `json:"web_form"`
	}

	ingroups := []InboundGroup{}
	for rows.Next() {
		var ig InboundGroup
		rows.Scan(&ig.GroupID, &ig.GroupName, &ig.GroupColor, &ig.Active, &ig.WebForm)
		ingroups = append(ingroups, ig)
	}

	respondWithSuccess(w, "Inbound groups retrieved", ingroups)
}

// InGroupStatus retrieves real-time inbound group status
func (h *Handler) InGroupStatus(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT ig.group_id, ig.group_name,
			   COUNT(DISTINCT iga.user) as agents_logged_in,
			   COUNT(DISTINCT CASE WHEN la.status = 'READY' THEN la.user END) as agents_ready,
			   COUNT(DISTINCT vic.uniqueid) as calls_waiting
		FROM vicidial_inbound_groups ig
		LEFT JOIN vicidial_inbound_group_agents iga ON ig.group_id = iga.group_id
		LEFT JOIN vicidial_live_agents la ON iga.user = la.user
		LEFT JOIN vicidial_auto_calls vic ON ig.group_id = vic.campaign_id AND vic.status = 'LIVE'
		WHERE ig.active = 'Y'
		GROUP BY ig.group_id, ig.group_name
		ORDER BY ig.group_name
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve ingroup status: "+err.Error())
		return
	}
	defer rows.Close()

	type IngroupStatus struct {
		GroupID        string `json:"group_id"`
		GroupName      string `json:"group_name"`
		AgentsLoggedIn int    `json:"agents_logged_in"`
		AgentsReady    int    `json:"agents_ready"`
		CallsWaiting   int    `json:"calls_waiting"`
	}

	statuses := []IngroupStatus{}
	for rows.Next() {
		var status IngroupStatus
		rows.Scan(&status.GroupID, &status.GroupName, &status.AgentsLoggedIn,
			&status.AgentsReady, &status.CallsWaiting)
		statuses = append(statuses, status)
	}

	respondWithSuccess(w, "Inbound group status retrieved", statuses)
}

// CallmenuList retrieves call menus
func (h *Handler) CallmenuList(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT menu_id, menu_name, menu_prompt, menu_timeout, menu_timeout_prompt, active
		FROM vicidial_call_menu
		ORDER BY menu_name
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve call menus: "+err.Error())
		return
	}
	defer rows.Close()

	type CallMenu struct {
		MenuID            string `json:"menu_id"`
		MenuName          string `json:"menu_name"`
		MenuPrompt        string `json:"menu_prompt"`
		MenuTimeout       int    `json:"menu_timeout"`
		MenuTimeoutPrompt string `json:"menu_timeout_prompt"`
		Active            string `json:"active"`
	}

	menus := []CallMenu{}
	for rows.Next() {
		var menu CallMenu
		rows.Scan(&menu.MenuID, &menu.MenuName, &menu.MenuPrompt,
			&menu.MenuTimeout, &menu.MenuTimeoutPrompt, &menu.Active)
		menus = append(menus, menu)
	}

	respondWithSuccess(w, "Call menus retrieved", menus)
}

// ContainerList retrieves configuration containers
func (h *Handler) ContainerList(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT container_id, container_name, container_notes, container_type, user_group
		FROM vicidial_settings_containers
		ORDER BY container_name
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve containers: "+err.Error())
		return
	}
	defer rows.Close()

	type Container struct {
		ContainerID    string `json:"container_id"`
		ContainerName  string `json:"container_name"`
		ContainerNotes string `json:"container_notes"`
		ContainerType  string `json:"container_type"`
		UserGroup      string `json:"user_group"`
	}

	containers := []Container{}
	for rows.Next() {
		var container Container
		rows.Scan(&container.ContainerID, &container.ContainerName,
			&container.ContainerNotes, &container.ContainerType, &container.UserGroup)
		containers = append(containers, container)
	}

	respondWithSuccess(w, "Containers retrieved", containers)
}

// ServerRefresh triggers a server configuration refresh
func (h *Handler) ServerRefresh(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ServerIP string `json:"server_ip"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Insert refresh request into server_updater table
	query := `
		INSERT INTO vicidial_server_refresh (server_ip, last_refresh, db_time)
		VALUES (?, NOW(), NOW())
		ON DUPLICATE KEY UPDATE last_refresh = NOW(), db_time = NOW()
	`

	_, err := h.DB.Exec(query, req.ServerIP)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to trigger refresh: "+err.Error())
		return
	}

	respondWithSuccess(w, "Server refresh triggered", map[string]string{
		"server_ip": req.ServerIP,
		"status":    "refresh_requested",
	})
}

// UserGroupStatus retrieves user group status
func (h *Handler) UserGroupStatus(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT ug.user_group, ug.group_name,
			   COUNT(DISTINCT u.user) as total_users,
			   COUNT(DISTINCT la.user) as logged_in_users
		FROM vicidial_user_groups ug
		LEFT JOIN vicidial_users u ON ug.user_group = u.user_group
		LEFT JOIN vicidial_live_agents la ON u.user = la.user
		GROUP BY ug.user_group, ug.group_name
		ORDER BY ug.group_name
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve user group status: "+err.Error())
		return
	}
	defer rows.Close()

	type UserGroupStatus struct {
		UserGroup      string `json:"user_group"`
		GroupName      string `json:"group_name"`
		TotalUsers     int    `json:"total_users"`
		LoggedInUsers  int    `json:"logged_in_users"`
	}

	statuses := []UserGroupStatus{}
	for rows.Next() {
		var status UserGroupStatus
		rows.Scan(&status.UserGroup, &status.GroupName, &status.TotalUsers, &status.LoggedInUsers)
		statuses = append(statuses, status)
	}

	respondWithSuccess(w, "User group status retrieved", statuses)
}
