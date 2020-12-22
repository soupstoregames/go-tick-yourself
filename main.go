package main

import (
	"github.com/codingconcepts/env"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"

	"github.com/soupstoregames/go-tick-yourself/api"
	"github.com/soupstoregames/go-tick-yourself/database"
	"github.com/soupstoregames/go-tick-yourself/database/migrations"
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

	db, err := database.OpenConnection("gotickyourself", conf.Database)
	if err != nil {
		logging.WithError(err).Fatal("Failed to connect to postgres")
	}

	if err := database.ValidateSchema(db, bindata.Resource(migrations.AssetNames(), migrations.Asset)); err != nil {
		logging.WithError(err).Error("Failed to validate schema")
	}

	routes := api.BuildRoutes(db)
	apiServer := api.NewHTTPServer("0.0.0.0:8080", routes)

	logging.Info("Starting HTTP Server")
	if err := apiServer.ListenAndServe(); err != nil {
		logging.WithError(err).Fatal("HTTP server failed")
	}
}
