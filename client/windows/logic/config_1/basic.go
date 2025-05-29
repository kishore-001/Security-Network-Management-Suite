package config_1

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
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

	hostnameBytes, err := exec.Command("cmd", "/C", "hostname").Output()
	if err != nil {
		http.Error(w, "Failed to get hostname", http.StatusInternalServerError)
		return
	}
	hostname := strings.TrimSpace(string(hostnameBytes))

	timezoneBytes, err := exec.Command("powershell", "-Command", "(Get-TimeZone).Id").Output()
	if err != nil {
		http.Error(w, "Failed to get timezone", http.StatusInternalServerError)
		return
	}
	timezone := strings.TrimSpace(string(timezoneBytes))

	data := BasicInfo{
		Hostname: hostname,
		Timezone: timezone,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
