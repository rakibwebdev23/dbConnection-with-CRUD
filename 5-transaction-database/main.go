package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Database transaction

// -----------------------------
// 1. User create account
// 2. Create a wallete for the user
// 3. Want to top up the wallet for the user
// 4. You want to write a transaction log
// -----------------------------

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY References users(id) ON DELETE CASCADE,
	avatar TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"_"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

type Profile struct {
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully")

}
