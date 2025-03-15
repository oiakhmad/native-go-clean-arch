package routes

import (
	"database/sql"
	"net/http"
	// "golang-clean-architecture/interface/http/handler"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	// Health check route
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
