package server

import (
	generaldb "backend/db/gen/server"
	"backend/logic/server/config1"
	"net/http"
)

func RegisterConfig1Routes(mux *http.ServeMux, queries *generaldb.Queries) {
	mux.HandleFunc("/api/admin/server/config1/basic", config1.HandleBasic(queries))
	mux.HandleFunc("/api/admin/server/config1/basic_update", config1.HandleBasicChange(queries))
	mux.HandleFunc("/api/admin/server/config1/create", config1.HandleCreateServer(queries))
	mux.HandleFunc("/api/admin/server/config1/delete", config1.HandleDeleteServer(queries))
	mux.HandleFunc("/api/admin/server/config1/device", config1.HandleGetAllServers(queries))
}
