package mongo

import (
	"log"

	"github.com/hilmansyafei/canopus-master-mservice/config"
	"gopkg.in/mgo.v2"
)

// MongodbConn : COnnection to mongodb
func MongodbConn() *mgo.Database {
	// Connect to database
	session, err := mgo.Dial(config.App.AppConfig.String("host") + "/" + config.App.AppConfig.String("database"))

	// Check for error
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db := session.DB(config.App.AppConfig.String("database"))

	return db
}
