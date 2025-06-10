package server

import (
	generaldb "backend/db/gen/server"
	"backend/logic/server/log"
	"net/http"
)

func RegisterLog(mux *http.ServeMux, queries *generaldb.Queries) {
	// GET-like operations via POST
	mux.HandleFunc("/api/admin/server/log", log.GetLog(queries))
}
