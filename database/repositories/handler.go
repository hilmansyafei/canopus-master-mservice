package repositories

import (
	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"github.com/hilmansyafei/go-package/database/mongo"
	"gopkg.in/mgo.v2/bson"
)

// Repositories : custon query
type Repositories interface {

	// Define Merchant Repositories
	GetMerchantByID(id string) (models.Merchants, error)
	GetMerchantByMID(mid string) (models.Merchants, error)
	DeleteMerchantByID(id string) error
	GetAllMerchant(merchants *[]interface{}) error
	CreateMerchant(models.Merchants) error
	UpdateMerchant(id string, data models.Merchants) error

	// Define Files Repositories
	GetPathFileByID(id string) (models.Files, error)
	GetAllFiles(files *[]interface{}) error

	// Define Condition Repositories
	GetConditionByPID(pid string, conditions *[]interface{}) error
	GetConditionByID(id bson.ObjectId, condition *models.Conditions) error
	GetAllCondition(conditions *[]interface{}) error
	GetConditionEvent(query bson.M, conditions *models.Conditions) error
}

// Env : ENV data
type Env struct {
	Mp mongo.MongoProvider
	// PagingQuery mongo.PagingQuery
}
