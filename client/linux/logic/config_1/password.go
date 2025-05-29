package config_1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"os/user"
	"strings"
)

type PasswordChange struct {
	Username    string `json:"username"` // Added username field
	NewPassword string `json:"new"`
}

// Secure password change function using chpasswd
func changePassword(username, newPassword string) error {
	// Check if running as root
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	if currentUser.Uid != "0" {
		return fmt.Errorf("API must be running as root/sudo to change passwords")
	}

	// Use chpasswd to change password securely
	cmd := exec.Command("chpasswd")
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s:%s", username, newPassword))
	return cmd.Run()
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

	// Validate input
	if p.Username == "" || p.NewPassword == "" {
		http.Error(w, "Username and new password are required", http.StatusBadRequest)
		return
	}

	// Use the secure password change function
	if err := changePassword(p.Username, p.NewPassword); err != nil {
		http.Error(w, fmt.Sprintf("Password change failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Password changed successfully for user: %s", p.Username)))
}
