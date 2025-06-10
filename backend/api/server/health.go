package server

import (
	generaldb "backend/db/gen/server"
	"backend/logic/server/health"
	"net/http"
)

func RegisterHealth(mux *http.ServeMux, queries *generaldb.Queries) {
	// GET-like operations via POST
	mux.HandleFunc("/api/admin/server/health", health.GetHealth(queries))
}
