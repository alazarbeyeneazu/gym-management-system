package rest

import (
	"bytes"
	"encoding/json"
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
	router := gin.New()
	
	testUser := models.User{
		Id:          1,
		FirstName:   utils.RandomUserName(),
		LastName:    utils.RandomUserName(),
		PhoneNumber: utils.RandomePhoneNumber(),
		Email:       utils.RandomeEmail(),
		Password:    utils.RandomPassword(),
		CreatedAt:   time.Now().GoString(),
	}
	mockController := gomock.NewController(t)
	service := app.NewMockUserService(mockController)
	service.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(testUser, models.Errors{})
	defer mockController.Finish()

	hdler := Init(service)

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
				t.Error("first name ", user.User.FirstName)

			},
		},
	}

	for _, tc := range testCase {

		t.Run(tc.name, func(t *testing.T) {
			json_data, err := json.Marshal(tc.user)
			if err != nil {
				log.Fatal(err)
			}
			url := "/register"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(string(json_data)))
			request.Header.Set("Content-Type", "application/json")
			require.NoError(t, err)
			router.POST("/register", hdler.RegisterUser)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)
			tc.checker(t, recorder)

		})
	}

}
