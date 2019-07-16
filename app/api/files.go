package api

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/ztrue/tracerr"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"
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

// GetPathFile : Get path file by MID and type file
// data["type"] : private / public
func GetPathFile(data map[string]string, h *Handler) (models.Files, error) {
	// Get merchant data.
	merchants, err := h.Repositories.GetMerchantByMID(data["mid"])
	// Define files struct.
	files := models.Files{}
	// Check error.
	if err != nil {
		return files, err
	}
	// Define idFile as string.
	var idFile string
	// Check type in data.
	switch data["type"] {
	// If want to get private key.
	case "private":
		idFile = hex.EncodeToString([]byte(merchants.PsaPrivKey))
	// If want to get public key.
	case "public":
		idFile = hex.EncodeToString([]byte(merchants.MerchantPubKey))
	// Return error.
	default:
		return files, errors.New("File type not found")
	}
	// Get path uri.
	Path, errPath := h.Repositories.GetPathFileByID(idFile)
	return Path, errPath
}
