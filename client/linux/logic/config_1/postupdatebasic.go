package config_1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

// BasicUpdateRequest defines the expected JSON structure for basic system updates
type BasicUpdateRequest struct {
	Hostname string `json:"hostname"`
	Timezone string `json:"timezone"`
}

// HandleBasicUpdate processes requests to update hostname and timezone
func HandleBasicUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request
	var updateReq BasicUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, "Failed to parse JSON request", http.StatusBadRequest)
		return
	}

	// Response data to track what was updated
	updated := make(map[string]bool)
	messages := make(map[string]string)

	// Update hostname if provided
	if updateReq.Hostname != "" {
		if err := updateHostname(updateReq.Hostname); err != nil {
			messages["hostname"] = fmt.Sprintf("Failed to update hostname: %v", err)
			updated["hostname"] = false
		} else {
			messages["hostname"] = "Hostname updated successfully"
			updated["hostname"] = true
		}
	}

	// Update timezone if provided
	if updateReq.Timezone != "" {
		if err := updateTimezone(updateReq.Timezone); err != nil {
			messages["timezone"] = fmt.Sprintf("Failed to update timezone: %v", err)
			updated["timezone"] = false
		} else {
			messages["timezone"] = "Timezone updated successfully"
			updated["timezone"] = true
		}
	}

	// Prepare response
	response := map[string]interface{}{
		"updated":  updated,
		"messages": messages,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// updateHostname changes the system hostname
func updateHostname(hostname string) error {
	// Check if running as root/sudo
	if os.Geteuid() != 0 {
		return fmt.Errorf("must run as root/sudo to change hostname")
	}

	// Execute hostname command to change hostname immediately
	cmd := exec.Command("hostname", hostname)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set hostname temporarily: %w", err)
	}

	// Update /etc/hostname file for persistence across reboots
	if err := os.WriteFile("/etc/hostname", []byte(hostname+"\n"), 0644); err != nil {
		return fmt.Errorf("failed to update /etc/hostname: %w", err)
	}

	// Also update /etc/hosts to ensure the hostname resolves locally
	// Note: This is a simplified approach; in production you might want to
	// parse and update the file more carefully
	hostsCmd := exec.Command("sed", "-i", fmt.Sprintf("s/127.0.1.1.*/127.0.1.1\t%s/", hostname), "/etc/hosts")
	if err := hostsCmd.Run(); err != nil {
		return fmt.Errorf("failed to update /etc/hosts: %w", err)
	}

	return nil
}

// updateTimezone changes the system timezone
func updateTimezone(timezone string) error {
	// Check if running as root/sudo
	if os.Geteuid() != 0 {
		return fmt.Errorf("must run as root/sudo to change timezone")
	}

	// Validate timezone by checking if it exists in /usr/share/zoneinfo/
	_, err := os.Stat(fmt.Sprintf("/usr/share/zoneinfo/%s", timezone))
	if err != nil {
		return fmt.Errorf("invalid timezone: %s", timezone)
	}

	// Execute timedatectl to set timezone
	cmd := exec.Command("timedatectl", "set-timezone", timezone)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set timezone with timedatectl: %w", err)
	}

	return nil
}
