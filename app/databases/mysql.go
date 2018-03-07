package databases

import (
	"database/sql"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"

	// Driver file to usage as input for migration
	_ "github.com/mattes/migrate/source/file"
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

// RestoreDB ...
func RestoreDB(sqlPath string, dsn string) error {
	db, err := openConnection(dsn)
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+sqlPath,
		"mysql", driver)
	m.Steps(2)

	return err
}
