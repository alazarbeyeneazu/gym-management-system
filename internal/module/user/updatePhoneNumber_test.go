package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

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
