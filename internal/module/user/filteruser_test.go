package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

func TestGetUserByFirstName(t *testing.T) {
	appUser := Initiate()
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
		appUser.RegisterUser(context.Background(), user)

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
			accounts, err := appUser.GetUsersByFirstName(context.Background(), tc.firstName)
			tc.checker(t, accounts, err)
		})
	}

}

func TestGetUserByLastName(t *testing.T) {
	appUser := Initiate()
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
		appUser.RegisterUser(context.Background(), user)

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
			accounts, err := appUser.GetUsersByLastName(context.Background(), tc.lastName)
			tc.checker(t, accounts, err)
		})
	}

}

func TestGetUserByPhoneNumber(t *testing.T) {
	appUser := Initiate()
	account := generateRandomUser()
	appUser.RegisterUser(context.Background(), account)
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
			accounts, err := appUser.GetUserByPhoneNumber(context.Background(), tc.phoneNumber)
			tc.checker(t, accounts, err)
		})
	}

}
