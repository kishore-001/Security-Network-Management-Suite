package api

import (
	"linux/logic/config_2"
	"net/http"
)

func RegisterConfig2Routes(mux *http.ServeMux) {
	mux.HandleFunc("/client/config2/network", config_2.HandleNetworkConfig)
}
