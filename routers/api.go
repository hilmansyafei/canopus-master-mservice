package routers

import (
	"github.com/hilmansyafei/canopus-master-mservice/app/api"
	"github.com/hilmansyafei/canopus-master-mservice/app/middleware"

	"github.com/labstack/echo"
)

//Gen function generate routing
func Gen(env *api.Handler) *echo.Echo {
	//Initiate router
	e := echo.New()

	//Add router
	h := &api.Handler{DB: env.DB}
	e.GET("/getConditionByPID/:pid", h.GetConditionByPID, middleware.Track)
	e.GET("/getConditionByID/:id", h.GetConditionByID, middleware.Track)
	e.GET("/getConditionAll", h.GetConditionAll, middleware.Track)
	e.GET("/getCondition/:pid/:event", h.GetCondition, middleware.Track)
	e.GET("/getMerchantByID/:mid", h.GetMerchantByIDHandler, middleware.Track)

	return e
}
