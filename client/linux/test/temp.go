package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// SSHKeyRequest defines the expected JSON structure for SSH key uploads
type SSHKeyRequest struct {
	Key string `json:"key"`
}

func main() {
	// Simulate a request with a JSON string
	jsonStr := `{"key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCz... your-test-key"}`

	fmt.Println("Current Date and Time (UTC): 2023-10-10 12:00:00")
	fmt.Println("Processing SSH key addition...")

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

	// Get the user's home directory and construct the .ssh path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error determining home directory:", err)
		os.Exit(1)
	}

	// Display the current user
	currentUser := os.Getenv("USER")
	if currentUser == "" {
		currentUser = "unknown"
	}
	fmt.Println("Current User's Login:", currentUser)

	// Create the .ssh directory if it doesn't exist
	sshDir := filepath.Join(homeDir, ".ssh")
	if err := os.MkdirAll(sshDir, 0700); err != nil {
		fmt.Println("Error creating .ssh directory:", err)
		os.Exit(1)
	}

	// Path to the authorized_keys file
	sshPath := filepath.Join(sshDir, "authorized_keys")
	fmt.Println("Path to authorized_keys:", sshPath)

	// Append a newline if the key doesn't end with one
	sshKey := keyRequest.Key
	if len(sshKey) > 0 && sshKey[len(sshKey)-1] != '\n' {
		sshKey += "\n"
	}

	// Check if file exists before trying to append
	fileInfo, err := os.Stat(sshPath)
	if err == nil {
		fmt.Printf("File exists. Current size: %d bytes\n", fileInfo.Size())
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist yet. Will create new file.")
	} else {
		fmt.Println("Error checking file:", err)
		os.Exit(1)
	}

	// Write the SSH key to the file
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

	fmt.Printf("Successfully wrote %d bytes to %s\n", bytesWritten, sshPath)
	fmt.Println("SSH key added successfully!")
}
