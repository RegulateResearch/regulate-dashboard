package setup

import (
	"database/sql"
	"frascati/exception"
	"frascati/prep"
)

func setupDb() (*sql.DB, exception.Exception) {
	db, err := prep.ConnectDB()
	if err != nil {
		return nil, exception.NewBaseException(exception.CAUSE_INTERNAL, "app startup", "fail to init db", err)
	}

	return db, nil
}
