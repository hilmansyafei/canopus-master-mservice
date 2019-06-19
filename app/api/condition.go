package api

import (
	"net/http"

	"github.com/alterra/canopus-master-mservice/database/models"
	"github.com/alterra/canopus-master-mservice/lib"
	"gopkg.in/mgo.v2/bson"

	"github.com/alterra/go-package/status"
	"github.com/labstack/echo"
)

// GetByPID : Get data from conditions collection
func (h *Handler) GetByPID(c echo.Context) error {
	pid := c.Param("pid")
	h.DB.Register(&models.Conditions{}, "conditions")
	Conditions := h.DB.Model("conditions")
	conditions := []*models.Conditions{}
	queryGetData := bson.M{}

	if bson.IsObjectIdHex(pid) {
		queryGetData["pid"] = bson.ObjectIdHex(pid)
	} else {
		sErr := status.NewSingleErrors("Canopus - Response: [GetByPID] function", "Invalid PID format", "app/api/condition.go")
		lib.GenLog(c, "", status.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, status.BuildError(sErr, status.InternalServerError))
	}

	err := Conditions.Find(queryGetData).Exec(&conditions)
	if err != nil {
		// Database error
		sErr := status.NewSingleErrors("Canopus - Response: [GetByPID] function", "Database Error", "app/api/condition.go")
		lib.GenLog(c, "", status.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, status.BuildError(sErr, status.InternalServerError))
	}
	lib.GenLog(c, "", status.BuildSuccess(conditions, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, status.BuildSuccess(conditions, status.OKSuccess))
}
