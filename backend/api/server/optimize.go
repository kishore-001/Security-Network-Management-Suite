package server

import (
	generaldb "backend/db/gen/server"
	"backend/logic/server/optimisation"
	"net/http"
)

func RegisterOptimisation(mux *http.ServeMux, queries *generaldb.Queries) {
	// GET-like operations via POST
	mux.HandleFunc("/api/admin/server/cleaninfo", optimisation.GetCleanInfo(queries))
	mux.HandleFunc("/api/admin/server/optimize", optimisation.PostClean(queries))
	mux.HandleFunc("/api/admin/server/service", optimisation.GetServices(queries))
	mux.HandleFunc("/api/admin/server/restartservice", optimisation.PostRestartService(queries))
}
