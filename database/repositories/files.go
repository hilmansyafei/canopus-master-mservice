package repositories

import (
	"errors"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"
)

// GetPathFileByID : get path file by ID
func (Hnd Env) GetPathFileByID(id string) (models.Files, error) {
	var getData interface{}
	// Define files struct.
	files := models.Files{}
	// Check valid id.
	if bson.IsObjectIdHex(id) {
		queryGetData := bson.M{"_id": bson.ObjectIdHex(id)}
		err := Hnd.Mp.GetOne(files.TableName(), queryGetData, &getData)
		err = files.ToModel(getData, &files)
		if err != nil {
			return files, err
		}
		// Return files error.
		return files, err
	}
	// Return error.
	return files, errors.New("Invalid ID")
}
