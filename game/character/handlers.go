package character

import (
	"encoding/json"
	"net/http"

	"github.com/soupstoregames/go-tick-yourself/logging"
)

type GetCharacterResponse struct {
	Balance    uint64 `json:"balance"`
	Reputation int16  `json:"reputation"`
}

func GetCharacter() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := GetCharacterResponse{}

		if err := json.NewEncoder(w).Encode(&response); err != nil {
			logging.WithError(err).Error("Failed to respond to GetCharacter")
		}
	})
}
