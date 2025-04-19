package database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlStatement := "INSERT INTO customers(id, name) VALUES('angga', 'Angga')"

	_, err := db.ExecContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlStatement := "SELECT id, name FROM customers"

	rows, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlStatement := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customers"

	rows, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, email string
		var name sql.NullString // field can be null
		var balance float64
		var rating float32
		var birthDate sql.NullTime // field can be null
		var createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id)
		fmt.Println("Name:", name.String) // Use .String to get the value
		fmt.Println("Email:", email)	
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate.Time)
		} else {
			fmt.Println("Birth Date: No data available")	
		}
		fmt.Println("Married:", married)
		fmt.Println("Created At:", createdAt)
		fmt.Println("===================================")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Example of SQL injection
	//username := "admin'; DROP TABLE users; --"
	username := "not-found-user' OR 1=1; #"
	password := "wrong-password"
	sqlStatement := fmt.Sprintf("SELECT username FROM users WHERE username = '%s' AND password = '%s'", username, password)

	rows, err := db.QueryContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Login success:", username)
		}
	} else {
		fmt.Println("Login failed")
	}
}

func TestQueryWithParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// Example of SQL injection
	//username := "admin'; DROP TABLE users; --"
	username := "admin' OR 1=1; #"
	password := "wrong-password"
	sqlStatement := "SELECT username FROM users WHERE username = ? AND password = ?" // query with params

	rows, err := db.QueryContext(ctx, sqlStatement, username, password) // passing value
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Login success:", username)
		}
	} else {
		fmt.Println("Login failed")
	}
}

func TestExecSqlWithParam(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlStatement := "INSERT INTO users(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, sqlStatement, "angga", "password")
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new user")
}

func TestLastInsertID(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlStatement := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, sqlStatement, "angga@mail.com", "My comment")
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with ID:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlStatement := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	stmt, err := db.PrepareContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := range 10 {
		email := "email" + strconv.Itoa(i) + "@mail.com"
		comment := "Comment by " + email
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Success insert new comment with ID:", insertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	sqlStatement := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	stmt, err := tx.PrepareContext(ctx, sqlStatement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := range 5 {
		email := "email" + strconv.Itoa(i) + "@mail.com"
		comment := "Comment by " + email
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		fmt.Println("Success insert new comment with ID:", insertId)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}