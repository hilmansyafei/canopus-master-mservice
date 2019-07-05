package models

import (
	"github.com/hilmansyafei/go-package/database/mongo"
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
