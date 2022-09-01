package user

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func (s *service) GetUsersByFirstName(ctx context.Context, first_name string) ([]models.User, models.Errors) {
	err := validation.Validate(first_name, validation.Required, validation.Length(2, 150))
	if err != nil {
		return []models.User{}, models.Errors{Err: err, ErrorLocation: "/internal/module/user/getusersbyfirstname.go", ErrLine: 96}
	}

	accounst, errModel := s.databasAdapter.GetUsersByFirstName(ctx, first_name)
	if errModel.Err != nil {
		return []models.User{}, errModel
	}
	return accounst, models.Errors{}
}

func (s *service) GetUsersByLastName(ctx context.Context, last_name string) ([]models.User, models.Errors) {
	err := validation.Validate(last_name, validation.Required, validation.Length(2, 150))
	if err != nil {
		return []models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}

	accounts, errModel := s.GetUsersByLastName(ctx, last_name)
	if errModel.Err != nil {
		return []models.User{}, errModel
	}
	return accounts, models.Errors{}

}

func (s *service) GetUserByPhoneNumber(ctx context.Context, phone_number string) (models.User, models.Errors) {

	err := validation.Validate(phone_number, validation.Required, validation.Length(13, 13))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, errorModel := s.databasAdapter.GetUserByPhoneNumber(ctx, phone_number)
	if errorModel.Err != nil {
		return models.User{}, errorModel
	}
	return account, models.Errors{}

}
