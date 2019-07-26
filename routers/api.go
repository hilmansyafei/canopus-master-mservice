package routers

import (
	"github.com/hilmansyafei/canopus-master-mservice/app/api"
	"github.com/hilmansyafei/canopus-master-mservice/app/middleware"

	"github.com/labstack/echo"
)

//Gen function generate routing
func Gen(h *api.Handler) *echo.Echo {
	//Initiate router
	e := echo.New()

	//Add router
	e.GET("/conditionByPID/:pid", h.GetConditionByPID, middleware.Track)
	e.GET("/conditionByID/:id", h.GetConditionByID, middleware.Track)
	e.GET("/conditionAll", h.GetConditionAll, middleware.Track)
	e.GET("/condition/:pid/:event", h.GetCondition, middleware.Track)
	e.GET("/merchantByID/:mid", h.GetMerchantByIDHandler, middleware.Track)

	// Files Routes.
	e.GET("/getPathFile/:mid", h.GetPathFileHandler, middleware.Track)

	// Methods Routes.
	e.GET("/methods/:id", h.GetMethodsByID, middleware.Track)
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler

	return e
}
