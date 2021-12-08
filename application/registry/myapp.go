package registry

import (
	"yourwishes/application"
	"yourwishes/controller"
	"yourwishes/controller/restapi"
	"yourwishes/gateway/localdb"
	"yourwishes/infrastructure/server"
	"yourwishes/usecase/addnewwishes"
	"yourwishes/usecase/displayallwishes"
)

type myapp struct {
	*server.GinHTTPHandler
	controller.Controller
}

func NewMyapp() func() application.RegistryContract {
	return func() application.RegistryContract {

		httpHandler := server.NewGinHTTPHandlerDefault()

		datasource := localdb.NewGateway()

		return &myapp{
			GinHTTPHandler: &httpHandler,
			Controller: &restapi.Controller{
				Router:                 httpHandler.Router,
				AddNewWishesInport:     addnewwishes.NewUsecase(datasource),
				DisplayAllWishesInport: displayallwishes.NewUsecase(datasource),
			},
		}

	}
}
