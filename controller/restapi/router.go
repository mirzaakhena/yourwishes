package restapi

import (
	"yourwishes/usecase/addnewwishes"
	"yourwishes/usecase/displayallwishes"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router                 gin.IRouter
	AddNewWishesInport     addnewwishes.Inport
	DisplayAllWishesInport displayallwishes.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/wishes", r.authorized(), r.addNewWishesHandler(r.AddNewWishesInport))
	r.Router.GET("/wishes", r.authorized(), r.displayAllWishesHandler(r.DisplayAllWishesInport))
}
