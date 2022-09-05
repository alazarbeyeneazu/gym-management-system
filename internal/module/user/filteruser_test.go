package user

import (
	"context"
	"database/sql"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	mockdb "gitlab.com/2ftimeplc/2fbackend/delivery-1/mocks/db"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

func TestGetUserByFirstName(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)

	testCase := []struct {
		name      string
		firstName string
		checker   func(t *testing.T, users []models.User, err models.Errors)
	}{
		{
			name:      "ok",
			firstName: "firstnamehere",
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, len(users), 10)
				require.Equal(t, users[0].FirstName, "firstnamehere")

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
			name:      "empty first name",
			firstName: "",
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, users)
				require.Equal(t, err.Err, errors.New("first name cannot be blank"))

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":

				var user []models.User
				for i := 0; i < 10; i++ {
					user = append(user, models.User{Id: rand.Int63(), FirstName: tc.firstName, LastName: utils.RandomUserName(), PhoneNumber: utils.RandomePhoneNumber(), Email: utils.RandomeEmail(), CreatedAt: time.Now().GoString(), State: 1})
				}
				db.EXPECT().GetUsersByFirstName(context.Background(), tc.firstName).Return(user, models.Errors{})
				defer ctl.Finish()
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUsersByFirstName(context.Background(), tc.firstName)
				tc.checker(t, accounts, err)
			case "not found":
				defer ctl.Finish()
				db.EXPECT().GetUsersByFirstName(context.Background(), tc.firstName).Return([]models.User{}, models.Errors{Err: sql.ErrNoRows})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUsersByFirstName(context.Background(), tc.firstName)
				tc.checker(t, accounts, err)

			case "empty first name":
				defer ctl.Finish()
				db.EXPECT().GetUsersByFirstName(context.Background(), tc.firstName).Return([]models.User{}, models.Errors{Err: errors.New("not empty")})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUsersByFirstName(context.Background(), tc.firstName)
				tc.checker(t, accounts, err)
			}

		})
	}

}

func TestGetUserByLastName(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)

	testCase := []struct {
		name     string
		lastName string
		checker  func(t *testing.T, users []models.User, err models.Errors)
	}{
		{
			name:     "ok",
			lastName: "lastnamehere",
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, len(users), 10)
				require.Equal(t, users[0].LastName, "lastnamehere")

			},
		},
		{
			name:     "not found",
			lastName: utils.RandomUserName(),
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, users)
				require.EqualError(t, err.Err, sql.ErrNoRows.Error())

			},
		},
		{
			name:     "empty last name",
			lastName: "",
			checker: func(t *testing.T, users []models.User, err models.Errors) {
				require.Empty(t, users)
				require.EqualError(t, err.Err, "LastName cannot be blank")

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				var user []models.User
				for i := 0; i < 10; i++ {
					user = append(user, models.User{Id: rand.Int63(), FirstName: utils.RandomUserName(), LastName: tc.lastName, PhoneNumber: utils.RandomePhoneNumber(), Email: utils.RandomeEmail(), CreatedAt: time.Now().GoString(), State: 1})
				}
				db.EXPECT().GetUsersByLastName(context.Background(), tc.lastName).Return(user, models.Errors{})
				defer ctl.Finish()
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUsersByLastName(context.Background(), tc.lastName)
				tc.checker(t, accounts, err)
			case "not found":
				defer ctl.Finish()
				db.EXPECT().GetUsersByLastName(context.Background(), tc.lastName).Return([]models.User{}, models.Errors{Err: sql.ErrNoRows})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUsersByLastName(context.Background(), tc.lastName)
				tc.checker(t, accounts, err)

			case "empty last name":
				defer ctl.Finish()
				db.EXPECT().GetUsersByLastName(context.Background(), tc.lastName).Return([]models.User{}, models.Errors{Err: errors.New("unexpected error")})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUsersByLastName(context.Background(), tc.lastName)
				tc.checker(t, accounts, err)
			}
		})
	}

}

func TestGetUserByPhoneNumber(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	account := generateRandomUser()
	shortPhoneNumber := utils.RandomePhoneNumber()
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
			phoneNumber: shortPhoneNumber,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, user)
				require.Error(t, err.Err)
				require.EqualError(t, err.Err, sql.ErrNoRows.Error())

			},
		},
		{
			name:        "short phone number",
			phoneNumber: "2131",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, user)
				require.EqualError(t, err.Err, "the length must be exactly 13")

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			switch tc.name {
			case "ok":
				defer ctl.Finish()
				db.EXPECT().GetUserByPhoneNumber(context.Background(), tc.phoneNumber).Return(account, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUserByPhoneNumber(context.Background(), tc.phoneNumber)
				tc.checker(t, accounts, err)
			case "not found":
				defer ctl.Finish()
				db.EXPECT().GetUserByPhoneNumber(context.Background(), tc.phoneNumber).Return(models.User{}, models.Errors{Err: sql.ErrNoRows})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUserByPhoneNumber(context.Background(), tc.phoneNumber)
				tc.checker(t, accounts, err)
			case "short phone number":
				defer ctl.Finish()
				db.EXPECT().GetUserByPhoneNumber(context.Background(), tc.phoneNumber).Return(models.User{}, models.Errors{Err: errors.New("un expected error")})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUserByPhoneNumber(context.Background(), tc.phoneNumber)
				tc.checker(t, accounts, err)
			}

		})
	}

}

func TestGetUserByEmail(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	account := generateRandomUser()

	testCase := []struct {
		name    string
		email   string
		checker func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:  "ok",
			email: account.Email,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.FirstName, account.FirstName)
				require.Equal(t, user.LastName, account.LastName)
				require.Equal(t, user.PhoneNumber, account.PhoneNumber)
				require.Equal(t, user.Email, account.Email)

			},
		},
		{
			name:  "not found",
			email: utils.RandomeEmail(),
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, user)
				require.Error(t, err.Err)
				require.EqualError(t, err.Err, sql.ErrNoRows.Error())
			},
		},
		{
			name:  "invalid email",
			email: "2131.com",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, user)
				require.Error(t, err.Err)
				require.EqualError(t, err.Err, "must be a valid email address")

			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				defer ctl.Finish()
				db.EXPECT().GetUserByEmail(context.Background(), tc.email).Return(account, models.Errors{})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUserByEmail(context.Background(), tc.email)
				tc.checker(t, accounts, err)
			case "not found":
				defer ctl.Finish()
				db.EXPECT().GetUserByEmail(context.Background(), tc.email).Return(models.User{}, models.Errors{Err: sql.ErrNoRows})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUserByEmail(context.Background(), tc.email)
				tc.checker(t, accounts, err)
			case "invalid email":
				defer ctl.Finish()
				db.EXPECT().GetUserByEmail(context.Background(), tc.email).Return(models.User{}, models.Errors{Err: errors.New("no error")})
				appUser := Initiate("../../../", db)
				accounts, err := appUser.GetUserByEmail(context.Background(), tc.email)
				tc.checker(t, accounts, err)
			}

		})
	}

}
