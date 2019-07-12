package api

import (
	"fmt"
	"net/http"

	"github.com/ztrue/tracerr"

	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
)

// GetPathFileHandler : api handler.
func (h *Handler) GetPathFileHandler(c echo.Context) error {
	// Get mid in param.
	mid := c.Param("mid")
	// Get query param.
	typeFile := c.QueryParam("type")
	// Bundle to interface.
	dataQuery := map[string]string{
		"mid":  mid,
		"type": typeFile,
	}
	fmt.Println(dataQuery)
	// Get file.
	path, err := GetPathFile(dataQuery, h)
	if err != nil {
		// Database error.
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetPathFileHandler] function",
			"Error get path file",
			"app/api/files.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}
	// Return valid response.
	sSuccess := response.BuildSuccess(path, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}
