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
	e.GET("/get_data_condition/:pid", h.GetByPID, middleware.Track)

	return e
}
