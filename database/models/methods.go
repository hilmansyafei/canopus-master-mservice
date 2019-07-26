package models

import (
	"github.com/hilmansyafei/go-package/database/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Methods holds data for methods Collection
type Methods struct {
	mongo.BaseStruct `json:",inline" bson:",inline"`
	MID              bson.ObjectId `json:"mid" bson:"mid"`
	Key              string        `json:"key" bson:"key"`
	Name             string        `json:"name" bson:"name"`
	DESC             string        `json:"desc" bson:"desc"`
	Status           int32         `json:"status" bson:"status"`
	ENV              int32         `json:"env" bson:"env"`
	Module           string        `json:"module" bson:"module"`
	ExpiredTime      string        `json:"expiredTime" bson:"expiredTime"`
	Settings         interface{}   `json:"settings" bson:"settings"`
	Revisions        []interface{} `json:"Revisions" bson:"Revisions"`
	Type             int32         `json:"type" bson:"type"`
}

// TableName : return table name
func (Methods) TableName() string {
	return "methods"
}

// ToModel : convert to model
func (Methods) ToModel(data interface{}, model *Methods) error {
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
