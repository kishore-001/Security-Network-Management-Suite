package api

import (
	"net/http"
	"windows/auth"
	"windows/logic/optimization"
)

func RegisterOptimizeRoutes(mux *http.ServeMux) {
	mux.Handle("/client/optimize", auth.TokenAuthMiddleware(http.HandlerFunc(optimization.HandleFileClean)))
	mux.Handle("/client/service", auth.TokenAuthMiddleware(http.HandlerFunc(optimization.HandleRestartableServices)))
	mux.Handle("/client/restartservice", auth.TokenAuthMiddleware(http.HandlerFunc(optimization.RestartServiceHandler)))
}
