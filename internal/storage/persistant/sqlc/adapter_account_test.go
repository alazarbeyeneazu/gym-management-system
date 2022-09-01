package db

// this page only to test user related operations
import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

var adapter *Adapter = NewAdapter("../../../../")

func generateRandomUser() models.User {
	var email = utils.RandomeEmail()
	var firstName = utils.RandomUserName()
	var lastName = utils.RandomUserName()
	var phoneNumber = utils.RandomePhoneNumber()
	return models.User{
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Email:       email,
		Password:    utils.RandomPassword(),
	}

}

//test create account
func TestCreateUser(t *testing.T) {
	randomUser := generateRandomUser()
	testCase := []struct {
		name        string
		account     models.User
		checkResult func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:    "ok",
			account: randomUser,
			checkResult: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.NotEmpty(t, user)
				require.NotEmpty(t, user.FirstName)
				require.NotEmpty(t, user.LastName)
				require.NotEmpty(t, user.Email)
				require.NotEmpty(t, user.State)
				require.Empty(t, user.Password)
				require.Equal(t, user.Email, randomUser.Email)
				require.Equal(t, user.FirstName, randomUser.FirstName)
				require.Equal(t, user.LastName, randomUser.LastName)
				require.Equal(t, user.PhoneNumber, randomUser.PhoneNumber)

			},
		},
		{
			name: "empty first name",
			account: models.User{
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Email:       utils.RandomeEmail(),
				Password:    utils.RandomPassword(),
			},
			checkResult: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Equal(t, err.Err.Error(), "first_name: cannot be blank.")
				require.Equal(t, err.ErrorLocation, "internal/storage/persistant/sqlc/Adapter.go")
				require.Empty(t, user)

			},
		},
		{
			name: "empty last name",
			account: models.User{
				FirstName:   utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Email:       utils.RandomeEmail(),
				Password:    utils.RandomPassword(),
			},
			checkResult: func(t *testing.T, user models.User, err models.Errors) {

				require.NotEmpty(t, err)
				require.Equal(t, err.Err.Error(), "last_name: cannot be blank.")
				require.Equal(t, err.ErrorLocation, "internal/storage/persistant/sqlc/Adapter.go")
				require.Empty(t, user)

			},
		},
		{
			name: "Empty PhoneNumber",
			account: models.User{
				FirstName: utils.RandomUserName(),
				LastName:  utils.RandomUserName(),
				Email:     utils.RandomeEmail(),
				Password:  utils.RandomPassword(),
			},
			checkResult: func(t *testing.T, user models.User, err models.Errors) {

				require.NotEmpty(t, err)
				require.Equal(t, err.Err.Error(), "phone_number: cannot be blank.")
				require.Equal(t, err.ErrorLocation, "internal/storage/persistant/sqlc/Adapter.go")
				require.Empty(t, user)
			},
		},
		{
			name: "empty email",
			account: models.User{
				FirstName:   utils.RandomUserName(),
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Password:    utils.RandomPassword(),
			},
			checkResult: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Equal(t, err.Err.Error(), "email: cannot be blank.")
				require.Equal(t, err.ErrorLocation, "internal/storage/persistant/sqlc/Adapter.go")
				require.Empty(t, user)
			},
		},
		{
			name: "empty password",
			account: models.User{
				FirstName:   utils.RandomUserName(),
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Email:       utils.RandomeEmail(),
			},
			checkResult: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Equal(t, err.Err.Error(), "password: cannot be blank.")
				require.Equal(t, err.ErrorLocation, "internal/storage/persistant/sqlc/Adapter.go")
				require.Empty(t, user)
			},
		},
	}

	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {
			account, err := adapter.CreateUser(context.Background(), tc.account)
			tc.checkResult(t, account, err)
		})
	}

}

