package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      string `json:"created_at"`
}

func main() {
	dbName := "users_database.db"

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Println("Closing database connection")
		db.Close()
	}()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully")

	// Get one user
	user, err := getUserByEmail(db, "john@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	userJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Single User:")
	fmt.Println(string(userJSON))

	// Get all users
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	usersJSON, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nAll Users:")
	fmt.Println(string(usersJSON))
}

// Get all users
func getAllUsers(db *sql.DB) ([]User, error) {
	query := `
		SELECT id, name, email, hashed_password, created_at
		FROM users
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Get one user by email
func getUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `
		SELECT id, name, email, hashed_password, created_at
		FROM users
		WHERE email = ?
	`

	row := db.QueryRow(query, email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}