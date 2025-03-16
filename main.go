package main

import (
	"log"
	"net/http"
	"os"

	"native-go-clean-arch/config"
	"native-go-clean-arch/infrastructure/db"
	routes "native-go-clean-arch/interface/http"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Initialize database
	dbConn := db.InitMySQL()
	defer dbConn.Close()

	// Setup HTTP server
	// httpMux := http.NewServeMux()
	// routes.RegisterRoutes(httpMux, dbConn)

	r := routes.RegisterRoutes(dbConn)
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	log.Println("Server running on port:", serverPort)
	if err := http.ListenAndServe(":"+serverPort, r); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
