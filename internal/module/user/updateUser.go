package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) UpdateUserPassword(ctx context.Context, user models.User, new_password string) (models.User, models.Errors) {

	err := validation.Validate(new_password, validation.Required, validation.Length(8, 1000))
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("password ", err), ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("Id ", err), ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, errModel := s.databasAdapter.UpdateUserPassword(ctx, user, new_password)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}

func (s *service) UpdateUserPhoneNumber(ctx context.Context, user models.User, new_phone_number string) (models.User, models.Errors) {

	err := validation.Validate(new_phone_number, validation.Required, validation.Length(13, 13))
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("phone number ", err), ErrorLocation: "internal/module/user/updatePhoneNumber.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("Id ", err), ErrorLocation: "internal/module/user/updatePhoneNumber.go", ErrLine: 96}
	}
	account, errModel := s.databasAdapter.UpdateUserPhoneNumber(ctx, user, new_phone_number)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}

func (s *service) UpdateUserLastName(ctx context.Context, user models.User, new_last_name string) (models.User, models.Errors) {

	err := validation.Validate(new_last_name, validation.Required, validation.Length(2, 100))
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("Last Name ", err), ErrorLocation: "internal/module/user/updateLastName.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("Id ", err), ErrorLocation: "internal/module/user/updateLastName.go", ErrLine: 96}
	}

	account, errModel := s.databasAdapter.UpdateUserLastName(ctx, user, new_last_name)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}

func (s *service) UpdateUserFirstName(ctx context.Context, user models.User, new_first_name string) (models.User, models.Errors) {
	err := validation.Validate(new_first_name, validation.Required, validation.Length(2, 100))
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("First Name ", err), ErrorLocation: "internal/module/user/updateFirstName.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("Id ", err), ErrorLocation: "internal/module/user/updateFirstName.go", ErrLine: 96}
	}

	account, errModel := s.databasAdapter.UpdateUserFirstName(ctx, user, new_first_name)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}

}

func (s *service) UpdateUserEmail(ctx context.Context, user models.User, new_email string) (models.User, models.Errors) {
	err := validation.Validate(new_email, validation.Required, is.Email)
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("Email ", err), ErrorLocation: "/internal/module/user/updateEmail.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: errorHelper("Id ", err), ErrorLocation: "/internal/module/user/updateEmail.go", ErrLine: 96}
	}
	account, errModel := s.databasAdapter.UpdateUserEmail(ctx, user, new_email)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}
