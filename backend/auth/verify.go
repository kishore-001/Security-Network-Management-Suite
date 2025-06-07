package auth

import (
	db "backend/db/gen/general"
	"encoding/json"
	"net/http"
	"strings"
)

type verifyResponse struct {
	Status   string `json:"status"`
	Username string `json:"username,omitempty"`
	Role     string `json:"role,omitempty"`
}

func HandleVerify(dbQueries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(verifyResponse{Status: "invalid"})
			return
		}

		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(verifyResponse{Status: "invalid"})
			return
		}

		// Validate token
		claims, err := ValidateAccessToken(tokenParts[1])
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(verifyResponse{Status: "invalid"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(verifyResponse{
			Status:   "valid",
			Username: claims.Username,
			Role:     claims.Role,
		})
	}
}
