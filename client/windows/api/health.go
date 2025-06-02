package api

import (
	"net/http"
	"windows/logic/health"
)

func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/health", health.HandleHealthConfig)
}
