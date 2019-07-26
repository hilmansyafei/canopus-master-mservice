package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/hilmansyafei/canopus-master-mservice/app/api"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// Track : track log path request
func Track(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		microservice := "canopus-notification"
		data := make(map[string]interface{})
		body, _ := ioutil.ReadAll(c.Request().Body)
		json.Unmarshal([]byte(body), &data)
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		msg := "Incoming Request"
		timeLog := time.Now().Format(time.RFC3339Nano)
		logData := map[string]interface{}{
			"msg":          msg,
			"microservice": microservice,
			"tags":         []string{c.Request().RequestURI, c.Request().Method},
			"Body":         data,
			"time":         timeLog,
		}
		api.ZapGlobal.Info(
			msg,
			zap.String("microservice", microservice),
			zap.Strings("tags", []string{c.Request().RequestURI, c.Request().Method}),
			zap.Any("Body", data),
			zap.String("time", timeLog),
		)
		api.LogGlobal.GenAnyLog(logData, msg)
		return next(c)
	}
}
