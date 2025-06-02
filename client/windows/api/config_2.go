package api

import (
	"net/http"
	"windows/logic/config_2"
)

func RegisterConfig2Routes(mux *http.ServeMux) {
	mux.HandleFunc("/client/config-2/basic", config_2.NetworkConfigHandler)
	mux.HandleFunc("/client/config-2/route", config_2.RouteHandler)
	mux.HandleFunc("/client/config-2/firewall", config_2.GetWindowsFirewallRules)
	mux.HandleFunc("/client/config-2/updateinterface", config_2.HandleUpdateInterface)
	mux.HandleFunc("/client/config-2/updatenetwork", config_2.HandleUpdateNetworkConfig)
	mux.HandleFunc("/client/config-2/postrestartinterface", config_2.HandleRestartInterfaces)
}
