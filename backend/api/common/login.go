package common

import (
	"backend/auth"
	"net/http"
)

func RegisterLoginRoute(mux *http.ServeMux) {
	mux.HandleFunc("/api/login", auth.HandleLogin)
}
