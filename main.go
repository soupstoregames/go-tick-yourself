package main

import (
	"github.com/soupstoregames/go-tick-yourself/api"
	"github.com/soupstoregames/go-tick-yourself/logging"
)

var (
	service = "go-tick-yourself"
	version = "dev"
)

func main() {
	logging.SetStandardFields(service, version)

	routes := api.BuildRoutes()
	apiServer := api.NewHTTPServer("0.0.0.0:8080", routes)

	logging.Info("Starting HTTP Server")
	if err := apiServer.ListenAndServe(); err != nil {
		logging.WithErr(err).Fatal("HTTP server failed")
	}
}
