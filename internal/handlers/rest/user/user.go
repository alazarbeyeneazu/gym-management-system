package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/module/user"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/routers"
)

type UserHandler interface {
	RegisterUser(ctx *gin.Context)
	StartRoutes() []routers.Router
}
type userHandler struct {
	userApp user.UserService
	routes  []routers.Router
}

func Init(user user.UserService) UserHandler {

	return &userHandler{userApp: user}
}

//validate user inpute
func validateUser(user models.User, functionName string) error {

	switch functionName {
	case "RegisterUser":
		err := validation.ValidateStruct(
			&user,
			validation.Field(&user.FirstName, validation.Required, validation.Length(2, 100), is.Alpha),
			validation.Field(&user.LastName, validation.Required, validation.Length(2, 100), is.Alpha),
			validation.Field(&user.Email, is.Email, validation.Required),
			validation.Field(&user.Password, validation.Required, validation.Length(8, 1000)),
			validation.Field(&user.PhoneNumber, validation.Required),
		)
		return err
	}
	return nil
}

func (u *userHandler) RegisterUser(ctx *gin.Context) {

	var newUser models.User
	if err := ctx.ShouldBind(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.RestResponse{
			Error:  errors.New("bad request").Error(),
			Status: "not registered",
			User:   models.User{},
		})
		return
	}
	if len(newUser.PhoneNumber) == 10 {
		phoneArray := newUser.PhoneNumber[1:]
		phoneNumber := "+251" + phoneArray
		newUser.PhoneNumber = phoneNumber

	}
	err := validateUser(newUser, "RegisterUser")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.RestResponse{
			Error:  err.Error(),
			Status: "not registered",
			User:   models.User{},
		})
		return
	}

	user, errmodel := u.userApp.RegisterUser(ctx, newUser)
	if errmodel.Err != nil {

		var duperr error

		if errmodel.Err.Error() == "pq: duplicate key value violates unique constraint \"users_phone_number_key\"" {
			duperr = errors.New("phone number already registered")
		} else if errmodel.Err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			duperr = errors.New("email already registered")
		} else {
			duperr = errors.New("internal error")
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.RestResponse{
			Error:  duperr.Error(),
			Status: "not registered",
			User:   user,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, models.RestResponse{
		Error:  "",
		Status: "registered",
		User:   user,
	})

}
