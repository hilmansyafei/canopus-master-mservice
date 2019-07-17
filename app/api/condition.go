package api

import (
	"errors"
	"net/http"

	"github.com/ztrue/tracerr"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"

	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
)

// GetConditionByPID : Get data from conditions collection
func (h *Handler) GetConditionByPID(c echo.Context) error {
	pid := c.Param("pid")
	var conditions []interface{}

	if !bson.IsObjectIdHex(pid) {
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetByPID] function",
			"Invalid PID format",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(errors.New("Invalid PID"))
	}

	err := h.Repositories.GetConditionByPID(pid, &conditions)
	if err != nil {
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetByPID] function",
			"Database Error",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}

	sSuccess := response.BuildSuccess(conditions, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}

// GetConditionByID : Get data from conditions collection
func (h *Handler) GetConditionByID(c echo.Context) error {
	id := c.Param("id")
	conditions := models.Conditions{}

	if !bson.IsObjectIdHex(id) {
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetByID] function",
			"Invalid ID format",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(errors.New("Invalid ID format"))
	}

	err := h.Repositories.GetConditionByID(bson.ObjectIdHex(id), &conditions)
	if err != nil {
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetByID] function",
			"Database Error",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}

	sSuccess := response.BuildSuccess(conditions, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	return c.JSON(http.StatusOK, sSuccess)
}

// GetConditionAll : Get data from conditions collection
func (h *Handler) GetConditionAll(c echo.Context) error {
	var conditions []interface{}
	err := h.Repositories.GetAllCondition(&conditions)
	if err != nil {
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetAll] function",
			"Database Error",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}

	sSuccess := response.BuildSuccess(conditions, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}

// GetCondition : Get data for conditions check
func (h *Handler) GetCondition(c echo.Context) error {
	pid := c.Param("pid")
	event := c.Param("event")
	conditions := models.Conditions{}
	queryGetData := bson.M{}

	if bson.IsObjectIdHex(pid) {
		queryGetData["pid"] = bson.ObjectIdHex(pid)
	} else {
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetCondition] function",
			"Invalid PID format",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(errors.New("Invalid PID format"))
	}

	if event != "" {
		queryGetData["event"] = event
	}

	err := h.Repositories.GetConditionEvent(queryGetData, &conditions)
	if err != nil {
		if err.Error() == "not found" {
			sSuccess := response.BuildSuccess("Data Not Found", status.OKSuccess)
			GenLog(c, "", sSuccess, "Response Log")
			c.JSON(http.StatusNotFound, sSuccess)
			return nil
		}
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetCondition] function",
			"Database Error",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}

	sSuccess := response.BuildSuccess(conditions, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}
