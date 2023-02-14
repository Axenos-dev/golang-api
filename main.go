package main

import (
	"api-go/apiserver"
	"database/sql"
	"log"

	//"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	create_tables()

	config := apiserver.NewConfig()

	server := apiserver.New(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

// func that creates required to us tables
func create_tables() {
	users_db, err := sql.Open("sqlite3", "./database/users.db")
	if err != nil {
		panic(err)
	}

	pictures_db, err := sql.Open("sqlite3", "./database/pictures.db")
	if err != nil {
		panic(err)
	}

	defer users_db.Close()
	defer pictures_db.Close()

	// Create a new table
	users_sqlStmt := `
        CREATE TABLE IF NOT EXISTS users(
            id TEXT PRIMARY KEY,
            username TEXT,
			pass_hash TEXT
        );
    `
	pictures_sqlStmt := `
        CREATE TABLE IF NOT EXISTS pictures (
            user_id TEXT,
            url TEXT,
			path TEXT
        );
    `

	// pass_hash, _ := HashPassword("123456")
	// _, err = users_db.Exec(`
	//   	INSERT INTO users (id, username, pass_hash) VALUES(?, ?, ?)
	// `, uuid.New().String(), "test", pass_hash)

	// if err != nil {
	// 	panic(err)
	// }

	_, err = users_db.Exec(users_sqlStmt)
	if err != nil {
		panic(err)
	}

	_, err = pictures_db.Exec(pictures_sqlStmt)
	if err != nil {
		panic(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
