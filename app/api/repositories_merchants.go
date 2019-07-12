package api

import (
	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"
)

// GetMerchantByID : Get data from merchant collection by ID
func GetMerchantByID(id string) (models.Merchants, error) {
	var getData interface{}
	merchants := models.Merchants{}
	err := MongoProvider.GetByID(merchants.TableName(), bson.ObjectIdHex(id), &getData)
	err = merchants.ToModel(getData, &merchants)
	if err != nil {
		return merchants, err
	}
	return merchants, nil
}

// GetMerchantByMID : Get data from merchant collection by MID
func GetMerchantByMID(mid string) (models.Merchants, error) {
	var getData interface{}
	merchants := models.Merchants{}
	queryGetData := bson.M{"mid": mid}
	err := MongoProvider.GetOne(merchants.TableName(), queryGetData, &getData)
	err = merchants.ToModel(getData, &merchants)
	if err != nil {
		return merchants, err
	}
	return merchants, nil
}
