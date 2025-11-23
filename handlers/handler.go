package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/vicidb/non-agent-api/config"
	"github.com/vicidb/non-agent-api/models"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	DB     *sql.DB
	Config *config.Config
}

// NewHandler creates a new Handler instance
func NewHandler(db *sql.DB, cfg *config.Config) *Handler {
	return &Handler{
		DB:     db,
		Config: cfg,
	}
}

// respondWithJSON sends a JSON response
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// respondWithError sends an error response
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.APIResponse{
		Success: false,
		Error:   message,
	})
}

// respondWithSuccess sends a success response
func respondWithSuccess(w http.ResponseWriter, message string, data interface{}) {
	respondWithJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
