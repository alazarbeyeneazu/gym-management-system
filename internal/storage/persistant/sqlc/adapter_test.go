package db

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

var adapter *Adapter = NewAdapter("../../../../")

var email = utils.RandomeEmail()
var firstName = utils.RandomUserName()
var lastName = utils.RandomUserName()
var phoneNumber = utils.RandomePhoneNumber()

func TestCreateUser(t *testing.T) {

	testCase := []struct {
		name        string
		account     models.User
		checkResult func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name: "ok",
			account: models.User{
				FirstName:   firstName,
				LastName:    lastName,
				PhoneNumber: phoneNumber,
				Email:       email,
				Password:    utils.RandomPassword(),
			},
			checkResult: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.NotEmpty(t, user)
				require.NotEmpty(t, user.FirstName)
				require.NotEmpty(t, user.LastName)
				require.NotEmpty(t, user.Email)
				require.NotEmpty(t, user.State)
				require.Empty(t, user.Password)
				require.Equal(t, user.Email, email)
				require.Equal(t, user.FirstName, firstName)
				require.Equal(t, user.LastName, lastName)
				require.Equal(t, phoneNumber, user.PhoneNumber)

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

func TestDeleteUser(t *testing.T) {
	account, err := adapter.CreateUser(context.Background(), models.User{
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Email:       email,
		Password:    utils.RandomPassword(),
	})

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