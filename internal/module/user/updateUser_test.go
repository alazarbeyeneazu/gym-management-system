package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	mockdb "gitlab.com/2ftimeplc/2fbackend/delivery-1/mocks/db"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

func TestUpdateEmail(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	randomUser := generateRandomUser()
	new_email := utils.RandomeEmail()

	testCase := []struct {
		name    string
		user    models.User
		email   string
		checker func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:  "ok",
			user:  randomUser,
			email: new_email,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.Email, new_email)
				require.Equal(t, user.Id, randomUser.Id)

			},
		},
		{
			name:  "empty email name",
			user:  models.User{Id: randomUser.Id},
			email: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.EqualError(t, err.Err, errors.New("Email  cannot be blank").Error())
				require.Empty(t, user)

			},
		},
		{
			name:  "invalid email format",
			user:  models.User{Id: randomUser.Id},
			email: "thisisnotemail.com",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "Email  must be a valid email address")

			},
		},
		{
			name:  "Empty Id",
			user:  models.User{},
			email: "test@gmail.com",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "Id  cannot be blank")

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				updatedUser := models.User{Id: randomUser.Id, FirstName: randomUser.FirstName, LastName: randomUser.LastName, PhoneNumber: randomUser.PhoneNumber, Email: tc.email, CreatedAt: randomUser.CreatedAt}
				db.EXPECT().UpdateUserEmail(context.Background(), tc.user, tc.email).Return(updatedUser, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserEmail(context.Background(), tc.user, tc.email)
				tc.checker(t, ant, err)
			case "empty email name":
				db.EXPECT().UpdateUserEmail(context.Background(), tc.user, tc.email).Return(models.User{}, models.Errors{Err: errors.New("un expected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserEmail(context.Background(), tc.user, tc.email)
				tc.checker(t, ant, err)
			case "invalid email format":
				db.EXPECT().UpdateUserEmail(context.Background(), tc.user, tc.email).Return(models.User{}, models.Errors{Err: errors.New("un expected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserEmail(context.Background(), tc.user, tc.email)
				tc.checker(t, ant, err)
			case "Empty Id":
				db.EXPECT().UpdateUserEmail(context.Background(), tc.user, tc.email).Return(models.User{}, models.Errors{Err: errors.New("un expected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserEmail(context.Background(), tc.user, tc.email)
				tc.checker(t, ant, err)
			}

		})
	}
}

func TestUpdatePhoneNumber(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	randomUser := generateRandomUser()
	new_phone_number := utils.RandomePhoneNumber()

	testCase := []struct {
		name        string
		user        models.User
		phoneNumber string
		checker     func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:        "ok",
			user:        randomUser,
			phoneNumber: new_phone_number,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.PhoneNumber, new_phone_number)
				require.Equal(t, user.Id, randomUser.Id)

			},
		},
		{
			name:        "empty phone number",
			user:        models.User{Id: randomUser.Id},
			phoneNumber: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "phone number  cannot be blank")

			},
		},
		{
			name:        "too short",
			user:        models.User{Id: randomUser.Id},
			phoneNumber: "09751461",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "phone number  the length must be exactly 13")

			},
		},
		{
			name:        "empty id",
			user:        models.User{},
			phoneNumber: "+251975146165",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "Id  cannot be blank")

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				updatedUser := models.User{Id: randomUser.Id, FirstName: randomUser.FirstName, LastName: randomUser.LastName, PhoneNumber: tc.phoneNumber, Email: randomUser.Email, CreatedAt: randomUser.CreatedAt}
				db.EXPECT().UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber).Return(updatedUser, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber)
				tc.checker(t, ant, err)
			case "empty phone number":
				db.EXPECT().UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber)
				tc.checker(t, ant, err)
			case "too short":
				db.EXPECT().UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber)
				tc.checker(t, ant, err)
			case "empty id":
				db.EXPECT().UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserPhoneNumber(context.Background(), tc.user, tc.phoneNumber)
				tc.checker(t, ant, err)

			}

		})
	}
}

