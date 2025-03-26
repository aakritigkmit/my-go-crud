package routes

import (
	"github.com/aakritigkmit/my-go-crud/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes(authHandler *handlers.AuthHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/register", authHandler.Register)
	r.Post("/login", authHandler.Login)

	return r
}
