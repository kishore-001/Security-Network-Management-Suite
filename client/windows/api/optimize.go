package api

import (
	"net/http"
	"windows/logic/optimization"
)

func RegisterOptimizeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/optimize", optimization.HandleFileClean)
	mux.HandleFunc("/client/service", optimization.HandleRestartableServices)
	mux.HandleFunc("/client/restartservice", optimization.RestartServiceHandler)
}
