package api

import (
	"linux/logic/log"
	"net/http"
)

func RegisterLogRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/log", log.HandleLog)
}
