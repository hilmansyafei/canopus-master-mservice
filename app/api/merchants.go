package api

import (
	"net/http"

	"github.com/hilmansyafei/canopus-master-mservice/lib"

	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
)

// GetMerchantByIDHandler : api handler.
func (h *Handler) GetMerchantByIDHandler(c echo.Context) error {
	_id := c.Param("mid")
	merchants, err := GetMerchantByID(_id, h)
	if err != nil {
		// Database error
		sErr := response.NewErrorInfo("Canopus - Response: [GetMerchantByMID] function", "Database Error", "app/api/merchant.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}

	lib.GenLog(c, "", response.BuildSuccess(merchants, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(merchants, status.OKSuccess))
}
