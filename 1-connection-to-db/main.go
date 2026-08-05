package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
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

func main() {
	dbName := "data.db"

	//this is very very important to remove the database file before creating a new one, otherwise it will not create a new database file and will use the existing one
	_ = os.Remove(dbName)

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

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully")

	_,err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database schema table was created successfully")
	
}
