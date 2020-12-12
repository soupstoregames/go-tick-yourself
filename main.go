package main

import (
	"github.com/codingconcepts/env"
	"github.com/soupstoregames/go-tick-yourself/api"
	"github.com/soupstoregames/go-tick-yourself/database"
	"github.com/soupstoregames/go-tick-yourself/logging"
)

var (
	service = "go-tick-yourself"
	version = "dev"
)

type Config struct {
	Database database.Config
}

func main() {
	logging.SetStandardFields(service, version)

	conf := Config{}
	if err := env.Set(&conf.Database); err != nil {
		logging.WithError(err).Fatal("Failed to read config from env vars")
	}

	_, err := database.OpenConnection("gotickyourself", conf.Database)
	if err != nil {
		logging.WithError(err).Error("Failed to connect to postgres")
	}

	routes := api.BuildRoutes()
	apiServer := api.NewHTTPServer("0.0.0.0:8080", routes)

	logging.Info("Starting HTTP Server")
	if err := apiServer.ListenAndServe(); err != nil {
		logging.WithError(err).Fatal("HTTP server failed")
	}
}
