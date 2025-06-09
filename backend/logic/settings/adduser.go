package settings

import (
	generaldb "backend/db/gen/general"
	"encoding/json"
	"net/http"
)
func HandleAddUser(queries *generaldb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name         string `json:"name"`
			Role         string `json:"role"`
			Email        string `json:"email"`
			PasswordHash string `json:"password_hash"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := queries.CreateUser(r.Context(), generaldb.CreateUserParams{
			Name:         req.Name,
			Role:         req.Role,
			Email:        req.Email,
			PasswordHash: req.PasswordHash,
		})
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}
