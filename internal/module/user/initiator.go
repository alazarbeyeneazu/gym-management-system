package user

import (
	"context"

	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/ports"
)

//go:generate echo $PWD - $GOPACKAGE - $GOFILE
//go:generate mockgen -package mockapp -destination ../../../mocks/app/appMock.go gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/module/user UserService
type UserService interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, models.Errors)
	DeleteUser(ctx context.Context, user models.User) models.Errors

	UpdateUserFirstName(ctx context.Context, user models.User, new_first_name string) (models.User, models.Errors)
	UpdateUserLastName(ctx context.Context, user models.User, new_last_name string) (models.User, models.Errors)
	UpdateUserPhoneNumber(ctx context.Context, user models.User, new_phone_number string) (models.User, models.Errors)
	UpdateUserEmail(ctx context.Context, user models.User, new_email string) (models.User, models.Errors)
	UpdateUserPassword(ctx context.Context, user models.User, new_password string) (models.User, models.Errors)

	GetUsersByLastName(ctx context.Context, last_name string) ([]models.User, models.Errors)
	GetUsersByFirstName(ctx context.Context, first_name string) ([]models.User, models.Errors)
	GetUserByPhoneNumber(ctx context.Context, phone_number string) (models.User, models.Errors)
	GetUserByEmail(ctx context.Context, email string) (models.User, models.Errors)
}
type service struct {
	databasAdapter ports.DBPort
}

func Initiate(env string, dbs ports.DBPort) *service {
	return &service{databasAdapter: dbs}
}
