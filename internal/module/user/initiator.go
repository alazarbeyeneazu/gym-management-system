package user

import (
	"context"

	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	db "gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/storage/persistant/sqlc"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/ports"
)

type UserService interface {
	RegisterUser(ctx context.Context, user models.User) (models.User, models.Errors)
}
type service struct {
	databasAdapter ports.DBPort
}

func Initiate() *service {
	database := db.NewAdapter("../../../")
	return &service{databasAdapter: database}
}
