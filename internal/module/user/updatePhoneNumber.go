package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) UpdateUserPhoneNumber(ctx context.Context, user models.User, new_phone_number string) (models.User, models.Errors) {

	err := validation.Validate(new_phone_number, validation.Required, validation.Length(13, 13))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/module/user/updatePhoneNumber.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/module/user/updatePhoneNumber.go", ErrLine: 96}
	}
	account, errModel := s.databasAdapter.UpdateUserPhoneNumber(ctx, user, new_phone_number)
	if errModel.Err != nil {
		return models.User{}, errModel
	}
	return account, models.Errors{}
}
