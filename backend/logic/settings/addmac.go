package settings

import (
	generaldb "backend/db/gen/general"
	"encoding/json"
	"net/http"
)

func HandleAddMac(queries *generaldb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Mac    string `json:"mac"`
			Status string `json:"status"` // Should be "BLACKLISTED" or "WHITELISTED"
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Insert MAC into DB
		mac, err := queries.AddMacAccess(r.Context(), generaldb.AddMacAccessParams{
			Mac:    req.Mac,
			Status: req.Status,
		})
		if err != nil {
			http.Error(w, "Failed to add MAC address", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(mac)
	}
}
