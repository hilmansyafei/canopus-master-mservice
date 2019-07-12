package middleware

import (
	"strconv"
	"time"

	"github.com/hilmansyafei/canopus-master-mservice/app/api"
	"github.com/ztrue/tracerr"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// CustomHTTPErrorHandler : custom error handling
func CustomHTTPErrorHandler(err error, c echo.Context) {
	start := time.Now()
	microservice := "canopus-notification"
	duration := time.Since(start).Seconds() * 1000
	durationStr := strconv.FormatFloat(duration, 'f', -1, 64)

	logData := map[string]interface{}{
		"microservice": microservice,
		"duration":     durationStr,
		"tags":         []string{c.Request().RequestURI, c.Request().Method},
		"file":         tracerr.Sprint(err),
	}

	api.ZapGlobal.Error(
		err.Error(),
		zap.String("microservice", microservice),
		zap.String("duration", durationStr),
		zap.Strings("tags", []string{c.Request().RequestURI, c.Request().Method}),
		zap.String("stacktrace", tracerr.Sprint(err)),
	)
	api.LogGlobal.GenErrLog(logData, err.Error())
}
