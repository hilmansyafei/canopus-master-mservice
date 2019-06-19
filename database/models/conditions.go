package models

import (
	"github.com/zebresel-com/mongodm"
	"gopkg.in/mgo.v2/bson"
)

// Conditions holds data for Conditions Collection
type Conditions struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	PID       bson.ObjectId `json:"pid" bson:"pid"`
	Key       string        `json:"key" bson:"key"`
	Name      string        `json:"name" bson:"name"`
	Event     string        `json:"event" bson:"event"`
	Weight    string        `json:"weight" bson:"weight"`
	Module    string        `json:"module" bson:"module"`
	Settings  []interface{} `json:"settings" bson:"settings"`
	Revisions []interface{} `json:"revisions" bson:"revisions"`
}
