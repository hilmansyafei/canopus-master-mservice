//main app
package main

import (
	"github.com/alterra/canopus-master-mservice/app/api"
	"github.com/alterra/canopus-master-mservice/database/mongo"
	"github.com/alterra/canopus-master-mservice/routers"
)

func main() {
	dbConn := mongo.MongodbConn()
	env := &api.Handler{DB: dbConn}
	e := routers.Gen(env)
	e.Logger.Fatal(e.Start(":1324"))
}
