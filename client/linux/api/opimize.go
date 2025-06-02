package api

import (
	"linux/logic/optimization"
	"net/http"
)

func RegisterOptimizeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/optimize", optimization.HandleFileClean)
	mux.HandleFunc("/client/service", optimization.HandleListService)
	mux.HandleFunc("/client/restartservice", optimization.HandleRestartService)
}
