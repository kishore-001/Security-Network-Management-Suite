package api

import (
	"linux/logic/optimization"
	"net/http"
)

func RegisterOptimizeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/client/optimize", optimization.HandleFileClean)
}
