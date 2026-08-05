package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	name TEXT NOT NULL,
    	email TEXT NOT NULL UNIQUE,
    	hashed_password BLOB NOT NULL
	)
`

type User struct {
	ID             int  `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      string `json:"created_at"`	
}

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

	fmt.Println("Database connection established successfully");

	john, err := getUserByEmail(db, "john@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("User found: %+v\n", john)

}

// table create for database 
func createdTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func getUserByEmail(db *sql.DB, email string) (*User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}