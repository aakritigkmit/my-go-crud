package services

import (
	"fmt"

	"github.com/aakritigkmit/my-go-crud/dto"
	"github.com/aakritigkmit/my-go-crud/internal/model"
	"github.com/aakritigkmit/my-go-crud/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		// DBClient repository.UserInterface
		Repo: repo,
	}
}

func (s *UserService) CreateUser(userReq dto.UserRequest) error {
	user := model.User{
		Name:    userReq.Name,
		Age:     userReq.Age,
		Country: userReq.Country,
	}

	userID, err := s.Repo.CreateUser(user) // Call the repository method
	if err != nil {
		fmt.Println("Error while inserting user in DB:", err)
		return err
	}

	fmt.Println("User successfully inserted", userID)
	return nil

}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	var user model.User
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		fmt.Println("Error fetching user by ID:", err)
		return model.User{}, err
	}

	fmt.Println("User fetched successfully")
	return user, nil
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	var users []model.User

	users, err := s.Repo.GetAllUsers()
	if err != nil {
		fmt.Println("Error fetching all users:", err)
		return nil, err
	}

	fmt.Println("Users fetched successfully")
	return users, nil

}

func (s *UserService) UpdateUserAgeByID(id string, age int) error {

	updatedCount, err := s.Repo.UpdateUserAgeByID(id, age)
	if err != nil {
		fmt.Println("Error while updating user age:", err)
		return err
	}

	fmt.Println("User successfully updated", updatedCount)
	return nil
}

// func (s *UserService) DeleteUserByID() {

// 	res := dto.UserResponse{}
// 	id := chi.URLParam(r, "id")

// 	if id == " " {
// 		fmt.Println("id field is empty")
// 		res.Error = "id doesn't exist"
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(res)
// 		return

// 	}

// 	result, err := srv.DBClient.DeleteUserByID(id)
// 	if err != nil {
// 		slog.Error(err.Error())
// 		res.Error = " error while deleting  user from db"
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

//		fmt.Println("user successfully deleted ", result)
//		json.NewEncoder(w).Encode(res)
//	}

func (s *UserService) DeleteUserByID(id string) error {
	deletedCount, err := s.Repo.DeleteUserByID(id)
	if err != nil {
		fmt.Println("Error while deleting user:", err)
		return err
	}

	fmt.Println("User successfully deleted", deletedCount)
	return nil
}

func (s *UserService) DeleteAllUsers() error {
	deletedCount, err := s.Repo.DeleteAllUsers()
	if err != nil {
		fmt.Println("Error while deleting all users:", err)
		return err
	}

	fmt.Println("All users successfully deleted", deletedCount)
	return nil
}
