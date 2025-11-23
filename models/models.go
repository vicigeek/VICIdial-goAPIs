package models

import "time"

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Lead represents a lead in the system
type Lead struct {
	LeadID              int       `json:"lead_id"`
	ListID              int       `json:"list_id"`
	PhoneNumber         string    `json:"phone_number"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	MiddleInitial       string    `json:"middle_initial"`
	Address1            string    `json:"address1"`
	Address2            string    `json:"address2"`
	Address3            string    `json:"address3"`
	City                string    `json:"city"`
	State               string    `json:"state"`
	Province            string    `json:"province"`
	PostalCode          string    `json:"postal_code"`
	CountryCode         string    `json:"country_code"`
	Gender              string    `json:"gender"`
	DateOfBirth         string    `json:"date_of_birth"`
	AltPhone            string    `json:"alt_phone"`
	Email               string    `json:"email"`
	Security            string    `json:"security"`
	Comments            string    `json:"comments"`
	Status              string    `json:"status"`
	EntryDate           time.Time `json:"entry_date"`
	ModifyDate          time.Time `json:"modify_date"`
	LastLocalCallTime   time.Time `json:"last_local_call_time"`
	CalledSinceLastRest string    `json:"called_since_last_reset"`
	CalledCount         int       `json:"called_count"`
	Rank                int       `json:"rank"`
	Owner               string    `json:"owner"`
}

// List represents a call list
type List struct {
	ListID          int       `json:"list_id"`
	ListName        string    `json:"list_name"`
	CampaignID      string    `json:"campaign_id"`
	Active          string    `json:"active"`
	ListDescription string    `json:"list_description"`
	Script          string    `json:"script"`
	WebForm         string    `json:"web_form"`
	ExpDate         time.Time `json:"exp_date"`
	ListChangeUser  string    `json:"list_changeuser"`
	ResetTime       string    `json:"reset_time"`
}

// User represents a VICIdial user
type User struct {
	UserID            string    `json:"user_id"`
	User              string    `json:"user"`
	Pass              string    `json:"pass,omitempty"`
	FullName          string    `json:"full_name"`
	UserLevel         int       `json:"user_level"`
	UserGroup         string    `json:"user_group"`
	PhoneLogin        string    `json:"phone_login"`
	PhonePass         string    `json:"phone_pass"`
	Active            string    `json:"active"`
	Email             string    `json:"email"`
	CustomOne         string    `json:"custom_one"`
	CustomTwo         string    `json:"custom_two"`
	CustomThree       string    `json:"custom_three"`
	CustomFour        string    `json:"custom_four"`
	CustomFive        string    `json:"custom_five"`
	VoicemailID       string    `json:"voicemail_id"`
	AgentChooseInGroups string  `json:"agent_choose_ingroups"`
	AgentChooseBlended  string  `json:"agent_choose_blended"`
	CloserDefaultBlended string `json:"closer_default_blended"`
	UserStart         time.Time `json:"user_start"`
}

// Campaign represents a campaign
type Campaign struct {
	CampaignID          string    `json:"campaign_id"`
	CampaignName        string    `json:"campaign_name"`
	Active              string    `json:"active"`
	DialStatus          string    `json:"dial_status"`
	LeadOrder           string    `json:"lead_order"`
	DialMethod          string    `json:"dial_method"`
	AutoDialLevel       string    `json:"auto_dial_level"`
	LocalCallTime       string    `json:"local_call_time"`
	DialPrefix          string    `json:"dial_prefix"`
	CampaignCIDOverride string    `json:"campaign_cid_override"`
	ManualDialPrefix    string    `json:"manual_dial_prefix"`
	RecordingTransfer   string    `json:"recording_transfer"`
	Script              string    `json:"script"`
	GetCallLaunch       string    `json:"get_call_launch"`
	AllowClosers        string    `json:"allow_closers"`
}

// Phone represents a phone extension
type Phone struct {
	Extension      string    `json:"extension"`
	Dialplan       string    `json:"dialplan_number"`
	VoicemailExt   string    `json:"voicemail_ext"`
	PhoneIP        string    `json:"phone_ip"`
	ComputerIP     string    `json:"computer_ip"`
	ServerIP       string    `json:"server_ip"`
	Login          string    `json:"login"`
	Pass           string    `json:"pass"`
	Status         string    `json:"status"`
	Active         string    `json:"active"`
	PhoneType      string    `json:"phone_type"`
	FullName       string    `json:"fullname"`
	CompanyName    string    `json:"company"`
	PictureURL     string    `json:"picture"`
	Messages       int       `json:"messages"`
	OutboundCID    string    `json:"outbound_cid"`
	Template       string    `json:"template_id"`
	LastUpdateTime time.Time `json:"last_update_time"`
}

// DID represents a Direct Inward Dial number
type DID struct {
	DIDID              string    `json:"did_id"`
	DIDPattern         string    `json:"did_pattern"`
	DIDDescription     string    `json:"did_description"`
	DIDRoute           string    `json:"did_route"`
	RecordCall         string    `json:"record_call"`
	Extension          string    `json:"extension"`
	Exten              string    `json:"exten"`
	VoicemailExt       string    `json:"voicemail_ext"`
	FilterInboundGroup string    `json:"filter_inbound_group"`
	Group              string    `json:"group"`
	User               string    `json:"user"`
	UserDirect         string    `json:"user_direct"`
	CallMenu           string    `json:"call_menu"`
	Active             string    `json:"active"`
	EntryDate          time.Time `json:"entry_date"`
}

// AgentStatus represents the current status of an agent
type AgentStatus struct {
	User          string    `json:"user"`
	Status        string    `json:"status"`
	ServerIP      string    `json:"server_ip"`
	LoginTime     time.Time `json:"login_time"`
	CampaignID    string    `json:"campaign_id"`
	Extension     string    `json:"extension"`
	Calls         int       `json:"calls"`
	LastCallTime  time.Time `json:"last_call_time"`
	LastStatus    string    `json:"last_status"`
	PauseCode     string    `json:"pause_code"`
}

// CallLog represents a call log entry
type CallLog struct {
	UniqueID      string    `json:"uniqueid"`
	LeadID        int       `json:"lead_id"`
	ListID        int       `json:"list_id"`
	CampaignID    string    `json:"campaign_id"`
	CallDate      time.Time `json:"call_date"`
	StartEpoch    int64     `json:"start_epoch"`
	EndEpoch      int64     `json:"end_epoch"`
	Length        int       `json:"length"`
	Status        string    `json:"status"`
	PhoneCode     string    `json:"phone_code"`
	PhoneNumber   string    `json:"phone_number"`
	User          string    `json:"user"`
	Comments      string    `json:"comments"`
	ProcessedFlag string    `json:"processed"`
	UserGroup     string    `json:"user_group"`
	Term          string    `json:"term_reason"`
	AltDial       string    `json:"alt_dial"`
}

// Recording represents a call recording
type Recording struct {
	RecordingID   int       `json:"recording_id"`
	Channel       string    `json:"channel"`
	ServerIP      string    `json:"server_ip"`
	Extension     string    `json:"extension"`
	StartTime     time.Time `json:"start_time"`
	StartEpoch    int64     `json:"start_epoch"`
	EndTime       time.Time `json:"end_time"`
	EndEpoch      int64     `json:"end_epoch"`
	Length        int       `json:"length_in_sec"`
	Filename      string    `json:"filename"`
	Location      string    `json:"location"`
	LeadID        int       `json:"lead_id"`
	User          string    `json:"user"`
	VicidialID    string    `json:"vicidial_id"`
}

// Hopper represents a campaign hopper entry
type Hopper struct {
	HopperID   int       `json:"hopper_id"`
	LeadID     int       `json:"lead_id"`
	CampaignID string    `json:"campaign_id"`
	Status     string    `json:"status"`
	User       string    `json:"user"`
	ListID     int       `json:"list_id"`
	GmtOffsetNow string  `json:"gmt_offset_now"`
	State      string    `json:"state"`
	Priority   int       `json:"priority"`
	Source     string    `json:"source"`
}

// InboundGroup represents an inbound call group
type InboundGroup struct {
	GroupID              string `json:"group_id"`
	GroupName            string `json:"group_name"`
	GroupColor           string `json:"group_color"`
	Active               string `json:"active"`
	WebForm              string `json:"web_form"`
	VoicemailExt         string `json:"voicemail_ext"`
	CallsWaiting         int    `json:"calls_waiting"`
	AgentsLoggedIn       int    `json:"agents_logged_in"`
	AgentsAvailable      int    `json:"agents_available"`
	LongestWaitTime      int    `json:"longest_wait_time"`
	LongestWaitCallID    string `json:"longest_wait_call_id"`
}

// DNCEntry represents a Do Not Call entry
type DNCEntry struct {
	PhoneNumber string    `json:"phone_number"`
	CampaignID  string    `json:"campaign_id"`
	EntryDate   time.Time `json:"entry_date"`
}

// SystemVersion represents system version information
type SystemVersion struct {
	Version  string `json:"version"`
	Build    string `json:"build"`
	Timezone string `json:"timezone"`
	Date     string `json:"date"`
}
