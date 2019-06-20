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
		sErr := response.NewErrorInfo("Canopus - Response: [GetByPID] function", "Invalid PID format", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	err := Conditions.Find(queryGetData).Exec(&conditions)
	if err != nil {
		// Database error
		sErr := response.NewErrorInfo("Canopus - Response: [GetByPID] function", "Database Error", "app/api/condition.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}
	lib.GenLog(c, "", response.BuildSuccess(conditions, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(conditions, status.OKSuccess))
}
