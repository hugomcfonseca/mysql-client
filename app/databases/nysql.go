package databases

import (
	"database/sql"
)

func openConnection(dsn string) (*sql.DB, error) {
	connection, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	//defer connection.Close()

	if err = connection.Ping(); err != nil {
		return nil, err
	}

	return connection, nil
}

func close(db *sql.DB) error {
	return db.Close()
}

// NewDB ...
func NewDB(dbName string, dsn string) error {
	db, err := openConnection(dsn)
	if err != nil {
		return err
	}

	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + " DEFAULT CHARACTER SET = 'utf8'"); err != nil {
		return err
	}

	err = close(db)

	return err
}

// RemoveDB ...
func RemoveDB(dbName string, dsn string) error {
	db, err := openConnection(dsn)
	if err != nil {
		return err
	}

	if _, err := db.Exec("DROP DATABASE IF EXISTS " + dbName); err != nil {
		return err
	}

	err = close(db)

	return err
}
