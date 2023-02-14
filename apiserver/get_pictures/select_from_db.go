package get_pictures

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Select_From_DB(user_id string) []Picture {
	pictures_db, err := sql.Open("sqlite3", "./database/pictures.db")
	if err != nil {
		panic(err)
	}

	row, err := pictures_db.Query("SELECT * FROM pictures WHERE user_id = ?", user_id)
	if err != nil {
		return []Picture{}
	}

	var results []Picture

	// iterating over result
	for row.Next() {
		result := Picture{}
		err = row.Scan(&result.User_ID, &result.URL, &result.Path)
		if err != nil {
			return []Picture{}
		}

		results = append(results, result)
	}
	return results
}
