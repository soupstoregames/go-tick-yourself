package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soupstoregames/go-tick-yourself/api/handlers"
	"github.com/soupstoregames/go-tick-yourself/game/character"
)

func BuildRoutes(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	r.Handle("/", handlers.NotImplemented()).Methods(http.MethodGet)

	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	r.Handle("/character", character.GetMyCharacter(db)).Methods(http.MethodGet)
	r.Handle("/character/{id:[0-9]+}", character.GetCharacter(db)).Methods(http.MethodGet)

	return r
}
