package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func showLeave(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM leaves")
		checkErr(err)
		var id string
		var name string
		var leaves int

		for rows.Next() {
			err = rows.Scan(&name, &id, &leaves)
			checkErr(err)
			fmt.Println(id)
			fmt.Println(name)
			fmt.Println(leaves)
		}

		rows.Close()

	})
}
