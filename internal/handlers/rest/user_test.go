package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	app "gitlab.com/2ftimeplc/2fbackend/delivery-1/mocks/app"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/utils"
)

func CreateJsonParams(user models.User) []byte {

	body, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func TestResgisterUser(t *testing.T) {

	mockController := gomock.NewController(t)
	service := app.NewMockUserService(mockController)
	hdler := Init(service)
	router := gin.New()
	url := "/register"

	router.POST("/register", hdler.RegisterUser)

	testUser := models.User{
		Id:          1,
		FirstName:   utils.RandomUserName(),
		LastName:    utils.RandomUserName(),
		PhoneNumber: utils.RandomePhoneNumber(),
		Email:       utils.RandomeEmail(),
		Password:    utils.RandomPassword(),
		CreatedAt:   time.Now().GoString(),
	}

	testCase := []struct {
		name    string
		user    models.User
		checker func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "ok",
			user: testUser,
			checker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var user models.RestResponse
				require.Equal(t, recorder.Code, http.StatusOK)
				body, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				json.Unmarshal(body, &user)
				require.Equal(t, user.User, testUser)
			},
		}, {
			name: "empty first Name",
			user: models.User{
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Email:       utils.RandomeEmail(),
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var user models.RestResponse
				require.Equal(t, recorder.Code, http.StatusBadRequest)
				body, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				json.Unmarshal(body, &user)
				require.Empty(t, user.User)
				require.Equal(t, user.Error, "first_name: cannot be blank.")

			},
		}, {
			name: "empty Last Name",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Email:       utils.RandomeEmail(),
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var user models.RestResponse
				require.Equal(t, recorder.Code, http.StatusBadRequest)
				body, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				json.Unmarshal(body, &user)
				require.Empty(t, user.User)
				require.Equal(t, user.Error, "last_name: cannot be blank.")
			},
		}, {
			name: "empty PhoneNumber",
			user: models.User{
				FirstName: utils.RandomUserName(),
				LastName:  utils.RandomUserName(),
				Email:     utils.RandomeEmail(),
				Password:  utils.RandomPassword(),
			},
			checker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var user models.RestResponse
				require.Equal(t, recorder.Code, http.StatusBadRequest)
				body, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				json.Unmarshal(body, &user)
				require.Empty(t, user.User)
				require.Equal(t, user.Error, "phone_number: cannot be blank.")
			},
		}, {
			name: "empty password",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Email:       utils.RandomeEmail(),
			},
			checker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var user models.RestResponse
				require.Equal(t, recorder.Code, http.StatusBadRequest)
				body, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				json.Unmarshal(body, &user)
				require.Empty(t, user.User)
				require.Equal(t, user.Error, "password: cannot be blank.")
			},
		}, {
			name: "short password",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				LastName:    utils.RandomUserName(),
				PhoneNumber: utils.RandomePhoneNumber(),
				Email:       utils.RandomeEmail(),
				Password:    "hello",
			},
			checker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var user models.RestResponse
				require.Equal(t, recorder.Code, http.StatusBadRequest)
				body, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				json.Unmarshal(body, &user)
				require.Empty(t, user.User)
				require.Equal(t, user.Error, "password: the length must be between 8 and 1000.")
			},
		}, {
			name: "invalide PhoneNumber",
			user: models.User{
				FirstName:   utils.RandomUserName(),
				LastName:    utils.RandomUserName(),
				PhoneNumber: "8238383",
				Email:       utils.RandomeEmail(),
				Password:    utils.RandomPassword(),
			},
			checker: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				var user models.RestResponse
				require.Equal(t, recorder.Code, http.StatusBadRequest)
				body, err := ioutil.ReadAll(recorder.Body)
				require.NoError(t, err)
				json.Unmarshal(body, &user)
				require.Empty(t, user.User)
				require.Equal(t, user.Error, "phone_number: the length must be exactly 13.")
			},
		},
	}

	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "ok":
				service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(testUser, models.Errors{})
				defer mockController.Finish()
				json_data, err := json.Marshal(tc.user)
				if err != nil {
					log.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
				request.Header.Set("Content-Type", "application/json")
				require.NoError(t, err)
				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				tc.checker(t, recorder)
			case "empty first Name":
				service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer mockController.Finish()
				json_data, err := json.Marshal(tc.user)
				if err != nil {
					log.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
				request.Header.Set("Content-Type", "application/json")
				require.NoError(t, err)

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, request)

				tc.checker(t, recorder)
			case "empty PhoneNumber":
				service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer mockController.Finish()
				json_data, err := json.Marshal(tc.user)
				if err != nil {
					log.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
				request.Header.Set("Content-Type", "application/json")
				require.NoError(t, err)

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				tc.checker(t, recorder)
			case "empty Last Name":
				service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer mockController.Finish()
				json_data, err := json.Marshal(tc.user)
				if err != nil {
					log.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
				request.Header.Set("Content-Type", "application/json")
				require.NoError(t, err)

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				tc.checker(t, recorder)
			case "invalide PhoneNumber":
				service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer mockController.Finish()
				json_data, err := json.Marshal(tc.user)
				if err != nil {
					log.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
				request.Header.Set("Content-Type", "application/json")
				require.NoError(t, err)

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				tc.checker(t, recorder)
			case "empty password":
				service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer mockController.Finish()
				json_data, err := json.Marshal(tc.user)
				if err != nil {
					log.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
				request.Header.Set("Content-Type", "application/json")
				require.NoError(t, err)

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				tc.checker(t, recorder)
			case "short password":
				service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(models.User{}, models.Errors{Err: errors.New("unexpected error")})
				defer mockController.Finish()
				json_data, err := json.Marshal(tc.user)
				if err != nil {
					log.Fatal(err)
				}

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
				request.Header.Set("Content-Type", "application/json")
				require.NoError(t, err)

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, request)
				tc.checker(t, recorder)
			}

		})
	}

}
