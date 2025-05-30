package api

import (
	"linux/logic/config_1"
	"net/http"
)

func RegisterConfig1Routes(mux *http.ServeMux) {
	mux.HandleFunc("/client/config1/ssh", config_1.HandleSSHUpload)
	mux.HandleFunc("/client/config1/basic_update", config_1.HandleBasicUpdate)
	mux.HandleFunc("/client/config1/pass", config_1.HandlePasswordChange)
	mux.HandleFunc("/client/config1/basic", config_1.HandleBasicInfo)
	mux.HandleFunc("/client/config1/cmd", config_1.HandleCommandExec)
}
