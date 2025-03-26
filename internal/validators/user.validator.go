package validators

import (
	"github.com/aakritigkmit/my-go-crud/internal/model"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateUser(user *model.User) error {
	return validate.Struct(user)
}
