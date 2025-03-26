package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aakritigkmit/my-go-crud/dto"
	"github.com/aakritigkmit/my-go-crud/internal/services"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(services *services.UserService) *UserHandler {
	return &UserHandler{
		Service: services,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		fmt.Println("Invalid request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		//map[string]string is key value pairs in string for json response
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	err = h.Service.CreateUser(userReq)
	if err != nil {
		fmt.Println("Error creating user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error creating user"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		fmt.Println("ID field is empty")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		fmt.Println("Error fetching user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	json.NewEncoder(w).Encode(user)
}
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error fetching users"})
		return
	}

	json.NewEncoder(w).Encode(users)
}
func (h *UserHandler) UpdateUserAgeByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		fmt.Println("ID field is empty")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}

	var userReq dto.UserRequest
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		fmt.Println("Invalid request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	err = h.Service.UpdateUserAgeByID(id, userReq.Age)
	if err != nil {
		fmt.Println("Error updating user age:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error updating user"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User age updated successfully"})

}
func (h *UserHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		fmt.Println("ID field is empty")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID is required"})
		return
	}

	err := h.Service.DeleteUserByID(id)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error deleting user"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
func (h *UserHandler) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	err := h.Service.DeleteAllUsers()
	if err != nil {
		fmt.Println("Error deleting users:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error deleting users"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "All users deleted successfully"})
}
