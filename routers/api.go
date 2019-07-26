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

	// Condition router
	e.GET("/conditionByPID/:pid", h.GetConditionByPID, middleware.Track)
	e.GET("/conditionByID/:id", h.GetConditionByID, middleware.Track)
	e.GET("/conditionAll", h.GetConditionAll, middleware.Track)
	e.GET("/condition/:pid/:event", h.GetCondition, middleware.Track)

	// Merchant Routes
	e.GET("/merchant/:id", h.GetMerchantByIDHandler, middleware.Track)
	e.POST("/merchant", h.CreateMerchant, middleware.Track)
	e.DELETE("/merchant/:id", h.DeleteMerchantByID, middleware.Track)
	e.GET("/merchant", h.GetAllMerchant, middleware.Track)
	e.PUT("/merchant/:id", h.UpdateMerchant, middleware.Track)

	// Files Routes.
	e.GET("/getPathFile/:mid", h.GetPathFileHandler, middleware.Track)
	e.GET("/getAllFiles", h.GetAllFiles, middleware.Track)

	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler

	return e
}
