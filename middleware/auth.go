package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/vicidb/non-agent-api/config"
)

type contextKey string

const (
	userContextKey contextKey = "user"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// AuthenticationMiddleware validates requests using a shared API key from the environment.
func AuthenticationMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if cfg.APIKey == "" {
				respondWithError(w, http.StatusInternalServerError, "Authentication not configured", "Missing API_KEY in environment")
				return
			}

			// Accept API key from header or query/form for flexibility
			providedKey := r.Header.Get("X-API-Key")
			if providedKey == "" {
				providedKey = r.URL.Query().Get("api_key")
			}
			if providedKey == "" {
				providedKey = r.FormValue("api_key")
			}

			if providedKey == "" {
				respondWithError(w, http.StatusUnauthorized, "Authentication required", "Missing API key (use header X-API-Key or query param api_key)")
				return
			}

			if providedKey != cfg.APIKey {
				respondWithError(w, http.StatusUnauthorized, "Authentication failed", "Invalid API key")
				return
			}

			// Optional user hint for downstream logging if provided
			user := r.Header.Get("X-User")
			if user == "" {
				user = r.URL.Query().Get("user")
			}
			if user == "" {
				user = r.FormValue("user")
			}
			if user == "" {
				user = "api-key"
			}

			// Add user to context for handlers to use
			ctx := context.WithValue(r.Context(), userContextKey, user)

			// Call next handler with updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserFromContext retrieves the user from request context
func GetUserFromContext(ctx context.Context) string {
	if user, ok := ctx.Value(userContextKey).(string); ok {
		return user
	}
	return ""
}

// respondWithError sends an error response
func respondWithError(w http.ResponseWriter, code int, error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   error,
		Message: message,
	})
}
