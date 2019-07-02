package api

import (
	"net/http"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"github.com/hilmansyafei/canopus-master-mservice/lib"
	"gopkg.in/mgo.v2/bson"

	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
)

// GetConditionByPID : Get data from conditions collection
func (h *Handler) GetConditionByPID(c echo.Context) error {
	pid := c.Param("pid")
	Conditions := h.DB.C("conditions")
	conditions := []models.Conditions{}
	queryGetData := bson.M{}

	if bson.IsObjectIdHex(pid) {
		queryGetData["pid"] = bson.ObjectIdHex(pid)
	} else {
		sErr := response.NewErrorInfo("Canopus - Response: [GetByPID] function", "Invalid PID format", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	err := Conditions.Find(queryGetData).All(&conditions)
	if err != nil {
		// Database error
		sErr := response.NewErrorInfo("Canopus - Response: [GetByPID] function", "Database Error", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	lib.GenLog(c, "", response.BuildSuccess(conditions, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(conditions, status.OKSuccess))
}

// GetConditionByID : Get data from conditions collection
func (h *Handler) GetConditionByID(c echo.Context) error {
	id := c.Param("id")
	Conditions := h.DB.C("conditions")
	conditions := models.Conditions{}
	queryGetData := bson.M{}

	if bson.IsObjectIdHex(id) {
		queryGetData["_id"] = bson.ObjectIdHex(id)
	} else {
		sErr := response.NewErrorInfo("Canopus - Response: [GetByID] function", "Invalid ID format", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	err := Conditions.Find(queryGetData).One(&conditions)
	if err != nil {
		// Database error
		sErr := response.NewErrorInfo("Canopus - Response: [GetByID] function", "Database Error", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	lib.GenLog(c, "", response.BuildSuccess(conditions, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(conditions, status.OKSuccess))
}

// GetConditionAll : Get data from conditions collection
func (h *Handler) GetConditionAll(c echo.Context) error {
	Conditions := h.DB.C("conditions")
	conditions := []models.Conditions{}

	err := Conditions.Find(nil).All(&conditions)
	if err != nil {
		// Database error
		sErr := response.NewErrorInfo("Canopus - Response: [GetAll] function", "Database Error", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	lib.GenLog(c, "", response.BuildSuccess(conditions, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(conditions, status.OKSuccess))
}

// GetCondition : Get data for conditions check
func (h *Handler) GetCondition(c echo.Context) error {
	pid := c.Param("pid")
	event := c.Param("event")
	Conditions := h.DB.C("conditions")
	conditions := models.Conditions{}
	queryGetData := bson.M{}

	if bson.IsObjectIdHex(pid) {
		queryGetData["pid"] = bson.ObjectIdHex(pid)
	} else {
		sErr := response.NewErrorInfo("Canopus - Response: [GetCondition] function", "Invalid PID format", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	if event != "" {
		queryGetData["event"] = event
	}

	err := Conditions.Find(queryGetData).One(&conditions)
	if err != nil {
		// Database error
		sErr := response.NewErrorInfo("Canopus - Response: [GetCondition] function", "Database Error", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	lib.GenLog(c, "", response.BuildSuccess(conditions, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(conditions, status.OKSuccess))
}
