package models

import (
	"github.com/hilmansyafei/go-package/database/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Files holds data for merchants Collection
type Files struct {
	mongo.BaseStruct `json:",inline" bson:",inline"`
	Removed          bool   `json:"removed" bson:"removed"`
	Name             string `json:"name" bson:"name"`
	URI              string `json:"uri" bson:"uri"`
	Mime             string `json:"mime" bson:"mime"`
	Size             int32  `json:"size" bson:"size"`
	Status           int32  `json:"status" bson:"status"`
	Type             int32  `json:"type" bson:"type"`
}

// TableName : return table name
func (Files) TableName() string {
	return "files"
}

// ToModel : convert to model
func (Files) ToModel(data interface{}, model *Files) error {
	bsonBytes, err := bson.Marshal(data.(bson.M))
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bsonBytes, &model)
	if err != nil {
		return err
	}

	return nil
}

// ToModels : convert to model array
func (Files) ToModels(data interface{}, model *[]Files) error {
	bsonBytes, err := bson.Marshal(data.(bson.M))
	if err != nil {
		return err
	}
	err = bson.Unmarshal(bsonBytes, &model)
	if err != nil {
		return err
	}

	return nil
}
