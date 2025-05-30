package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"resfull-api/helper"
	"time"
)

func NewDB() *sql.DB {
	// Assuming you have a function to initialize your database connection
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/sandbox")
	helper.PanicIfError(err)
	
	// Set the maximum number of open connections to the database
	db.SetMaxOpenConns(25)
	
	// Set the maximum number of idle connections in the pool
	db.SetMaxIdleConns(5)
	
	// Set the maximum lifetime of a connection
	db.SetConnMaxLifetime(time.Minute * 10)

	// Set the maximum idle time for connections
	db.SetConnMaxIdleTime(time.Minute * 5)
	
	return db
}