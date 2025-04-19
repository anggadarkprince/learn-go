package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

func TestEmpty(t *testing.T) {
	// This test is intentionally left empty
	// It serves as a placeholder to ensure the test suite runs without errors
}

func TestOpenConnection(t *testing.T) {
	// Open a connection to the database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learn_golangdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}