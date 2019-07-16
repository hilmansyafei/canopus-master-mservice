package api

import (
	"github.com/hilmansyafei/canopus-master-mservice/database/repositories"
	"github.com/hilmansyafei/go-package/database/mongo"
	"github.com/hilmansyafei/go-package/modules"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// LogGlobal : Global Log
var LogGlobal modules.LogProvider

// ZapGlobal : global zap
var ZapGlobal *zap.Logger

// Handler connection database
type Handler struct {
	MongoProvider mongo.MongoProvider
	Repositories  repositories.Repositories
}

// GenLog : Generate general error
func GenLog(c echo.Context, dataRequest interface{}, resp interface{}, info string) {
	log := status.Log{
		IP:       c.RealIP(),
		Protocol: c.Request().Proto,
		Host:     c.Request().Host,
		URI:      c.Request().RequestURI,
		Headers:  c.Request().Header,
	}
	LogGlobal.GenLog(log, dataRequest, resp, info)
}
