package api

import (
	"errors"

	"encoding/json"
	"net/http"

	"github.com/ztrue/tracerr"
	"gopkg.in/mgo.v2/bson"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"github.com/hilmansyafei/go-package/response"
	"github.com/hilmansyafei/go-package/status"
	"github.com/labstack/echo"
)

// GetMerchantByIDHandler : api handler get data by Id from merchants collection
func (h *Handler) GetMerchantByIDHandler(c echo.Context) error {
	_id := c.Param("id")
	if !bson.IsObjectIdHex(_id) {
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetMerchantByIDHandler] function",
			"Invalid PID format",
			"app/api/merchants.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(errors.New("Invalid PID format"))
	}

	merchants, err := h.Repositories.GetMerchantByID(_id)
	if err != nil {
		if err.Error() == "not found" {
			sSuccess := response.BuildSuccess("Merchant Not Found", status.OKSuccess)
			GenLog(c, "", sSuccess, "Response Log")
			c.JSON(http.StatusNotFound, sSuccess)
			return nil
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

// GetAllMerchant : api handler to get all data from merchants collection
func (h *Handler) GetAllMerchant(c echo.Context) error {
	var merchants []interface{}
	err := h.Repositories.GetAllMerchant(&merchants)

	if err != nil {
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetAll] function",
			"Database Error",
			"app/api/merchants.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}

	sSuccess := response.BuildSuccess(merchants, status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}

// CreateMerchant : api handler.
func (h *Handler) CreateMerchant(c echo.Context) error {

	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return err
	}

	statusInt := int(jsonMap["Status"].(float64))
	statusInt32 := int32(statusInt)

	env32 := int(jsonMap["ENV"].(float64))
	thisEnv := int32(env32)

	//json_map has the JSON Payload decoded into a map
	data := models.Merchants{
		ShortName:        jsonMap["ShortName"].(string),
		OfficialName:     jsonMap["OfficialName"].(string),
		Email:            jsonMap["Email"].(string),
		Status:           statusInt32,
		MID:              jsonMap["MID"].(string),
		ENV:              thisEnv,
		SecretKey:        jsonMap["SecretKey"].([]interface{}),
		MerchantPubKey:   bson.ObjectIdHex(jsonMap["MerchantPubKey"].(string)),
		PsaPrivKey:       bson.ObjectIdHex(jsonMap["PsaPrivKey"].(string)),
		PsaPubKey:        bson.ObjectIdHex(jsonMap["PsaPubKey"].(string)), // bson.NewObjectId(),
		Method:           jsonMap["Method"].([]interface{}),
		NotificationURLs: jsonMap["NotificationURLs"].(interface{}),
		WhitelistIP:      jsonMap["WhitelistIP"].(interface{}),
	}

	err = h.Repositories.CreateMerchant(data)
	if err != nil {
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [CreateMerchant] function",
			"Database Error",
			"app/api/merchant.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}
	sSuccess := response.BuildSuccess("Successfully create merchant", status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}

// UpdateMerchant : api handler.
func (h *Handler) UpdateMerchant(c echo.Context) error {
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return err
	}
	statusInt := int(jsonMap["Status"].(float64))
	statusInt32 := int32(statusInt)

	env32 := int(jsonMap["ENV"].(float64))
	thisEnv := int32(env32)

	//json_map has the JSON Payload decoded into a map
	dataUpdate := models.Merchants{
		ShortName:        jsonMap["ShortName"].(string),
		OfficialName:     jsonMap["OfficialName"].(string),
		Email:            jsonMap["Email"].(string),
		Status:           statusInt32,
		MID:              jsonMap["MID"].(string),
		ENV:              thisEnv,
		SecretKey:        jsonMap["SecretKey"].([]interface{}),
		MerchantPubKey:   bson.ObjectIdHex(jsonMap["MerchantPubKey"].(string)),
		PsaPrivKey:       bson.ObjectIdHex(jsonMap["PsaPrivKey"].(string)),
		PsaPubKey:        bson.ObjectIdHex(jsonMap["PsaPubKey"].(string)), // bson.NewObjectId(),
		Method:           jsonMap["Method"].([]interface{}),
		NotificationURLs: jsonMap["NotificationURLs"].(interface{}),
		WhitelistIP:      jsonMap["WhitelistIP"].(interface{}),
	}
	_id := c.Param("id")
	err3 := h.Repositories.UpdateMerchant(_id, dataUpdate)

	if err3 != nil {
		if err3.Error() == "not found" {
			sSuccess := response.BuildSuccess("Data Not Found", status.OKSuccess)
			GenLog(c, "", sSuccess, "Response Log")
			c.JSON(http.StatusNotFound, sSuccess)
			return nil
		}
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [UpdateMerchant] function",
			"Database Error",
			"app/api/merchant.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}
	sSuccess := response.BuildSuccess("Successfully update merchant", status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}

// DeleteMerchantByID : api handler to delete data from merchants collection
func (h *Handler) DeleteMerchantByID(c echo.Context) error {
	_id := c.Param("id")
	err := h.Repositories.DeleteMerchantByID(_id)
	if err != nil {
		if err.Error() == "not found" {
			sSuccess := response.BuildSuccess("Merchant Not Found", status.OKSuccess)
			GenLog(c, "", sSuccess, "Response Log")
			c.JSON(http.StatusNotFound, sSuccess)
			return nil
		}
		// Database error
		sErr := response.BuildError(response.NewErrorInfo(
			"Canopus - Response: [GetMerchantByID] function",
			"Database Error",
			"app/api/merchant.go"), status.InternalServerError)
		GenLog(c, "", sErr, "Error Log")
		c.JSON(http.StatusInternalServerError, sErr)
		return tracerr.Wrap(err)
	}

	sSuccess := response.BuildSuccess("Successfully delete merchant", status.OKSuccess)
	GenLog(c, "", sSuccess, "Response Log")
	c.JSON(http.StatusOK, sSuccess)
	return nil
}
