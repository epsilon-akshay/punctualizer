package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func addUser(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("Name")
		id := r.FormValue("Id")
		stmt, err := db.Prepare("insert into leaves(name, id ,leaves) values(?,?,?)")
		checkErr(err)

		exec, err := stmt.Exec(name, id, 0)
		checkErr(err)

		affect, err := exec.RowsAffected()
		checkErr(err)

		fmt.Println(affect)
	})
}
