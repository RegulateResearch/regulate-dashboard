package prep

import (
	"database/sql"
	"frascati/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", config.GetDatabaseURL())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
