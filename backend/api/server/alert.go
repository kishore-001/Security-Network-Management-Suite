package server

import (
	serverdb "backend/db/gen/server"
	"backend/logic/server/alert"
	"net/http"
)

func RegisterAlertRoute(mux *http.ServeMux, queries *serverdb.Queries) {
	mux.HandleFunc("/api/server/alert", alert.HandleGetAlerts(queries))
}
