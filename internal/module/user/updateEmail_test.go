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
