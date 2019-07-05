package api

import (
	"errors"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"
)

// GetMerchantByID : Get data from merchant collection by ID
func GetMerchantByID(id string, h *Handler) (models.Merchants, error) {
	Merchant := h.DB.C("merchants")
	merchants := models.Merchants{}
	if bson.IsObjectIdHex(id) {
		queryGetData := bson.M{"_id": bson.ObjectIdHex(id)}
		err := Merchant.Find(queryGetData).One(&merchants)
		return merchants, err
	}
	return merchants, errors.New("Invalid ID")
}

// GetMerchantByMID : Get data from merchant collection by MID
func GetMerchantByMID(mid string, h *Handler) (models.Merchants, error) {
	Merchant := h.DB.C("merchants")
	merchants := models.Merchants{}
	queryGetData := bson.M{"mid": mid}
	err := Merchant.Find(queryGetData).One(&merchants)
	return merchants, err
}
