package api

import (
	"net/http"
	"windows/logic/config_2"
)

func RegisterConfig1Routes(mux *http.ServeMux) {
	mux.HandleFunc("/client/config-1/basic", config_2.BasicInfo)

}
