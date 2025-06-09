package main

import (
	"backend/api/common"
	"backend/api/server"
	"backend/config"
	"log"
	"net/http"
)

func main() {
	// Create separate muxes for different protection levels
	publicMux := http.NewServeMux()
	protectedMux := http.NewServeMux()
	adminMux := http.NewServeMux()

	// This is necessary for Database
	queries := config.GenerateQueries()

	// 🌐 Public routes (no authentication required)
	common.RegisterAuthRoutes(publicMux, queries)

	// 🔒 Protected routes (authentication required)

	//  Server Protected Routes
	//	server.RegisterHealthRoutes(protectedMux, queries)
	//	server.RegisterLogRoutes(protectedMux, queries)

	//  Network Protected Routes

	//	network.RegisterHealthRoutes(protectedMux, queries)
	//	network.RegisterLogRoutes(protectedMux, queries)

	// 👑 Admin-only routes

	//  Common Admin Routes
	common.RegisterSettingsRoutes(adminMux, queries)

	//  Server Admin Routes
	server.RegisterConfig1Routes(adminMux, queries)
	//	server.RegisterBackupRoutes(adminMux, queries)

	//  Network Admin Routes

	//	network.RegisterConfigRoutes(adminMux, queries)
	//	network.RegisterBackupRoutes(adminMux, queries)

	// Create main mux and apply appropriate middlewares
	mainMux := http.NewServeMux()

	// Mount with different middleware chains
	mainMux.Handle("/api/auth/", config.ApplyPublicMiddlewares(publicMux))
	mainMux.Handle("/api/server/", config.ApplyProtectedMiddlewares(protectedMux))
	mainMux.Handle("/api/network/", config.ApplyProtectedMiddlewares(protectedMux))
	mainMux.Handle("/api/admin/", config.ApplyAdminMiddlewares(adminMux))

	log.Println("✅ SNSMS backend running on port 8000...")
	if err := http.ListenAndServe(":8000", mainMux); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
