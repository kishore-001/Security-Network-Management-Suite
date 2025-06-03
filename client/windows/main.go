// main.go
package main

import (
	"log"
	"net/http"
	"windows/api"
)

func main() {
	mux := http.NewServeMux()

	// Register routes for config-1 page
	api.RegisterConfig1Routes(mux)
	// Register routes for config-2 page
	api.RegisterConfig2Routes(mux)
	//Register routes for health endpoint
	api.RegisterHealthRoutes(mux)
	// Register routes for optimization
	api.RegisterOptimizeRoutes(mux)
	// Register routes for logs
	api.RegisterLogRoutes(mux)

	log.Println("Starting client server on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
