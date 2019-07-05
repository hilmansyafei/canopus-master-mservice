package api

import (
	"fmt"
	"net/http"

	"github.com/hilmansyafei/canopus-master-mservice/lib"
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
		sErr := response.NewErrorInfo("Canopus - Response: [GetPathFileHandler] function", err.Error(), "app/api/files.go")
		lib.GenLog(c, "", response.BuildError(sErr, status.InternalServerError), "Error Log")
		return c.JSON(http.StatusInternalServerError, response.BuildError(sErr, status.InternalServerError))
	}
	// Return valid response.
	lib.GenLog(c, "", response.BuildSuccess(path, status.OKSuccess), "Response Log")
	return c.JSON(http.StatusOK, response.BuildSuccess(path, status.OKSuccess))
}
