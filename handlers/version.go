package handlers

import (
	"net/http"
	"time"

	"github.com/vicidb/non-agent-api/models"
)

// GetVersion returns API version information
func (h *Handler) GetVersion(w http.ResponseWriter, r *http.Request) {
	version := models.SystemVersion{
		Version:  "1.0.0",
		Build:    "20250108",
		Timezone: h.Config.Timezone,
		Date:     time.Now().Format("2006-01-02 15:04:05"),
	}

	respondWithSuccess(w, "Version information retrieved", version)
}
