package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aakritigkmit/my-go-crud/dto"
	"github.com/aakritigkmit/my-go-crud/internal/model"
	"github.com/aakritigkmit/my-go-crud/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthRepo *repository.AuthRepo
	Repo     *repository.UserRepo
}

func NewAuthService(authRepo *repository.AuthRepo, repo *repository.UserRepo) *AuthService {
	return &AuthService{AuthRepo: authRepo, Repo: repo}
}

func (s *AuthService) RegisterUser(userReq dto.UserRequest) error {

	existingUser, _ := s.AuthRepo.FindUserByEmail(userReq.Email)
	if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", userReq.Email)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	user := model.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: string(hashedPassword),
	}

	userID, err := s.Repo.CreateUser(user) // Call the repository method
	if err != nil {
		fmt.Println("Error while register user in DB:", err)
		return err
	}

	fmt.Println("User successfully Registered ", userID)
	return nil

}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.AuthRepo.FindUserByEmail(email)
	if err != nil {
		fmt.Println("Error finding user:", err) // Print error
		return "", errors.New("invalid credentials")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println("Password mismatch:", err) // Print error
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	tokenString, err := generateJWT(user.ID.Hex(), user.Email)
	if err != nil {
		fmt.Println("Error generating JWT:", err) // Print error
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func generateJWT(userID string, email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("JWT_SECRET is not set") // Print error
		return "", fmt.Errorf("JWT_SECRET is not set")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Error signing JWT token:", err) // Print error
		return "", err
	}

	return tokenString, nil
}
