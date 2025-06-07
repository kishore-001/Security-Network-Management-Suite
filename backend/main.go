package main

import (
	"backend/api/common"
	"backend/config"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	queries := config.GenerateQueries()
	// 🧠 Registering Server API Routes
	// server.RegisterAlertRoutes(mux)
	// server.RegisterBackupRoutes(mux)
	// server.RegisterConfig1Routes(mux)
	// server.RegisterConfig2Routes(mux)
	// server.RegisterHealthRoutes(mux)
	// server.RegisterLogRoutes(mux)
	// server.RegisterOptimizeRoutes(mux)

	// 🌐 Registering Network API Routes
	// network.RegisterAlertRoutes(mux)
	//	network.RegisterBackupRoutes(mux)
	//	network.RegisterConfig1Routes(mux)
	//	network.RegisterConfig2Routes(mux)
	//	network.RegisterHealthRoutes(mux)
	//	network.RegisterLogRoutes(mux)
	//	network.RegisterOptimizeRoutes(mux)

	// 🛠️ Register Common Routes (Login, Settings, etc.)
	//	common.RegisterSettingsRoutes(mux)
	common.RegisterAuthRoutes(mux, queries)
	handlerWithCORS := config.CORS(mux)

	log.Println("✅ SNSMS backend running on port 8000...")
	if err := http.ListenAndServe(":8000", handlerWithCORS); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
