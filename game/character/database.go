package character

import "database/sql"

type character struct {
	id         uint64
	balance    uint64
	reputation int16
}

func getCharacter(db *sql.DB, id uint64) (character, error) {
	row := db.QueryRow("SELECT * FROM characters WHERE id = $1", id)

	if err := row.Err(); err != nil {
		return character{}, err
	}

	var c character
	if err := row.Scan(&c.id, &c.balance, &c.reputation); err != nil {
		return character{}, err
	}

	return c, nil
}