//test Delete Account
func TestDeleteUser(t *testing.T) {
	randomUser := generateRandomUser()
	account, err := adapter.CreateUser(context.Background(), randomUser)

	if err.Err != nil {
		t.Error(t, err)
	}

	testCase := []struct {
		name    string
		user    models.User
		checker func(t *testing.T, err models.Errors)
	}{
		{
			name: "ok",
			user: account,
			checker: func(t *testing.T, err models.Errors) {
				require.Empty(t, err)

			},
		},
		{
			name: "not found ",
			user: models.User{
				Id: rand.Int63(),
			},
			checker: func(t *testing.T, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)

			},
		},
		{
			name: "empty id",
			user: models.User{},
			checker: func(t *testing.T, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)

			},
		},
	}
	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {
			err := adapter.DeleteUser(context.Background(), tc.user)
			tc.checker(t, err)
		})
	}

}

//update first name  test
func TestUpdateFirstName(t *testing.T) {
	randomUser := generateRandomUser()
	new_Name := utils.RandomUserName()
	account, err := adapter.CreateUser(context.Background(), randomUser)
	if err.Err != nil {
		t.Error(err)
	}
	testCase := []struct {
		name    string
		user    models.User
		newName string
		checker func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:    "ok",
			user:    account,
			newName: new_Name,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.FirstName, new_Name)
				require.Equal(t, user.Id, account.Id)

			},
		},
		{
			name:    "empty first name",
			user:    models.User{Id: account.Id},
			newName: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
		{
			name:    "too short",
			user:    models.User{Id: account.Id},
			newName: "h",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ant, err := adapter.UpdateUserFirstName(context.Background(), tc.user, tc.newName)
			tc.checker(t, ant, err)
		})
	}
}

//update last  name  test
func TestUpdateLastName(t *testing.T) {
	randomUser := generateRandomUser()
	new_Name := utils.RandomUserName()
	account, err := adapter.CreateUser(context.Background(), randomUser)
	if err.Err != nil {
		t.Error(err)
	}
	testCase := []struct {
		name    string
		user    models.User
		newName string
		checker func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:    "ok",
			user:    account,
			newName: new_Name,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.LastName, new_Name)
				require.Equal(t, user.Id, account.Id)

			},
		},
		{
			name:    "empty Last name",
			user:    models.User{Id: account.Id},
			newName: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
		{
			name:    "too short",
			user:    models.User{Id: account.Id},
			newName: "h",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ant, err := adapter.UpdateUserLastName(context.Background(), tc.user, tc.newName)
			tc.checker(t, ant, err)
		})
	}
}

//update phone number  test
func TestUpdatePhoneNumber(t *testing.T) {
	randomUser := generateRandomUser()
	new_phone_number := utils.RandomePhoneNumber()
	account, err := adapter.CreateUser(context.Background(), randomUser)
	if err.Err != nil {
		t.Error(err)
	}
	testCase := []struct {
		name        string
		user        models.User
		phoneNumber string
		checker     func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:        "ok",
			user:        account,
			phoneNumber: new_phone_number,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.PhoneNumber, new_phone_number)
				require.Equal(t, user.Id, account.Id)

			},
		},
		{
			name:        "empty phone number",
			user:        models.User{Id: account.Id},
			phoneNumber: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
		{
			name:        "too short",
			user:        models.User{Id: account.Id},
			phoneNumber: "09751461",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ant, err := adapter.UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber)
			tc.checker(t, ant, err)
		})
	}
}

//update phone number  test
func TestUpdateEmail(t *testing.T) {
	randomUser := generateRandomUser()
	new_email := utils.RandomeEmail()
	account, err := adapter.CreateUser(context.Background(), randomUser)
	if err.Err != nil {
		t.Error(err)
	}
	testCase := []struct {
		name    string
		user    models.User
		email   string
		checker func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:  "ok",
			user:  account,
			email: new_email,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.Email, new_email)
				require.Equal(t, user.Id, account.Id)

			},
		},
		{
			name:  "empty email name",
			user:  models.User{Id: account.Id},
			email: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
		{
			name:  "invalid email format",
			user:  models.User{Id: account.Id},
			email: "thisisnotemail.com",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ant, err := adapter.UpdateUserEmail(context.Background(), tc.user, tc.email)
			tc.checker(t, ant, err)
		})
	}
}

