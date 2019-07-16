package api

import (
	"net/http"

	"github.com/ztrue/tracerr"

	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
)

// GetMerchantByIDHandler : api handler.
func (h *Handler) GetMerchantByIDHandler(c echo.Context) error {
	_id := c.Param("mid")
	merchants, err := h.Repositories.GetMerchantByID(_id)
	if err != nil {
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
