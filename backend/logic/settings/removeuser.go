package settings

import (
	generaldb "backend/db/gen/general"
	"encoding/json"
	"net/http"
)

func HandleRemoveUser(queries *generaldb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name         string `json:"name"`
			PasswordHash string `json:"password_hash"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err := queries.DeleteUserByEmail(r.Context(), generaldb.DeleteUserByEmailParams{
			Name:         req.Name,
			PasswordHash: req.PasswordHash,
		})
		if err != nil {
			http.Error(w, "Failed to remove user. Check name or password.", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User removed successfully"))
	}
}
