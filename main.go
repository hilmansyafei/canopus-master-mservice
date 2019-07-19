//main app
package main

import (
	"fmt"

	"github.com/hilmansyafei/canopus-master-mservice/app/api"
	"github.com/hilmansyafei/canopus-master-mservice/config"
	"github.com/hilmansyafei/canopus-master-mservice/database/repositories"
	"github.com/hilmansyafei/canopus-master-mservice/routers"
	"github.com/hilmansyafei/go-package/database/mongo"
	"github.com/hilmansyafei/go-package/modules"
	"go.uber.org/zap"
)

func main() {
	logInit, _ := modules.NewLogger(config.App.AppConfig.String("log_path"))
	api.ZapGlobal, _ = zap.NewProduction()
	configMongo := mongo.Configuration{
		Host:     config.App.AppConfig.String("host"),
		Port:     config.App.AppConfig.String("dbport"),
		Database: config.App.AppConfig.String("database"),
		User:     config.App.AppConfig.String("username"),
		Password: config.App.AppConfig.String("password"),
	}
	dbConn, err := mongo.New(configMongo)
	if err != nil {
		fmt.Errorf("Cannot connect to Mongo")
	}
	env := measure(logInit, dbConn)
	e := routers.Gen(env)
	e.Logger.Fatal(e.Start(":" + config.App.AppConfig.String("port")))
}

func measure(lP modules.LogProvider, mP mongo.MongoProvider) *api.Handler {
	api.LogGlobal = lP
	repoProvider := repositories.Env{
		Mp: mP,
	}
	return &api.Handler{
		MongoProvider: mP,
		Repositories:  repoProvider,
	}
}