func TestUpdatePassword(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	randomUser := generateRandomUser()
	new_password := utils.RandomPassword()

	testCase := []struct {
		name     string
		user     models.User
		password string
		checker  func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:     "ok",
			user:     randomUser,
			password: new_password,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.Id, randomUser.Id)

			},
		},
		{
			name:     "empty password",
			user:     models.User{Id: randomUser.Id},
			password: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "password  cannot be blank")

			},
		},
		{
			name:     "short password",
			user:     models.User{Id: randomUser.Id},
			password: "thisis",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "password  the length must be between 8 and 1000")

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				updatedUser := models.User{Id: randomUser.Id, FirstName: randomUser.FirstName, LastName: randomUser.LastName, PhoneNumber: randomUser.PhoneNumber, Email: randomUser.Email, CreatedAt: randomUser.CreatedAt}
				db.EXPECT().UpdateUserPassword(context.Background(), tc.user, tc.password).Return(updatedUser, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserPassword(context.Background(), tc.user, tc.password)
				tc.checker(t, ant, err)
			case "empty password":
				db.EXPECT().UpdateUserPassword(context.Background(), tc.user, tc.password).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserPassword(context.Background(), tc.user, tc.password)
				tc.checker(t, ant, err)
			case "short password":
				db.EXPECT().UpdateUserPassword(context.Background(), tc.user, tc.password).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserPassword(context.Background(), tc.user, tc.password)
				tc.checker(t, ant, err)

			}

		})
	}
}

func TestUpdateLastName(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	randomUser := generateRandomUser()
	new_Name := utils.RandomUserName()

	testCase := []struct {
		name    string
		user    models.User
		newName string
		checker func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:    "ok",
			user:    randomUser,
			newName: new_Name,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.LastName, new_Name)
				require.Equal(t, user.Id, randomUser.Id)

			},
		},
		{
			name:    "empty Last name",
			user:    models.User{Id: randomUser.Id},
			newName: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "Last Name  cannot be blank")

			},
		},
		{
			name:    "too short",
			user:    models.User{Id: randomUser.Id},
			newName: "h",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "Last Name  the length must be between 2 and 100")

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				updatedUser := models.User{Id: randomUser.Id, FirstName: randomUser.FirstName, LastName: tc.newName, PhoneNumber: randomUser.PhoneNumber, Email: randomUser.Email, CreatedAt: randomUser.CreatedAt}
				db.EXPECT().UpdateUserLastName(context.Background(), tc.user, tc.newName).Return(updatedUser, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserLastName(context.Background(), tc.user, tc.newName)
				tc.checker(t, ant, err)
			case "empty Last name":
				db.EXPECT().UpdateUserLastName(context.Background(), tc.user, tc.newName).Return(models.User{}, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserLastName(context.Background(), tc.user, tc.newName)
				tc.checker(t, ant, err)
			case "too short":
				db.EXPECT().UpdateUserLastName(context.Background(), tc.user, tc.newName).Return(models.User{}, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserLastName(context.Background(), tc.user, tc.newName)
				tc.checker(t, ant, err)

			}

		})
	}
}

func TestUpdateFirstName(t *testing.T) {
	ctl := gomock.NewController(t)
	db := mockdb.NewMockDBPort(ctl)
	randomUser := generateRandomUser()
	new_Name := utils.RandomUserName()

	testCase := []struct {
		name    string
		user    models.User
		newName string
		checker func(t *testing.T, user models.User, err models.Errors)
	}{
		{
			name:    "ok",
			user:    randomUser,
			newName: new_Name,
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.Empty(t, err)
				require.Equal(t, user.FirstName, new_Name)
				require.Equal(t, user.Id, randomUser.Id)

			},
		},
		{
			name:    "empty first name",
			user:    models.User{Id: randomUser.Id},
			newName: "",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "First Name  cannot be blank")

			},
		},
		{
			name:    "too short",
			user:    models.User{Id: randomUser.Id},
			newName: "h",
			checker: func(t *testing.T, user models.User, err models.Errors) {
				require.NotEmpty(t, err)
				require.Error(t, err.Err)
				require.Empty(t, user)
				require.EqualError(t, err.Err, "First Name  the length must be between 2 and 100")

			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				updatedUser := models.User{Id: randomUser.Id, FirstName: tc.newName, LastName: randomUser.LastName, PhoneNumber: randomUser.PhoneNumber, Email: randomUser.Email, CreatedAt: randomUser.CreatedAt}
				db.EXPECT().UpdateUserFirstName(context.Background(), tc.user, tc.newName).Return(updatedUser, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserFirstName(context.Background(), tc.user, tc.newName)
				tc.checker(t, ant, err)
			case "empty first name":
				db.EXPECT().UpdateUserFirstName(context.Background(), tc.user, tc.newName).Return(models.User{}, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserFirstName(context.Background(), tc.user, tc.newName)
				tc.checker(t, ant, err)
			case "too short":
				db.EXPECT().UpdateUserFirstName(context.Background(), tc.user, tc.newName).Return(models.User{}, models.Errors{})
				defer ctl.Finish()

				appUser := Initiate("../../../", db)
				ant, err := appUser.UpdateUserFirstName(context.Background(), tc.user, tc.newName)
				tc.checker(t, ant, err)

			}

		})
	}
}
