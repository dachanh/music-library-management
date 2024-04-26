package main

import (
	"fmt"
	"os"
	"time"

	"music-library-management/models"
	"music-library-management/sdk"

	"music-library-management/conf"

	"music-library-management/sdk/common"
	"music-library-management/sdk/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type infoData struct {
	Service     string    `json:"service"`
	Environment string    `json:"environment"`
	Version     string    `json:"version"`
	StartTime   time.Time `json:"startTime"`
}

var globalInfo *infoData

func info(req sdk.APIRequest, res sdk.APIResponder) error {
	return res.Respond(&common.APIResponse{
		Status: common.APIStatus.Ok,
		Data:   []interface{}{globalInfo},
	})
}

// onDBLogConnected ...
func onDBLogConnected(s *mongo.Database) error {
	return nil
}

// onDBConnected function that handle on connected to DB event
func onDBConnected(s *mongo.Database) error {

	models.InitMusicTrackModel(s)
	// model

	return nil
}

// onCacheConnected is func handle event connected to db cache
func onCacheConnected(s *mongo.Database) error {
	return nil
}

func main() {

	globalInfo = &infoData{
		Service:     "API",
		Version:     os.Getenv("version"),
		Environment: conf.Config.Env,
		StartTime:   time.Now(),
	}

	// setup new app
	var app = sdk.NewApp("service")
	configMap, err := app.GetConfigFromEnv()
	if err != nil {
		fmt.Println("Parse config error: " + err.Error())
		fmt.Println("Exiting app ...")
		return
	}

	// DB main
	app.SetupDBClient(db.Configuration{
		Address:  configMap["dbAddr"],
		Username: configMap["dbUser"],
		Password: configMap["dbPassword"],
		DBName:   conf.Config.MainDBName,
		AuthDB:   conf.Config.MainAuthDB,
	}, onDBConnected)

	// DB log
	logAddr := configMap["logAddr"]
	if logAddr == "" {
		logAddr = configMap["dbAddr"]
	}
	logUser := configMap["logUser"]
	if logUser == "" {
		logUser = configMap["dbUser"]
	}
	logPassword := configMap["logPassword"]
	if logPassword == "" {
		logPassword = configMap["dbPassword"]
	}
	app.SetupDBClient(db.Configuration{
		Address:  logAddr,
		Username: logUser,
		Password: logPassword,
		DBName:   conf.Config.LogDBName,
		AuthDB:   conf.Config.LogAuthDB,
	}, onDBLogConnected)

	// // setup second database
	// app.SetupDBClient(db.Configuration{
	// 	Address:  configMap["cacheAddr"],
	// 	Username: configMap["cacheUser"],
	// 	Password: configMap["cachePassword"],
	// 	DBName:   conf.Config.CacheDBName,
	// 	AuthDB:   "admin",
	// }, onCacheConnected)

	// setup API Server
	protocol := os.Getenv("protocol")

	var server, _ = app.SetupAPIServer(protocol)

	server.Expose(8080)

	// launch app
	if err := app.Launch(); err != nil {
		panic(err)
	}
}
