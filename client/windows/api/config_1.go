package api

import (
	"net/http"
	"windows/logic/config_1"
)

func RegisterConfig1Routes(mux *http.ServeMux) {
	mux.HandleFunc("/client/config-1/ssh", config_1.HandleSSHUpload)
	mux.HandleFunc("/client/config-1/pass", config_1.HandlePasswordChange)
	mux.HandleFunc("/client/config-1/basic", config_1.HandleBasicInfo)
	mux.HandleFunc("/client/config-1/cmd", config_1.HandleCommandExec)
}
