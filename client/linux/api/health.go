package api

import (
	"linux/logic/health"
	"net/http"
)

func RegisterHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/health", health.HandleHealthConfig)
}
