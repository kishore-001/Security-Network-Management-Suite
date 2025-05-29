package config_1

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type BasicInfo struct {
	Hostname string `json:"hostname"`
	Timezone string `json:"timezone"`
}

func HandleBasicInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	hostname, _ := exec.Command("hostname").Output()
	timezone, _ := exec.Command("timedatectl").Output()

	data := BasicInfo{
		Hostname: string(hostname),
		Timezone: string(timezone),
	}

	json.NewEncoder(w).Encode(data)
}
