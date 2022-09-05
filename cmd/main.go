package main

import (
	handler "gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/handlers/rest/user"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/module/user"
	db "gitlab.com/2ftimeplc/2fbackend/delivery-1/internal/storage/persistant/sqlc"
	"gitlab.com/2ftimeplc/2fbackend/delivery-1/platform/routers"
)

func main() {

	dbs := db.NewAdapter("../")
	service := user.Initiate("../", dbs)
	user := handler.Init(service)
	routes := user.StartRoutes()
	router := routers.Initialize(":8181", routes)
	router.Serve()
}
