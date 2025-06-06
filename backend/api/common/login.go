package common

import (
	"backend/auth"
	generaldb "backend/db/gen/general"
	"net/http"
)

func RegisterLoginRoute(mux *http.ServeMux, queries *generaldb.Queries) {
	mux.HandleFunc("/api/login", auth.HandleLogin(queries))
}
