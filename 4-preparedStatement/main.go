package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	hashed_password BLOB NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

func main() {
	db, err := sql.Open("sqlite3", "users_database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully")

	// Create table
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Insert user
	_, err = createUserWithPreparedCtx(
		ctx,
		db,
		"Rakib",
		"rakib3@gmail.com",
		"password",
	)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserWithPreparedCtx(ctx context.Context, db *sql.DB, name, email, password string) (int64, error) {
	// Prepare SQL statement
	stmt, err := db.Prepare(`
		INSERT INTO users (name, email, hashed_password)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return 0, err
	}

	// Execute statement
	result, err := stmt.ExecContext(ctx, name, email, hashedPassword)
	if err != nil {
		return 0, err
	}

	// Return inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}
