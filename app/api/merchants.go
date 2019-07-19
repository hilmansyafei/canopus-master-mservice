package api

import (
	"errors"
	"net/http"

	"github.com/ztrue/tracerr"
	"gopkg.in/mgo.v2/bson"

	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
)

// GetMerchantByIDHandler : api handler.
func (h *Handler) GetMerchantByIDHandler(c echo.Context) error {
	ID := c.Param("mid")
	if !bson.IsObjectIdHex(ID) {
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetMerchantByIDHandler] function",
			"Invalid PID format",
			"app/api/merchants.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(errors.New("Invalid PID format"))
	}
	merchants, err := h.Repositories.GetMerchantByID(ID)
	if err != nil {
		if err.Error() == "not found" {
			sSuccess := response.BuildSuccess("Data Not Found", status.DataNotFound)
			GenLog(c, "", sSuccess, "Response Log")
			c.JSON(http.StatusNotFound, sSuccess)
			return tracerr.Wrap(err)
		}
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetMerchantByMID] function",
			"Database Error",
			"app/api/merchant.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}

	sSuccess := response.BuildSuccess(merchants, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}
