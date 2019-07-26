package repositories

import (
	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"github.com/hilmansyafei/go-package/database/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Repositories : custon query
type Repositories interface {
	GetMerchantByID(id string) (models.Merchants, error)
	GetMerchantByMID(mid string) (models.Merchants, error)
	GetPathFileByID(id string) (models.Files, error)
	GetConditionByPID(pid string, conditions *[]interface{}) error
	GetConditionByID(id bson.ObjectId, condition *models.Conditions) error
	GetAllCondition(conditions *[]interface{}) error
	GetConditionEvent(query bson.M, conditions *models.Conditions) error
	GetMethodByID(id string, methods *models.Methods) error
}

// Env : ENV data
type Env struct {
	Mp mongo.MongoProvider
}
