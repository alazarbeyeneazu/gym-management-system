package user

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	mockdb "gitlab.com/2ftimeplc/2fbackend/delivery-1/mocks/db"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

func generateRandomUser() models.User {
	var email = utils.RandomeEmail()
	var firstName = utils.RandomUserName()
	var lastName = utils.RandomUserName()
	var phoneNumber = utils.RandomePhoneNumber()
	return models.User{
		Id:          rand.Int63(),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Email:       email,
		State:       1,
		CreatedAt:   time.Now().GoString(),
		Password:    utils.RandomPassword(),
	}

}

func TestRegisterUser(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
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
				require.Equal(t, err.ErrorLocation, "/internal/module/user/registration.go")
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
				require.Equal(t, err.ErrorLocation, "/internal/module/user/registration.go")
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
				require.Equal(t, err.ErrorLocation, "/internal/module/user/registration.go")
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
				require.Equal(t, err.ErrorLocation, "/internal/module/user/registration.go")
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
				require.Equal(t, err.ErrorLocation, "/internal/module/user/registration.go")
				require.Empty(t, user)
			},
		},
	}

	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				defer ctl.Finish()
				db.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(randomUser, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.RegisterUser(context.Background(), tc.account)
				tc.checkResult(t, accounts, err)
			case "empty first name":
				defer ctl.Finish()
				db.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(randomUser, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.RegisterUser(context.Background(), tc.account)
				tc.checkResult(t, accounts, err)
			case "empty last name":
				defer ctl.Finish()
				db.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(randomUser, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.RegisterUser(context.Background(), tc.account)
				tc.checkResult(t, accounts, err)
			case "Empty PhoneNumber":
				defer ctl.Finish()
				db.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(randomUser, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.RegisterUser(context.Background(), tc.account)
				tc.checkResult(t, accounts, err)
			case "empty email":
				defer ctl.Finish()
				db.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(randomUser, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.RegisterUser(context.Background(), tc.account)
				tc.checkResult(t, accounts, err)
			case "empty password":
				defer ctl.Finish()
				db.EXPECT().CreateUser(context.Background(), gomock.Any()).Return(randomUser, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.RegisterUser(context.Background(), tc.account)
				tc.checkResult(t, accounts, err)

			}
		})
	}

}
