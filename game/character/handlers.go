package character

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soupstoregames/go-tick-yourself/logging"
)

type GetCharacterResponse struct {
	ID         uint64 `json:"id"`
	Balance    uint64 `json:"balance"`
	Reputation int16  `json:"reputation"`
}

func GetCharacter(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logging.WithError(err).Error("Failed to respond to GetCharacter")
			return
		}
		character, err := getCharacter(db, uint64(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			logging.WithError(err).Error("Failed to get character")
			return
		}

		response := GetCharacterResponse{
			ID:         character.id,
			Balance:    character.balance,
			Reputation: character.reputation,
		}

		if err := json.NewEncoder(w).Encode(&response); err != nil {
			logging.WithError(err).Error("Failed to respond to GetCharacter")
		}
	})
}

func GetMyCharacter(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement logic to get id of signed in user and assign to id
		id := 1
		character, err := getCharacter(db, uint64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logging.WithError(err).Error("Failed to get character")
			return
		}

		response := GetCharacterResponse{
			ID:         character.id,
			Balance:    character.balance,
			Reputation: character.reputation,
		}

		if err := json.NewEncoder(w).Encode(&response); err != nil {
			logging.WithError(err).Error("Failed to respond to GetCharacter")
		}
	})
}
