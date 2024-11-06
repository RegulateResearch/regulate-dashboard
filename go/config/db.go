package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", getDatabaseURL())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDatabaseURL() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		name := os.Getenv("DATABASE_NAME")
		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, name)
	}

	return url
}
