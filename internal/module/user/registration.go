package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/encription"
)

func validateUser(user models.User) error {

	err := validation.ValidateStruct(
		&user,
		validation.Field(&user.FirstName, validation.Required, validation.Length(2, 100), is.Alpha),
		validation.Field(&user.LastName, validation.Required, validation.Length(2, 100), is.Alpha),
		validation.Field(&user.Email, is.Email, validation.Required),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 1000)),
		validation.Field(&user.PhoneNumber, validation.Required),
	)
	return err
}
func createHasedUser(user models.User) (models.User, error) {
	hasedPassword, err := encription.GenerateHashedPassword(user.Password)
	if err != nil {
		return models.User{}, err
	}
	hashed_password_user := models.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    hasedPassword,
	}
	return hashed_password_user, nil
}
func (s *service) RegisterUser(ctx context.Context, user models.User) (models.User, models.Errors) {
	err := validateUser(user)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "/internal/module/user/registration.go", ErrLine: 15}
	}
	newUser, err := createHasedUser(user)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "/internal/module/user/registration.go", ErrLine: 15}
	}

	return s.databasAdapter.CreateUser(ctx, newUser)

}
