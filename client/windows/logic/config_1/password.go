package config_1

import (
	"encoding/json"
	"net/http"
	"os/exec"
)

type PasswordChange struct {
	Username        string `json:"username"`
	CurrentPassword string `json:"current"`
	NewPassword     string `json:"new"`
	ConfirmPassword string `json:"confirm"`
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

	if p.NewPassword != p.ConfirmPassword {
		http.Error(w, "New passwords do not match", http.StatusBadRequest)
		return
	}

	// NOTE: This skips old password verification and requires admin rights
	cmd := exec.Command("powershell", "-Command",
		`net user "`+p.Username+`" "`+p.NewPassword+`"`)

	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Password change failed: "+string(out), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password changed successfully"))
}
