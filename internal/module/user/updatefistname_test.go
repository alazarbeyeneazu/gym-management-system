package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

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
