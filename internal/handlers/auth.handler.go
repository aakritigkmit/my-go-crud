package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aakritigkmit/my-go-crud/dto"
	"github.com/aakritigkmit/my-go-crud/internal/services"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(services *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: services,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReq dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.Service.RegisterUser(userReq)
	if err != nil {
		fmt.Println("Error creating user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error creating user"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// // Login an existing user
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
