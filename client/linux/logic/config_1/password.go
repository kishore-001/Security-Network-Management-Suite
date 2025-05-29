package config_1

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type PasswordChange struct {
	CurrentPassword string `json:"current"`
	NewPassword     string `json:"new"`
}

func HandlePasswordChange(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var p PasswordChange
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("bash", "-c", "echo -e '"+p.CurrentPassword+":"+p.NewPassword+"' | chpasswd")
	err := cmd.Run()
	if err != nil {
		http.Error(w, "Password change failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password changed successfully"))
}
