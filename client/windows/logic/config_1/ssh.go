package config_1

import (
	"encoding/json"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

// Define the structure of the expected JSON
type SSHKeyPayload struct {
	PublicKey string `json:"public_key"`
}

func HandleSSHUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON body
	var payload SSHKeyPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if payload.PublicKey == "" {
		http.Error(w, "Public key is required", http.StatusBadRequest)
		return
	}

	// Get current user home directory
	currentUser, err := user.Current()
	if err != nil {
		http.Error(w, "Unable to determine current user", http.StatusInternalServerError)
		return
	}

	sshDir := filepath.Join(currentUser.HomeDir, ".ssh")
	authKeysPath := filepath.Join(sshDir, "authorized_keys")

	// Create .ssh directory if it doesn't exist
	if _, err := os.Stat(sshDir); os.IsNotExist(err) {
		if err := os.MkdirAll(sshDir, 0700); err != nil {
			http.Error(w, "Failed to create .ssh directory", http.StatusInternalServerError)
			return
		}
	}

	// Open the authorized_keys file for appending
	f, err := os.OpenFile(authKeysPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		http.Error(w, "Failed to open authorized_keys", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Write the public key with a newline
	if _, err := f.WriteString(payload.PublicKey + "\n"); err != nil {
		http.Error(w, "Failed to write SSH key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SSH key uploaded successfully"))
}