//update phone number  test
func TestUpdatePassword(t *testing.T) {
	randomUser := generateRandomUser()
	new_password := utils.RandomPassword()
	account, err := adapter.CreateUser(context.Background(), randomUser)
	if err.Err != nil {
		t.Error(err)
	}
	testCase := []struct {
		name     string
		user     models.User
		password string
		checker  func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:     "ok",
			user:     account,
			password: new_password,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.Password, new_password)
				require.Equal(t, user.Id, account.Id)

			},
		},
		{
			name:     "empty email name",
			user:     models.User{Id: account.Id},
			password: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
		{
			name:     "short password",
			user:     models.User{Id: account.Id},
			password: "thisis",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ant, err := adapter.UpdateUserPassword(context.Background(), tc.user, tc.password)
			tc.checker(t, ant, err)
		})
	}
}

//testing get user by first name
func TestGetUserByFirstName(t *testing.T) {
	fistname := utils.RandomUserName()
	for i := 0; i < 10; i++ {
		email := utils.RandomeEmail()
		lastName := utils.RandomUserName()
		phoneNumber := utils.RandomePhoneNumber()
		user := models.User{
			FirstName:   fistname,
			LastName:    lastName,
			Email:       email,
			PhoneNumber: phoneNumber,
			Password:    utils.RandomPassword(),
		}
		adapter.CreateUser(context.Background(), user)

	}
	testCase := []struct {
		name      string
		firstName string
		checker   func(t *testing.T, users []models.User, err models.Errors)
	}{
		{
			name:      "ok",
			firstName: fistname,
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, len(users), 10)
				require.Equal(t, users[0].FirstName, fistname)

			},
		},
		{
			name:      "not found",
			firstName: utils.RandomUserName(),
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, users)

			},
		},
		{
			name:      "empty first name ",
			firstName: "",
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, users)

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			accounts, err := adapter.GetUsersByFirstName(context.Background(), tc.firstName)
			tc.checker(t, accounts, err)
		})
	}

}

//testing get user by last name
func TestGetUserByLastName(t *testing.T) {
	lastname := utils.RandomUserName()
	for i := 0; i < 10; i++ {
		email := utils.RandomeEmail()
		firstname := utils.RandomUserName()
		phoneNumber := utils.RandomePhoneNumber()
		user := models.User{
			FirstName:   firstname,
			LastName:    lastname,
			Email:       email,
			PhoneNumber: phoneNumber,
			Password:    utils.RandomPassword(),
		}
		adapter.CreateUser(context.Background(), user)

	}
	testCase := []struct {
		name     string
		lastName string
		checker  func(t *testing.T, users []models.User, err models.Errors)
	}{
		{
			name:     "ok",
			lastName: lastname,
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, len(users), 10)
				require.Equal(t, users[0].LastName, lastname)

			},
		},
		{
			name:     "not found",
			lastName: utils.RandomUserName(),
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, users)

			},
		},
		{
			name:     "empty first name ",
			lastName: "",
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, users)

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			accounts, err := adapter.GetUsersByLastName(context.Background(), tc.lastName)
			tc.checker(t, accounts, err)
		})
	}

}

//testing get user by phoneNumber
func TestGetUserByPhoneNumber(t *testing.T) {
	account := generateRandomUser()
	adapter.CreateUser(context.Background(), account)
	testCase := []struct {
		name        string
		phoneNumber string
		checker     func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:        "ok",
			phoneNumber: account.PhoneNumber,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.FirstName, account.FirstName)
				require.Equal(t, user.LastName, account.LastName)
				require.Equal(t, user.PhoneNumber, account.PhoneNumber)
				require.Equal(t, user.Email, account.Email)

			},
		},
		{
			name:        "not found",
			phoneNumber: utils.RandomUserName(),
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, user)

			},
		},
		{
			name:        "short phone number ",
			phoneNumber: "2131",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, user)

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			accounts, err := adapter.GetUserByPhoneNumber(context.Background(), tc.phoneNumber)
			tc.checker(t, accounts, err)
		})
	}

}
