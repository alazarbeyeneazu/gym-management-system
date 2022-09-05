package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/middlewares"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/routers"
)

func (u *userHandler) StartRoutes() []routers.Router {

	registerUser := []routers.Router{
		{
			Method:      http.MethodPost,
			Path:        "account",
			Domain:      "system",
			Handler:     u.RegisterUser,
			MiddleWares: []gin.HandlerFunc{middlewares.TestMiddleWare()},
		},
	}

	u.routes = registerUser

	return u.routes

}
