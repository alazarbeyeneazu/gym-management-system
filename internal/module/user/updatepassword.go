package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) UpdateUserPassword(ctx context.Context, user models.User, new_password string) (models.User, models.Errors) {

	err := validation.Validate(new_password, validation.Required, validation.Length(8, 1000))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, errModel := s.databasAdapter.UpdateUserPassword(ctx, user, new_password)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}
