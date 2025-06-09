package settings

import (
	generaldb "backend/db/gen/general"
	"encoding/json"
	"net/http"
)

func HandleRemoveMac(queries *generaldb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Mac string `json:"mac"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		err := queries.RemoveMacAccess(r.Context(), req.Mac)
		if err != nil {
			http.Error(w, "Failed to remove MAC address", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("MAC address removed successfully"))
	}
}
