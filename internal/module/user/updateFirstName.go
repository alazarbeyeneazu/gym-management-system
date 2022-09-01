package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) UpdateUserFirstName(ctx context.Context, user models.User, new_first_name string) (models.User, models.Errors) {
	err := validation.Validate(new_first_name, validation.Required, validation.Length(2, 100))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/module/user/updateFirstName.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/module/user/updateFirstName.go", ErrLine: 96}
	}

	account, errModel := s.databasAdapter.UpdateUserFirstName(ctx, user, new_first_name)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}

}
