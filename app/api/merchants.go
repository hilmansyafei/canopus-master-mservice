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

// GetMerchantByID : Get data from merchant collection
func (h *Handler) GetMerchantByID(c echo.Context) error {
	_id := c.Param("mid")
	Merchant := h.DB.C("merchants")
	merchants := models.Merchants{}
	queryGetData := bson.M{"_id": bson.ObjectIdHex(_id)}

	err := Merchant.Find(queryGetData).One(&merchants)
	if err != nil {
		// Database error
		sErr := response.NewErrorInfo("Canopus - Response: [GetMerchantByMID] function", "Database Error", "app/api/merchant.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	lib.GenLog(c, "", response.BuildSuccess(merchants, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(merchants, status.OKSuccess))
}
