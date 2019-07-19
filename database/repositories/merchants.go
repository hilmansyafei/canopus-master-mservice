package repositories

import (
	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"
)

// GetMerchantByID : Get data from merchant collection by ID
func (Hnd Env) GetMerchantByID(id string) (models.Merchants, error) {
	var getData interface{}
	merchants := models.Merchants{}
	err := Hnd.Mp.GetByID(merchants.TableName(), bson.ObjectIdHex(id), &getData)
	if err != nil {
		return merchants, err
	}
	err = merchants.ToModel(getData, &merchants)
	if err != nil {
		return merchants, err
	}
	return merchants, nil
}

// GetMerchantByMID : Get data from merchant collection by MID
func (Hnd Env) GetMerchantByMID(mid string) (models.Merchants, error) {
	var getData interface{}
	merchants := models.Merchants{}
	queryGetData := bson.M{"mid": mid}
	err := Hnd.Mp.GetOne(merchants.TableName(), queryGetData, &getData)
	if err != nil {
		return merchants, err
	}
	errMdl := merchants.ToModel(getData, &merchants)
	if errMdl != nil {
		return merchants, errMdl
	}
	return merchants, nil
}
