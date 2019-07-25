package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hilmansyafei/canopus-master-mservice/app/api"
	"github.com/hilmansyafei/canopus-master-mservice/database/repositories"
	"github.com/hilmansyafei/go-package/database/mongo"
	"github.com/hilmansyafei/go-package/modules"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func measure(lP modules.LogProvider, mP mongo.MongoProvider) *api.Handler {
	api.ZapGlobal, _ = zap.NewProduction()
	api.LogGlobal = lP
	repoProvider := repositories.Env{
		Mp: mP,
	}
	return &api.Handler{
		MongoProvider: mP,
		Repositories:  repoProvider,
	}
}

func Test_GetMethodsByID_success(t *testing.T) {
	data := bson.M{}
	dbMock := mongo.MongoMock{
		InterfaceReturn: data,
	}

	mockLogInit, _ := test.NewNullLogger()
	mockLog := &modules.Log{
		Logger: mockLogInit,
	}

	env := measure(mockLog, &dbMock)
	// Setup Mock Request.
	e := echo.New()

	// Create mock Error Body.
	requestBodyError := httptest.NewRequest(http.MethodPost, "/", errReader(0))
	recorder := httptest.NewRecorder()
	echoContext := e.NewContext(requestBodyError, recorder)
	echoContext.SetPath("/methods/:id")
	echoContext.SetParamNames("id")
	echoContext.SetParamValues("4fb41cf8e688128b6e714bf4")

	// Assertions.
	if assert.NoError(t, env.GetMethodsByID(echoContext)) {
		// Error code should be 200.
		assert.Equal(t, http.StatusOK, recorder.Code)
		// Define expected error.
		errorResponse := `{"data":{"_id":"","createdAt":"","updatedAt":"","mid":"","key":"","name":"","desc":"","status":0,"env":0,"module":"","expiredTime":"","settings":null,"Revisions":null,"type":0},"status":{"type":"Success","code":200,"message":"Request Successed"}}`
		assert.Equal(t, errorResponse, strings.TrimSpace(recorder.Body.String()))
	}
}
