package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/soupstoregames/go-tick-yourself/logging"
)

func OpenConnection(dbName string, config Config) (*pgx.Conn, error) {
	uri := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d", dbName, config.User, config.Password, config.Host, config.Port)
	if !config.SSL {
		uri += " sslmode=disable"
	}

	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}

	// conn.SetMaxOpenConns(10)
	// conn.SetMaxIdleConns(4)
	// conn.SetConnMaxLifetime(4 * time.Second)

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	logging.Info(fmt.Sprintf("Connected to PostgreSQL %s/%s", config.Host, dbName))

	return conn, nil
}
