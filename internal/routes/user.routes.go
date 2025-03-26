package routes

import (
	"github.com/aakritigkmit/my-go-crud/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(userHandler *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", userHandler.CreateUser)
	r.Get("/{id}", userHandler.GetUserByID)
	r.Get("/", userHandler.GetAllUsers)
	r.Put("/{id}", userHandler.UpdateUserAgeByID)
	r.Delete("/", userHandler.DeleteAllUsers)
	r.Delete("/{id}", userHandler.DeleteUserByID)

	return r
}
