package config_1

import (
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

func HandleSSHUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Get current user
	currentUser, err := user.Current()
	if err != nil {
		http.Error(w, "Unable to determine current user", http.StatusInternalServerError)
		return
	}

	// Windows-style SSH directory path
	sshDir := filepath.Join(currentUser.HomeDir, ".ssh")
	authKeysPath := filepath.Join(sshDir, "authorized_keys")

	// Create .ssh directory if it doesn't exist
	if _, err := os.Stat(sshDir); os.IsNotExist(err) {
		if err := os.MkdirAll(sshDir, 0700); err != nil {
			http.Error(w, "Failed to create .ssh directory", http.StatusInternalServerError)
			return
		}
	}

	// Overwrite (or write) SSH key to authorized_keys
	err = os.WriteFile(authKeysPath, body, 0600)
	if err != nil {
		http.Error(w, "Failed to write SSH key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SSH key uploaded successfully"))
}
