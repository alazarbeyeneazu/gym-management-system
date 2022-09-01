package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

func TestUpdateEmail(t *testing.T) {
	appUser := Initiate()
	randomUser := generateRandomUser()
	new_email := utils.RandomeEmail()
	account, err := appUser.RegisterUser(context.Background(), randomUser)
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
			ant, err := appUser.UpdateUserEmail(context.Background(), tc.user, tc.email)
			tc.checker(t, ant, err)
		})
	}
}

func TestUpdatePhoneNumber(t *testing.T) {
	appUser := Initiate()
	randomUser := generateRandomUser()
	new_phone_number := utils.RandomePhoneNumber()
	account, err := appUser.RegisterUser(context.Background(), randomUser)
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
			ant, err := appUser.UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber)
			tc.checker(t, ant, err)
		})
	}
}

func TestUpdatePassword(t *testing.T) {
	appUser := Initiate()
	randomUser := generateRandomUser()
	new_password := utils.RandomPassword()
	account, err := appUser.RegisterUser(context.Background(), randomUser)
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
			ant, err := appUser.UpdateUserPassword(context.Background(), tc.user, tc.password)
			tc.checker(t, ant, err)
		})
	}
}

func TestUpdateLastName(t *testing.T) {
	appUser := Initiate()
	randomUser := generateRandomUser()
	new_Name := utils.RandomUserName()
	account, err := appUser.RegisterUser(context.Background(), randomUser)
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
			ant, err := appUser.UpdateUserLastName(context.Background(), tc.user, tc.newName)
			tc.checker(t, ant, err)
		})
	}
}

func TestUpdateFirstName(t *testing.T) {
	appUser := Initiate()
	randomUser := generateRandomUser()
	new_Name := utils.RandomUserName()
	account, err := appUser.RegisterUser(context.Background(), randomUser)
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
			ant, err := appUser.UpdateUserFirstName(context.Background(), tc.user, tc.newName)
			tc.checker(t, ant, err)
		})
	}
}
