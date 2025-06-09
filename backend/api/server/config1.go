package server

import (
	generaldb "backend/db/gen/server"
	"backend/logic/server/config1"
	"net/http"
)

func RegisterConfig1Routes(mux *http.ServeMux, queries *generaldb.Queries) {
	mux.HandleFunc("/api/admin/server/config1/basic", config1.HanldeBasic())
}
