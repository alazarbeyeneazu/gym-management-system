package db

//this is adapter page which implements DBPort interface
import (
	"context"
	"database/sql"
	"log"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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
		validation.Field(&user.Password, validation.Required, validation.Length(8, 1000)),
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
		Password:    account.Password,
	}
	return returnUser
}

//used to map the values from accounts to model.User
func mapAccountsToUsers(accounts []User) []models.User {
	var users []models.User
	for _, account := range accounts {
		returnUser := models.User{
			Id:          account.ID,
			FirstName:   account.FirstName,
			LastName:    account.LastName,
			PhoneNumber: account.PhoneNumber,
			Email:       account.Email,
			CreatedAt:   account.CreatedAt,
			State:       account.State,
		}
		users = append(users, returnUser)
	}

	return users
}

// addapte to implement database port
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

//close the database
func (a *Adapter) Close(ctx context.Context) models.Errors {
	err := a.query.Close()
	if err != nil {
		return models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 93}
	}
	return models.Errors{}
}

//implement create user for database
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

//implement delete user

func (a *Adapter) DeleteUser(ctx context.Context, user models.User) models.Errors {

	_, err := a.query.DeleteUser(ctx, user.Id)
	if err != nil {
		return models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 86}
	}

	return models.Errors{}
}

//implement update first name for database
func (a *Adapter) UpdateUserFirstName(ctx context.Context, user models.User, new_first_name string) (models.User, models.Errors) {
	err := validation.Validate(new_first_name, validation.Required, validation.Length(2, 100))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
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

//update last name
func (a *Adapter) UpdateUserLastName(ctx context.Context, user models.User, new_last_name string) (models.User, models.Errors) {

	err := validation.Validate(new_last_name, validation.Required, validation.Length(2, 100))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}

	account, err := a.query.UpdateUserLastName(ctx, UpdateUserLastNameParams{
		LastName: new_last_name,
		ID:       user.Id,
	})
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 108}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}
}

//update phone number for user
func (a *Adapter) UpdateUserPhoneNumber(ctx context.Context, user models.User, new_phone_number string) (models.User, models.Errors) {

	err := validation.Validate(new_phone_number, validation.Required, validation.Length(13, 13))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, err := a.query.UpdateUserPhoneNumber(ctx, UpdateUserPhoneNumberParams{
		PhoneNumber: new_phone_number,
		ID:          user.Id,
	})
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 119}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}
}

//update email for the user
func (a *Adapter) UpdateUserEmail(ctx context.Context, user models.User, new_email string) (models.User, models.Errors) {
	err := validation.Validate(new_email, validation.Required, is.Email)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, err := a.query.UpdateUserEmail(ctx, UpdateUserEmailParams{
		Email: new_email,
		ID:    user.Id,
	})
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 130}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}

}

//update password
func (a *Adapter) UpdateUserPassword(ctx context.Context, user models.User, new_password string) (models.User, models.Errors) {

	err := validation.Validate(new_password, validation.Required, validation.Length(8, 1000))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	err = validation.Validate(user.Id, validation.Required)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, err := a.query.UpdateUsersPassword(ctx, UpdateUsersPasswordParams{
		Password: new_password,
		ID:       user.Id,
	})
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 142}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}
}

// get user by first name
func (a *Adapter) GetUsersByFirstName(ctx context.Context, first_name string) ([]models.User, models.Errors) {
	err := validation.Validate(first_name, validation.Required, validation.Length(2, 150))
	if err != nil {
		return []models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}

	accounts, err := a.query.GetUsersByFirstName(ctx, first_name)
	if err != nil {
		return []models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 150}
	}
	returnUsers := mapAccountsToUsers(accounts)
	return returnUsers, models.Errors{}
}

//get user by last name
func (a *Adapter) GetUsersByLastName(ctx context.Context, last_name string) ([]models.User, models.Errors) {
	err := validation.Validate(last_name, validation.Required, validation.Length(2, 150))
	if err != nil {
		return []models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}

	accounts, err := a.query.GetUsersByLastName(ctx, last_name)
	if err != nil {
		return []models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 158}
	}
	returnUsers := mapAccountsToUsers(accounts)
	return returnUsers, models.Errors{}

}

//get user by phone number
func (a *Adapter) GetUserByPhoneNumber(ctx context.Context, phone_number string) (models.User, models.Errors) {

	err := validation.Validate(phone_number, validation.Required, validation.Length(13, 13))
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, err := a.query.GetUserByPhoneNumber(ctx, phone_number)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 167}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}
}

//get user by email

func (a *Adapter) GetUserByEmail(ctx context.Context, email string) (models.User, models.Errors) {
	err := validation.Validate(email, validation.Required, is.Email)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 96}
	}
	account, err := a.query.GetUserEmail(ctx, email)
	if err != nil {
		return models.User{}, models.Errors{Err: err, ErrorLocation: "internal/storage/persistant/sqlc/Adapter.go", ErrLine: 175}
	}
	returnUser := mapAccountToUser(account)
	return returnUser, models.Errors{}
}
