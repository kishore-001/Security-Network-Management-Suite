package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// SSHKeyRequest defines the expected JSON structure for SSH key uploads
type SSHKeyRequest struct {
	Key string `json:"key"`
}

func main() {
	// Simulate a request with a JSON string (you can replace this with real input)
	jsonStr := `{"key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCz... your-test-key"}`

	fmt.Println("Current Date and Time (UTC): 2023-10-10 12:00:00")
	fmt.Println("Processing SSH key addition on", runtime.GOOS)

	// Parse the JSON
	var keyRequest SSHKeyRequest
	err := json.Unmarshal([]byte(jsonStr), &keyRequest)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	// Validate that the key was provided
	if keyRequest.Key == "" {
		fmt.Println("Error: SSH key is empty")
		os.Exit(1)
	}

	// Get the current user details
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user info:", err)
		os.Exit(1)
	}

	homeDir := currentUser.HomeDir
	fmt.Println("Current User:", currentUser.Username)
	fmt.Println("Home Directory:", homeDir)

	// Build the .ssh directory path
	sshDir := filepath.Join(homeDir, ".ssh")

	// Create .ssh directory if it doesn't exist
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		fmt.Println("Error creating .ssh directory:", err)
		os.Exit(1)
	}

	// Build the authorized_keys file path
	sshPath := filepath.Join(sshDir, "authorized_keys")
	fmt.Println("Authorized Keys Path:", sshPath)

	// Ensure SSH key ends with newline
	sshKey := keyRequest.Key
	if sshKey[len(sshKey)-1] != '\n' {
		sshKey += "\n"
	}

	// Check file status
	fileInfo, err := os.Stat(sshPath)
	if err == nil {
		fmt.Printf("authorized_keys exists. Current size: %d bytes\n", fileInfo.Size())
	} else if os.IsNotExist(err) {
		fmt.Println("authorized_keys does not exist yet. A new file will be created.")
	} else {
		fmt.Println("Error checking file:", err)
		os.Exit(1)
	}

	// Open and append the key to the file
	file, err := os.OpenFile(sshPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println("Error opening authorized_keys file:", err)
		os.Exit(1)
	}
	defer file.Close()

	bytesWritten, err := file.WriteString(sshKey)
	if err != nil {
		fmt.Println("Error writing SSH key:", err)
		os.Exit(1)
	}

	fmt.Printf("✅ Successfully wrote %d bytes to %s\n", bytesWritten, sshPath)
	fmt.Println("🎉 SSH key added successfully!")

	if runtime.GOOS == "windows" {
		fmt.Println("ℹ️  If using OpenSSH on Windows, ensure the 'sshd' service is installed and your SSH agent is enabled.")
	}
}
