package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soupstoregames/go-tick-yourself/api/handlers"
	"github.com/soupstoregames/go-tick-yourself/game/character"
)

func BuildRoutes() http.Handler {
	r := mux.NewRouter()

	r.Handle("/", handlers.NotImplemented()).Methods(http.MethodGet)

	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	r.Handle("/character", character.GetCharacter()).Methods(http.MethodGet)

	return r
}
