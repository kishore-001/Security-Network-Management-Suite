package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"linux/api"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const tokenFilePath = "./auth/token.hash"

func main() {
	ensureTokenHashExists()

	mux := http.NewServeMux()

	// Register API routes
	api.RegisterConfig1Routes(mux)
	api.RegisterConfig2Routes(mux)
	api.RegisterHealthRoutes(mux)
	api.RegisterOptimizeRoutes(mux)
	api.RegisterLogRoutes(mux)

	log.Println("Starting client server on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// ------------------------------
// Check and create token hash if missing
// ------------------------------
func ensureTokenHashExists() {
	if _, err := os.Stat(tokenFilePath); os.IsNotExist(err) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("🔐 Enter token to register this client: ")
		token, _ := reader.ReadString('\n')
		token = trim(token)

		hash := sha256.Sum256([]byte(token))
		hashHex := hex.EncodeToString(hash[:])

		// Ensure auth folder exists
		os.MkdirAll(filepath.Dir(tokenFilePath), 0755)

		if err := os.WriteFile(tokenFilePath, []byte(hashHex), 0644); err != nil {
			log.Fatalf("❌ Failed to save token hash: %v", err)
		}

		fmt.Println("✅ Token hash saved. Server starting...")
	}
}

func trim(s string) string {
	return string([]byte(s)[:len(s)-1]) // remove newline character
}
