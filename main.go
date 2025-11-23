package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/vicidb/non-agent-api/config"
	"github.com/vicidb/non-agent-api/database"
	"github.com/vicidb/non-agent-api/handlers"
	"github.com/vicidb/non-agent-api/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Successfully connected to database")

	// Initialize router
	router := mux.NewRouter()

	// API v1 routes
	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	// Apply authentication middleware
	apiRouter.Use(middleware.AuthenticationMiddleware(cfg))
	apiRouter.Use(middleware.LoggingMiddleware)

	// Initialize handlers
	h := handlers.NewHandler(db, cfg)

	// Version endpoint
	apiRouter.HandleFunc("/version", h.GetVersion).Methods("GET")

	// Lead Management
	apiRouter.HandleFunc("/leads", h.AddLead).Methods("POST")
	apiRouter.HandleFunc("/leads/{lead_id}", h.UpdateLead).Methods("PUT")
	apiRouter.HandleFunc("/leads/batch", h.BatchUpdateLead).Methods("PUT")
	apiRouter.HandleFunc("/leads/search", h.LeadSearch).Methods("GET")
	apiRouter.HandleFunc("/leads/{lead_id}/info", h.LeadAllInfo).Methods("GET")
	apiRouter.HandleFunc("/leads/{lead_id}/field-info", h.LeadFieldInfo).Methods("GET")
	apiRouter.HandleFunc("/leads/status-search", h.LeadStatusSearch).Methods("GET")
	apiRouter.HandleFunc("/leads/{lead_id}/callback-info", h.LeadCallbackInfo).Methods("GET")
	apiRouter.HandleFunc("/leads/{lead_id}/dearchive", h.LeadDearchive).Methods("POST")
	apiRouter.HandleFunc("/phone/check", h.CheckPhoneNumber).Methods("GET")

	// List Management
	apiRouter.HandleFunc("/lists", h.AddList).Methods("POST")
	apiRouter.HandleFunc("/lists/{list_id}", h.UpdateList).Methods("PUT")
	apiRouter.HandleFunc("/lists/{list_id}/info", h.ListInfo).Methods("GET")
	apiRouter.HandleFunc("/lists/{list_id}/custom-fields", h.ListCustomFields).Methods("GET", "POST", "PUT")

	// User/Agent Management
	apiRouter.HandleFunc("/users", h.AddUser).Methods("POST")
	apiRouter.HandleFunc("/users/{user_id}", h.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/users/{user_id}/copy", h.CopyUser).Methods("POST")
	apiRouter.HandleFunc("/users/{user_id}/details", h.UserDetails).Methods("GET")
	apiRouter.HandleFunc("/users/logged-in", h.LoggedInAgents).Methods("GET")
	apiRouter.HandleFunc("/agents/status", h.AgentStatus).Methods("GET")
	apiRouter.HandleFunc("/agents/{agent_id}/ingroup-info", h.AgentIngroupInfo).Methods("GET")
	apiRouter.HandleFunc("/agents/{agent_id}/campaigns", h.AgentCampaigns).Methods("GET")
	apiRouter.HandleFunc("/remote-agents/{agent_id}", h.UpdateRemoteAgent).Methods("PUT")

	// Campaign Management
	apiRouter.HandleFunc("/campaigns/{campaign_id}", h.UpdateCampaign).Methods("PUT")
	apiRouter.HandleFunc("/campaigns", h.CampaignsList).Methods("GET")
	apiRouter.HandleFunc("/campaigns/with-lists", h.GetCampaignsWithLists).Methods("GET")
	apiRouter.HandleFunc("/campaigns/{campaign_id}/hopper", h.HopperList).Methods("GET")
	apiRouter.HandleFunc("/campaigns/{campaign_id}/hopper/bulk", h.HopperBulkInsert).Methods("POST")

	// SIP/Carrier Logs
	apiRouter.HandleFunc("/sip/carrier-log", h.GetSIPLog).Methods("GET")
	apiRouter.HandleFunc("/sip/event-log", h.GetSIPEventLog).Methods("GET")
	apiRouter.HandleFunc("/sip/live-channels", h.GetLiveSIPChannels).Methods("GET")

	// KPI & Analytics
	apiRouter.HandleFunc("/kpi/dispositions", h.GetKPIDispositions).Methods("GET")

	// Test Calls
	apiRouter.HandleFunc("/test-call/send", h.SendTestCall).Methods("POST")
	apiRouter.HandleFunc("/test-call/status", h.GetTestCallStatus).Methods("GET")
	apiRouter.HandleFunc("/test-call/list", h.ListTestCalls).Methods("GET")

	// Phone/DID Management
	apiRouter.HandleFunc("/phones", h.AddPhone).Methods("POST")
	apiRouter.HandleFunc("/phones/{phone_id}", h.UpdatePhone).Methods("PUT")
	apiRouter.HandleFunc("/phone-aliases", h.AddPhoneAlias).Methods("POST")
	apiRouter.HandleFunc("/phone-aliases/{alias_id}", h.UpdatePhoneAlias).Methods("PUT")
	apiRouter.HandleFunc("/dids", h.AddDID).Methods("POST")
	apiRouter.HandleFunc("/dids/{did_id}", h.UpdateDID).Methods("PUT")
	apiRouter.HandleFunc("/dids/{did_id}/copy", h.CopyDID).Methods("POST")

	// DNC Management
	apiRouter.HandleFunc("/dnc", h.AddDNCPhone).Methods("POST")
	apiRouter.HandleFunc("/dnc/{phone}", h.DeleteDNCPhone).Methods("DELETE")
	apiRouter.HandleFunc("/fpg", h.AddFPGPhone).Methods("POST")
	apiRouter.HandleFunc("/fpg/{phone}", h.DeleteFPGPhone).Methods("DELETE")

	// Reporting & Monitoring
	apiRouter.HandleFunc("/recordings/lookup", h.RecordingLookup).Methods("GET")
	apiRouter.HandleFunc("/did-logs/export", h.DIDLogExport).Methods("GET")
	apiRouter.HandleFunc("/phone-logs/{phone}", h.PhoneNumberLog).Methods("GET")
	apiRouter.HandleFunc("/agent-stats/export", h.AgentStatsExport).Methods("GET")
	apiRouter.HandleFunc("/call-stats/status", h.CallStatusStats).Methods("GET")
	apiRouter.HandleFunc("/call-stats/dispo", h.CallDispoReport).Methods("GET")
	apiRouter.HandleFunc("/monitor/blind", h.BlindMonitor).Methods("POST")

	// System Management
	apiRouter.HandleFunc("/system/sounds", h.SoundsList).Methods("GET")
	apiRouter.HandleFunc("/system/moh", h.MOHList).Methods("GET")
	apiRouter.HandleFunc("/system/voicemail", h.VMList).Methods("GET")
	apiRouter.HandleFunc("/ingroups", h.IngroupList).Methods("GET")
	apiRouter.HandleFunc("/ingroups/status", h.InGroupStatus).Methods("GET")
	apiRouter.HandleFunc("/callmenus", h.CallmenuList).Methods("GET")
	apiRouter.HandleFunc("/containers", h.ContainerList).Methods("GET")
	apiRouter.HandleFunc("/system/refresh", h.ServerRefresh).Methods("POST")
	apiRouter.HandleFunc("/user-groups/status", h.UserGroupStatus).Methods("GET")

	// Advanced Features
	apiRouter.HandleFunc("/group-aliases", h.AddGroupAlias).Methods("POST")
	apiRouter.HandleFunc("/log-entries/{entry_id}", h.UpdateLogEntry).Methods("PUT")
	apiRouter.HandleFunc("/cid-groups/{entry_id}", h.UpdateCIDGroupEntry).Methods("PUT")
	apiRouter.HandleFunc("/alt-urls/{url_id}", h.UpdateAltURL).Methods("PUT")
	apiRouter.HandleFunc("/presets/{preset_id}", h.UpdatePresets).Methods("PUT")
	apiRouter.HandleFunc("/calls/{call_id}/info", h.CallidInfo).Methods("GET")
	apiRouter.HandleFunc("/ccc/lead-info/{lead_id}", h.CCCLeadInfo).Methods("GET")

	// Health check endpoint (no auth required)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	// Start server
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting VICIdial Non-Agent API server on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
