package config_1

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleSSHUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	sshPath := os.Getenv("HOME") + "/.ssh/authorized_keys"
	err = ioutil.WriteFile(sshPath, body, 0600)
	if err != nil {
		http.Error(w, "Failed to write SSH key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SSH key uploaded successfully"))
}
