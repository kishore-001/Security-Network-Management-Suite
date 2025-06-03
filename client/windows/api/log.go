package api

import (
	"net/http"
	log "windows/logic/logs"
)

func RegisterLogRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/log", log.HandleAllSystemLogs)
}
