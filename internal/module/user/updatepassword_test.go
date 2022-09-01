package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

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
