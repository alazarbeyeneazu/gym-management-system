package user

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

func TestDeleteUser(t *testing.T) {
	appUser := Initiate()
	randomUser := generateRandomUser()
	account, err := appUser.RegisterUser(context.Background(), randomUser)

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
			err := appUser.DeleteUser(context.Background(), tc.user)
			tc.checker(t, err)
		})
	}

}
