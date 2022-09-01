package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) UpdateUserEmail(ctx context.Context, user models.User, new_email string) (models.User, models.Errors) {
	err := validation.Validate(new_email, validation.Required, is.Email)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "/internal/module/user/updateEmail.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "/internal/module/user/updateEmail.go", ErrLine: 96}
	}
	account, errModel := s.databasAdapter.UpdateUserEmail(ctx, user, new_email)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}
