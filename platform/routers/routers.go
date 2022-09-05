package routers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/constants/models"
)

type Router struct {
	Method      string
	Path        string
	Domain      string
	MiddleWares []gin.HandlerFunc
	Handler     gin.HandlerFunc
}

type routing struct {
	serverAddress string
	routers       []Router
}

type Routers interface {
	Serve()
}

func Initialize(serverAddress string, routers []Router) Routers {
	return &routing{serverAddress: serverAddress, routers: routers}
}

func (r *routing) Serve() {
	// to store different groups of domains
	groups := make(map[string]string)

	// to store name and router gin router group
	routerGroup := make(map[string]*gin.RouterGroup)

	router := gin.New()
	// catogorize domain group in one if they are repeated
	for i, group := range r.routers {
		groups[r.routers[i].Domain] = group.Domain
	}

	//register groups to gine router
	for _, domain := range groups {
		if domain == "system" {
			routerGroup[domain] = router.Group("/api/v1/system")
		} else {
			// the result will be /api/suppliers/domain
			// or /api/users/domain
			spdomain := strings.Split(domain, "/")
			if len(spdomain) < 2 {
				log.Println(models.Errors{Err: errors.New("con not splite customers domain"), ErrorLocation: "platform/routers/routers.go", ErrLine: 53})
			}
			routerGroup[spdomain[0]] = router.Group(fmt.Sprintf("/api/v1/%s/%s", spdomain[0], spdomain[1]))
		}

	}
	//assign path and method for the requests
	for _, route := range r.routers {

		if route.Domain == "system" {
			method := route.Method
			switch method {
			case http.MethodPost:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].POST(route.Path, route.MiddleWares...)
				} else {

					routerGroup[route.Domain].POST(route.Path, route.Handler)
				}
			case http.MethodGet:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].GET(route.Path, route.MiddleWares...)
				} else {
					routerGroup[route.Domain].GET(route.Path, route.Handler)
				}

			case http.MethodPut:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].PUT(route.Path, route.MiddleWares...)
				} else {
					routerGroup[route.Domain].PUT(route.Path, route.Handler)
				}

			case http.MethodDelete:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].DELETE(route.Path, route.MiddleWares...)
				} else {
					routerGroup[route.Domain].DELETE(route.Path, route.Handler)
				}

			}

		} else {
			spdomain := strings.Split(route.Domain, "/")
			if len(spdomain) < 2 {
				log.Println(models.Errors{Err: errors.New("con not splite customers domain"), ErrorLocation: "platform/routers/routers.go", ErrLine: 53})
			}
			method := route.Method
			switch method {
			case http.MethodPost:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].POST(route.Path, route.MiddleWares...)
				} else {

					routerGroup[route.Domain].POST(route.Path, route.Handler)
				}
			case http.MethodGet:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].GET(route.Path, route.MiddleWares...)
				} else {
					routerGroup[route.Domain].GET(route.Path, route.Handler)
				}

			case http.MethodPut:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].PUT(route.Path, route.MiddleWares...)
				} else {
					routerGroup[route.Domain].PUT(route.Path, route.Handler)
				}

			case http.MethodDelete:
				if len(route.MiddleWares) > 0 {
					route.MiddleWares = append(route.MiddleWares, route.Handler)
					routerGroup[route.Domain].DELETE(route.Path, route.MiddleWares...)
				} else {
					routerGroup[route.Domain].DELETE(route.Path, route.Handler)
				}
			}
		}

	}

	router.Run(r.serverAddress)

}
