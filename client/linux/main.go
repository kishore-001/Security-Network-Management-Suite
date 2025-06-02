// main.go
package main

import (
	"linux/api"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Register routes for config-1 page
	api.RegisterConfig1Routes(mux)
	api.RegisterConfig2Routes(mux)
	api.RegisterHealthRoutes(mux)
	api.RegisterOptimizeRoutes(mux)

	log.Println("Starting client server on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
