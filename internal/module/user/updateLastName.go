package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) UpdateUserLastName(ctx context.Context, user models.User, new_last_name string) (models.User, models.Errors) {

	err := validation.Validate(new_last_name, validation.Required, validation.Length(2, 100))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/module/user/updateLastName.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/module/user/updateLastName.go", ErrLine: 96}
	}

	account, errModel := s.databasAdapter.UpdateUserLastName(ctx, user, new_last_name)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}
