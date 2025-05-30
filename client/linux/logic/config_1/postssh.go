package config_1

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

// SSHKeyRequest defines the expected JSON structure for SSH key uploads
type SSHKeyRequest struct {
	Key string `json:"key"`
}

func HandleSSHUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request
	var keyRequest SSHKeyRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&keyRequest); err != nil {
		http.Error(w, "Failed to parse JSON request", http.StatusBadRequest)
		return
	}

	// Validate that the key was provided
	if keyRequest.Key == "" {
		http.Error(w, "SSH key is empty", http.StatusBadRequest)
		return
	}

	// Get the user's home directory and construct the .ssh path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		http.Error(w, "Failed to determine home directory", http.StatusInternalServerError)
		return
	}

	// Create the .ssh directory if it doesn't exist
	sshDir := filepath.Join(homeDir, ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		http.Error(w, "Failed to create .ssh directory", http.StatusInternalServerError)
		return
	}

	// Path to the authorized_keys file
	sshPath := filepath.Join(sshDir, "authorized_keys")

	// Append a newline if the key doesn't end with one
	sshKey := keyRequest.Key
	if sshKey[len(sshKey)-1] != '\n' {
		sshKey += "\n"
	}

	// Write the SSH key to the file
	// Using os.O_APPEND to add to existing keys rather than overwrite
	file, err := os.OpenFile(sshPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		http.Error(w, "Failed to open authorized_keys file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(sshKey); err != nil {
		http.Error(w, "Failed to write SSH key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SSH key uploaded successfully"))
}
