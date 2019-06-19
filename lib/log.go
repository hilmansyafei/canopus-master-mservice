package lib

import (
	"github.com/alterra/canopus-master-mservice/config"
	"github.com/alterra/go-package/modules"
	"github.com/alterra/go-package/status"
	"github.com/labstack/echo"
)

// GenLog for general log
func GenLog(c echo.Context, dataRequest interface{}, resp interface{}, info string) {
	log := status.Log{
		IP:       c.RealIP(),
		Protocol: c.Request().Proto,
		Host:     c.Request().Host,
		URI:      c.Request().RequestURI,
		Headers:  c.Request().Header,
	}
	modules.GenLog(log, dataRequest, resp, info, config.App.AppConfig.String("log_path"))
}
