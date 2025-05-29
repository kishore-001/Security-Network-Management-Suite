package api

import (
	"linux/logic/config_1"
	"net/http"
)

func RegisterConfig1Routes(mux *http.ServeMux) {
	mux.HandleFunc("/client/config-1/ssh", config_1.HandleSSHUpload)
	mux.HandleFunc("/client/config-1/pass", config_1.HandlePasswordChange)
	mux.HandleFunc("/client/config-1/basic", config_1.HandleBasicInfo)
	mux.HandleFunc("/client/config-1/cmd", config_1.HandleCommandExec)
}
