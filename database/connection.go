package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/soupstoregames/go-tick-yourself/logging"
)

func OpenConnection(dbName string, config Config) (*sql.DB, error) {
	uri := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d", dbName, config.User, config.Password, config.Host, config.Port)
	if !config.SSL {
		uri += " sslmode=disable"
	}

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(4)
	db.SetConnMaxLifetime(4 * time.Second)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	logging.Info(fmt.Sprintf("Connected to PostgreSQL %s/%s", config.Host, dbName))

	return db, nil
}
