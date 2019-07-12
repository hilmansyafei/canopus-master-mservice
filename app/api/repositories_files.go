package api

import (
	"encoding/hex"
	"errors"

	"github.com/hilmansyafei/canopus-master-mservice/database/models"
	"gopkg.in/mgo.v2/bson"
)

// GetPathFile : Get path file by MID and type file
// data["type"] : private / public
func GetPathFile(data map[string]string, h *Handler) (models.Files, error) {
	// Get merchant data.
	merchants, err := GetMerchantByMID(data["mid"])
	// Define files struct.
	files := models.Files{}
	// Check error.
	if err != nil {
		return files, err
	}
	// Define idFile as string.
	var idFile string
	// Check type in data.
	switch data["type"] {
	// If want to get private key.
	case "private":
		idFile = hex.EncodeToString([]byte(merchants.PsaPrivKey))
	// If want to get public key.
	case "public":
		idFile = hex.EncodeToString([]byte(merchants.MerchantPubKey))
	// Return error.
	default:
		return files, errors.New("File type not found")
	}
	// Get path uri.
	Path, errPath := GetPathFileByID(idFile)
	return Path, errPath
}

// GetPathFileByID : get path file by ID
func GetPathFileByID(id string) (models.Files, error) {
	var getData interface{}
	// Define files struct.
	files := models.Files{}
	// Check valid id.
	if bson.IsObjectIdHex(id) {
		queryGetData := bson.M{"_id": bson.ObjectIdHex(id)}
		err := MongoProvider.GetOne(files.TableName(), queryGetData, &getData)
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
