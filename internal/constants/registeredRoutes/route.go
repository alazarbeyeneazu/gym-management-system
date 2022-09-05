package registeredroutes

import (
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/routers"
)

type RouterManager interface {
	RegisterRouters(route routers.Router)
}
type routerManger struct {
	routers []routers.Router
}

func NewRouterManager() RouterManager {
	return &routerManger{}
}
func (r *routerManger) RegisterRouters(route routers.Router) {
	r.routers = append(r.routers, route)
}
func (r *routerManger) GetRoutes() []routers.Router {

	return r.routers
}
