package alert

import (
	"encoding/json"
	"net/http"

	"backend/config"
	serverdb "backend/db/gen/server"
)

type AlertListRequest struct {
	Host  string `json:"host,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type AlertListResponse struct {
	Status string           `json:"status"`
	Alerts []serverdb.Alert `json:"alerts"`
	Count  int              `json:"count"`
}

func HandleGetAlerts(queries *serverdb.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Check user authorization
		_, ok := config.GetUserFromContext(r)
		if !ok {
			http.Error(w, "User context not found", http.StatusInternalServerError)
			return
		}

		// Parse JSON request body
		var req AlertListRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Set default limit if not provided
		if req.Limit <= 0 {
			req.Limit = 100 // default limit
		}

		var alerts []serverdb.Alert
		var err error

		if req.Host != "" {
			// Get alerts for specific host
			alerts, err = queries.GetAlertsByHost(r.Context(), req.Host)
		} else {
			// Get all recent alerts
			alerts, err = queries.GetAllAlerts(r.Context(), int32(req.Limit))
		}

		if err != nil {
			http.Error(w, "Failed to fetch alerts", http.StatusInternalServerError)
			return
		}

		response := AlertListResponse{
			Status: "success",
			Alerts: alerts,
			Count:  len(alerts),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
