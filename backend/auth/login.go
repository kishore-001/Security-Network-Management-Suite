package auth

import (
	db "backend/db/gen/general" // replace with your actual module path
	"context"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token,omitempty"`
}

func respondUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(loginResponse{Status: "unauthorized"})
}

func HandleLogin(dbQueries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req loginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		user, err := dbQueries.GetUserByName(context.Background(), req.Username)
		if err == sql.ErrNoRows {
			respondUnauthorized(w)
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			respondUnauthorized(w)
			return
		}

		// Generate JWT token with username and role
		token, err := GenerateJWTToken(req.Username, user.Role)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		resp := loginResponse{
			Status: "ok",
			Token:  token,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
