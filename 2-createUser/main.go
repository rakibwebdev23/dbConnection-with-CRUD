package main

import (
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
	)
`

var db *sql.DB

// main function to establish database connection and create table
func main() {
	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Closing database connection")

		if err := db.Close(); err != nil {
			log.Print("Error closing database connection: ", err)
		}
	}()

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully")

	// createdTable(db)

	lastID, err := createUser(db, "John", "john@gmail.com", "password")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User created with ID:", lastID)

}

// table create for database 
func createdTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

// api create here 
func createUser(db *sql.DB, name, email, hashedPassword string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)`

	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec(stmt, name, email, string(hp))
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}