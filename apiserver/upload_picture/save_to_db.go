package upload_picture

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func insert_into_db(user_id string, url string, path string) {
	pictures_db, err := sql.Open("sqlite3", "./database/pictures.db")
	if err != nil {
		panic(err)
	}

	pictures_db.Exec(`
	INSERT INTO pictures (user_id, url, path) VALUES(?, ?, ?)
	`, user_id, url, path)
}
