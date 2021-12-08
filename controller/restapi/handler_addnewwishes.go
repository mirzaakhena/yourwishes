package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yourwishes/infrastructure/log"
	"yourwishes/infrastructure/util"
	"yourwishes/usecase/addnewwishes"
)

// addNewWishesHandler ...
func (r *Controller)addNewWishesHandler(inputPort addnewwishes.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := log.Context(c.Request.Context(), traceID)

		var req addnewwishes.InportRequest
		if err := c.BindJSON(&req); err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
			return
		}

		log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
			return
		}

		log.Info(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res, traceID))

	}
}
