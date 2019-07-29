package repositories

import (
	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"
)

// GetMethodByID : Get data from method collection by ID
func (Hnd Env) GetMethodByID(id string, methods *models.Methods) error {
	var getData interface{}
	err := Hnd.Mp.GetByID(methods.TableName(), bson.ObjectIdHex(id), &getData)
	if err != nil {
		return err
	}
	err = methods.ToModel(getData, methods)
	if err != nil {
		return err
	}
	return nil
}
