package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	_ "github.com/lib/pq"

	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/state"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

type Adapter struct {
	query *Queries
}

func validateCreateAccount(user models.User) error {
	err := validation.ValidateStruct(
		&user,
		validation.Field(&user.FirstName, validation.Required, validation.Length(2, 100), is.Alpha),
		validation.Field(&user.LastName, validation.Required, validation.Length(2, 100), is.Alpha),
		validation.Field(&user.Email, is.Email, validation.Required),
		validation.Field(&user.Password, validation.Required),
		validation.Field(&user.PhoneNumber, validation.Required),
	)
	return err
}

//used to map the values from account to model.User
func mapAccountToUser(account User) models.User {
	returnUser := models.User{
		Id:          account.ID,
		FirstName:   account.FirstName,
		LastName:    account.LastName,
		PhoneNumber: account.PhoneNumber,
		Email:       account.Email,
		CreatedAt:   account.CreatedAt,
		State:       account.State,
	}
	return returnUser
}

func NewAdapter(env string) *Adapter {
	config, err := utils.LoadConfig(env)
	if err != nil {
		log.Fatal("can not load config file on database ", models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 50})
	}
	DBSouce := config.DBSource
	DBDriver := config.DBDriver
	db, err := sql.Open(DBDriver, DBSouce)
	if err != nil {
		log.Fatal("can not connect to database on database", models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 56})

	}
	query := New(db)

	return &Adapter{
		query: query,
	}
}
func (a *Adapter) Close(ctx context.Context) error {
	return a.query.Close()
}
func (a *Adapter) CreateUser(ctx context.Context, user models.User) (models.User, models.Errors) {
	err := validateCreateAccount(user)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 78}
	}

	account, err := a.query.CreateUser(ctx, CreateUserParams{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
		CreatedAt:   time.Now().GoString(),
		State:       state.Active,
	})
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 91}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}
}

func (a *Adapter) DeleteUser(ctx context.Context, user models.User) models.Errors {

	_, err := a.query.DeleteUser(ctx, user.Id)
	if err != nil {
		return models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 86}
	}

	return models.Errors{}
}

func (a *Adapter) UpdateUserFirstName(ctx context.Context, user models.User, new_first_name string) (models.User, models.Errors) {
	err := validation.Validate(new_first_name, validation.Required, validation.Length(2, 100))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, err := a.query.UpdateUserFirstName(ctx, UpdateUserFirstNameParams{
		FirstName: new_first_name,
		ID:        user.Id,
	})
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}
}
