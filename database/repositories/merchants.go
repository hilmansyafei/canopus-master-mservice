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

// GetAllMerchant : get all merchant
func (Hnd Env) GetAllMerchant(merchants *[]interface{}) error {
	err := Hnd.Mp.GetAll("merchants", merchants)
	if err != nil {
		return err
	}
	return nil
}

// CreateMerchant : Create data to merchant collection
func (Hnd Env) CreateMerchant(data models.Merchants) error {
	_, _, err := Hnd.Mp.Create("merchants", data, data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateMerchant : Update data from merchant collection
func (Hnd Env) UpdateMerchant(id string, data models.Merchants) error {
	_id := bson.ObjectIdHex(id)
	queryGetData := bson.M{"_id": _id}
	err := Hnd.Mp.Update("merchants", queryGetData, data)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMerchantByID : Delete data from merchant collection by ID
func (Hnd Env) DeleteMerchantByID(id string) error {
	merchants := models.Merchants{}
	err := Hnd.Mp.DeleteID(merchants.TableName(), bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}
