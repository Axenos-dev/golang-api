package login

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func check_in_db(user User) UserInDB {
	users_db, err := sql.Open("sqlite3", "./database/users.db")
	if err != nil {
		fmt.Println(err)
		return UserInDB{}
	}

	defer users_db.Close()

	var find_results UserInDB

	row, _ := users_db.Query("SELECT * FROM users WHERE username=?", user.Username)

	for row.Next() {
		if err = row.Scan(&find_results.Id, &find_results.Username, &find_results.Password_hash); err != nil {
			fmt.Println(err.Error())
			return UserInDB{}
		}
	}
	return find_results
}
