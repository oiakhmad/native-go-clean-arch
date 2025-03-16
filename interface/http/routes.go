package routes

import (
	"database/sql"
	"fmt"
	"native-go-clean-arch/interface/http/handler"
	"native-go-clean-arch/internal/repository"
	"native-go-clean-arch/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
	// "golang-clean-architecture/interface/http/handler"
)

func RegisterRoutes(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// Inisialisasi repository, usecase, dan handler
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Tambahkan user route
	r.Get("/users", userHandler.GetUsers)

	// Debug: Print semua route yang terdaftar
	fmt.Println("Registered Routes:")
	_ = chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	})

	return r
}

// func RegisterRoutes(db *sql.DB) {
// 	// Health check route
// 	// mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.WriteHeader(http.StatusOK)
// 	// 	w.Write([]byte("OK"))
// 	// })
// 	r := chi.NewRouter()

// 	// Inisialisasi repository, usecase, dan handler
// 	userRepo := repository.NewUserRepository(db)
// 	userUsecase := usecase.NewUserUsecase(userRepo)
// 	userHandler := handler.NewUserHandler(userUsecase)

// 	// Route untuk User
// 	r.Get("/users", userHandler.GetUsers)
// }
