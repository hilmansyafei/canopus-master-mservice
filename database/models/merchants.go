package models

import (
	"github.com/hilmansyafei/go-package/database/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Merchants holds data for merchants Collection
type Merchants struct {
	mongo.BaseStruct `json:",inline" bson:",inline"`
	ShortName        string        `json:"shortName" bson:"shortName"`
	OfficialName     string        `json:"officialName" bson:"officialName"`
	Email            string        `json:"email" bson:"email"`
	Status           int32         `json:"status" bson:"status"`
	MID              string        `json:"mid" bson:"mid"`
	ENV              int32         `json:"env" bson:"env"`
	SecretKey        []interface{} `json:"secretKey" bson:"secretKey"`
	MerchantPubKey   bson.ObjectId `json:"merchantPubKey" bson:"merchantPubKey"`
	PsaPrivKey       bson.ObjectId `json:"psaPrivKey" bson:"psaPrivKey"`
	PsaPubKey        bson.ObjectId `json:"psaPubKey" bson:"psaPubKey"`
	Method           []interface{} `json:"method" bson:"method"`
	NotificationURLs interface{}   `json:"notificationURLs" bson:"notificationURLs"`
	WhitelistIP      interface{}   `json:"whitelistIP" bson:"whitelistIP"`
}
