package character

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/soupstoregames/go-tick-yourself/logging"
)

type GetCharacterResponse struct {
	ID 		   uint64 `json:"id"`
	Balance    uint64 `json:"balance"`
	Reputation int16  `json:"reputation"`
}

func GetCharacter(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		character, err := getCharacter(db, 1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logging.WithError(err).Error("Failed to get character")
			return
		}


		response := GetCharacterResponse{
			ID: character.id,
			Balance:    character.balance,
			Reputation: character.reputation,
		}

		if err := json.NewEncoder(w).Encode(&response); err != nil {
			logging.WithError(err).Error("Failed to respond to GetCharacter")
		}
	})
}
