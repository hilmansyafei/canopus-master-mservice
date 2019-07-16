package repositories

import (
	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"
)

// GetConditionByPID : get condition by PID
func (Hnd Env) GetConditionByPID(pid string, conditions *[]interface{}) error {
	queryGetData := bson.M{"pid": bson.ObjectIdHex(pid)}
	err := Hnd.Mp.Get("conditions", queryGetData, conditions)
	if err != nil {
		return err
	}
	return nil
}

// GetConditionByID : get condition by ID
func (Hnd Env) GetConditionByID(id bson.ObjectId, condition *models.Conditions) error {
	var getData interface{}
	err := Hnd.Mp.GetByID("conditions", id, &getData)
	if err != nil {
		return err
	}
	err = condition.ToModel(getData, condition)
	if err != nil {
		return err
	}

	return nil
}

// GetAllCondition : get all condition
func (Hnd Env) GetAllCondition(conditions *[]interface{}) error {
	err := Hnd.Mp.GetAll("conditions", conditions)
	if err != nil {
		return err
	}
	return nil
}

// GetConditionEvent : get condition by PID and event
func (Hnd Env) GetConditionEvent(query bson.M, conditions *models.Conditions) error {
	var getData interface{}
	err := Hnd.Mp.GetOne("conditions", query, &getData)
	if err != nil {
		return err
	}
	condition := models.Conditions{}
	err = condition.ToModel(getData, conditions)
	if err != nil {
		return err
	}

	return nil
}
