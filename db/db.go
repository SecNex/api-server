package db

import (
	"database/sql"
	"fmt"
	"log"
)

func EmptyConnection() *Connection {
	return &Connection{}
}

func (db *DB) ConnectInit() (*Connection, error) {
	log.Printf("Connecting to database %s...\n", "postgres")
	cnx, err := db.ConnectDatabase("postgres")
	if err != nil {
		return nil, err
	}

	log.Printf("Creating database %s...\n", db.Database)
	err = cnx.CreateDatabase(db.Database)
	if err != nil {
		return nil, err
	}

	cnx.Connection.Close()

	log.Printf("Connecting to database %s...\n", db.Database)
	return db.ConnectDatabase(db.Database)
}

func (c *Connection) CreateDatabase(name string) error {
	log.Printf("Dropping and creating database %s...\n", name)
	_, err := c.Connection.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", name))
	if err != nil {
		return err
	}

	_, err = c.Connection.Exec(fmt.Sprintf("CREATE DATABASE %s", name))
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) Search(table string, column string, value string) (*sql.Rows, error) {
	log.Printf("Searching %s for %s...\n", table, value)
	return c.Connection.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", table, column), value)
}

func (c *Connection) Exists(table string, column string, value string) (bool, error) {
	log.Printf("Checking if %s exists in %s:%s...\n", value, table, column)
	rows, err := c.Search(table, column, value)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
