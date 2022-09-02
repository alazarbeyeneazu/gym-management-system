package user

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	mockdb "gitlab.com/2ftimeplc/2fbackend/delivery-1/mocks/db"
)

func TestDeleteUser(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	randomUser := generateRandomUser()
	testCase := []struct {
		name    string
		user    models.User
		checker func(t *testing.T, err models.Errors)
	}{
		{
			name: "ok",
			user: randomUser,
			checker: func(t *testing.T, err models.Errors) {
				require.Empty(t, err)

			},
		},
		{
			name: "not found",
			user: models.User{
				Id: rand.Int63(),
			},
			checker: func(t *testing.T, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Equal(t, err.Err, sql.ErrNoRows)

			},
		},
		{
			name: "empty id",
			user: models.User{},
			checker: func(t *testing.T, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Equal(t, err.Err.Error(), "cannot be blank")

			},
		},
	}
	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {

			switch tc.name {
			case "ok":
				db.EXPECT().DeleteUser(context.Background(), tc.user).Return(models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				err := appUser.DeleteUser(context.Background(), tc.user)
				tc.checker(t, err)
			case "not found":
				db.EXPECT().DeleteUser(context.Background(), tc.user).Return(models.Errors{Err: sql.ErrNoRows})
				defer ctl.Finish()
				appUser := Initiate("../../../", db)
				err := appUser.DeleteUser(context.Background(), tc.user)
				tc.checker(t, err)
			case "empty id":
				db.EXPECT().DeleteUser(context.Background(), tc.user).Return(models.Errors{Err: sql.ErrNoRows})
				defer ctl.Finish()
				appUser := Initiate("../../../", db)
				err := appUser.DeleteUser(context.Background(), tc.user)
				tc.checker(t, err)

			}

		})
	}

}
