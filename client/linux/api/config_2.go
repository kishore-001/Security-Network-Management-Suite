package api

import (
	"linux/logic/config_2"
	"net/http"
)

func RegisterConfig2Routes(mux *http.ServeMux) {
	mux.HandleFunc("/client/config2/network", config_2.HandleNetworkConfig)
	mux.HandleFunc("/client/config2/route", config_2.HandleRouteTable)
	mux.HandleFunc("/client/config2/firewall", config_2.HandleFirewallRules)
	mux.HandleFunc("/client/config2/updatenetwork", config_2.HandleUpdateNetworkConfig)
	mux.HandleFunc("/client/config2/updateinterface", config_2.HandleUpdateInterface)
	mux.HandleFunc("/client/config2/restartinterface", config_2.HandleRestartInterfaces)
}
