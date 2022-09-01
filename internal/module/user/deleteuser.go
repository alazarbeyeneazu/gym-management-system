package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) DeleteUser(ctx context.Context, user models.User) models.Errors {
	err := validation.Validate(user.Id)
	if err != nil {
		return models.Errors{Err: err, ErrorLocation: "/internal/module/user/registration.go"}
	}
	errModel := s.databasAdapter.DeleteUser(context.Background(), user)
	if errModel.Err != nil {
		return errModel
	}
	return models.Errors{}

}
