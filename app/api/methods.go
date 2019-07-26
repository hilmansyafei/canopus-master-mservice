package api

import (
	"errors"
	"net/http"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"

	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
	"github.com/ztrue/tracerr"
	"gopkg.in/mgo.v2/bson"
)

// GetMethodsByID : Get data from methods collection
func (h *Handler) GetMethodsByID(c echo.Context) error {
	ID := c.Param("id")
	methods := &models.Methods{}
	if !bson.IsObjectIdHex(ID) {
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetMethodsByID] function",
			"Invalid ID format",
			"app/api/methods.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(errors.New("Invalid PID"))
	}

	err := h.Repositories.GetMethodByID(ID, methods)
	if err != nil {
		if err.Error() == "not found" {
			sSuccess := response.BuildSuccess("Data Not Found", status.DataNotFound)
			GenLog(c, "", sSuccess, "Response Log")
			c.JSON(http.StatusNotFound, sSuccess)
			return tracerr.Wrap(err)
		}
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetByPID] function",
			"Database Error",
			"app/api/condition.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}
	sSuccess := response.BuildSuccess(methods, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}
