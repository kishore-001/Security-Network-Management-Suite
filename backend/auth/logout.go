package auth

import (
	db "backend/db/gen/general"
	"context"
	"encoding/json"
	"net/http"
)

func HandleLogout(dbQueries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Get refresh token from cookie
		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			// No cookie, but that's okay for logout
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
			return
		}

		// Delete refresh token from database
		err = dbQueries.DeleteRefreshToken(context.Background(), cookie.Value)
		if err != nil {
			// Log error but don't fail logout
			// log.Printf("Failed to delete refresh token: %v", err)
		}

		// Clear the refresh token cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			MaxAge:   -1, // Delete the cookie
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}
}
