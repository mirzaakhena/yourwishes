package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yourwishes/infrastructure/log"
	"yourwishes/infrastructure/util"
	"yourwishes/usecase/displayallwishes"
)

// displayAllWishesHandler ...
func (r *Controller) displayAllWishesHandler(inputPort displayallwishes.Inport) gin.HandlerFunc {

	type Wishes struct {
		Message string `json:"message"`
		Date    string `json:"date"`
		ID      string `json:"id"`
	}

	type response struct {
		AllWishes []Wishes `json:"all_wishes"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := log.Context(c.Request.Context(), traceID)

		var req displayallwishes.InportRequest
		//if err := c.BindJSON(&req); err != nil {
		//	log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
		//	return
		//}

		log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response

		for _, w := range res.ListOfWishes {
			jsonRes.AllWishes = append(jsonRes.AllWishes, Wishes{
				Message: w.Message,
				Date:    w.Date.Format("2006-01-02 15:04:05"),
				ID:      w.ID,
			})
		}

		log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, NewSuccessResponse(jsonRes, traceID))

	}
}
