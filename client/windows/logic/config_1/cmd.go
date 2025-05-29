package config_1

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

type CmdRequest struct {
	Command string `json:"command"`
}

func HandleCommandExec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CmdRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Execute the command using Windows shell
	cmd := exec.Command("cmd", "/C", req.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Command execution failed: "+strings.TrimSpace(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(out)
}
