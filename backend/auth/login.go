package auth

import (
	db "backend/db/gen/general"
	"context"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Status      string `json:"status"`
	AccessToken string `json:"access_token,omitempty"`
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
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		accessToken, err := GenerateAccessToken(req.Username, user.Role)
		if err != nil {
			http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
			return
		}

		// 🔍 Check if user already has a valid refresh token
		existingSession, err := dbQueries.GetValidRefreshTokenByUser(context.Background(), req.Username)

		var refreshToken string

		if err == sql.ErrNoRows {
			// 🆕 No valid session exists, create new refresh token
			refreshToken, err = GenerateRefreshToken()
			if err != nil {
				http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
				return
			}

			// Save new refresh token to DB
			err = dbQueries.SaveRefreshToken(context.Background(), db.SaveRefreshTokenParams{
				Username:     req.Username,
				RefreshToken: refreshToken,
				ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
			})
			if err != nil {
				http.Error(w, "Failed to save session", http.StatusInternalServerError)
				return
			}
		} else if err != nil {
			// Database error
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		} else {
			// ♻️ Valid session exists, reuse the existing refresh token
			refreshToken = existingSession.RefreshToken
		}

		// Set refresh token as HttpOnly secure cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   7 * 24 * 60 * 60,
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(loginResponse{
			Status:      "ok",
			AccessToken: accessToken,
		})
	}
}
