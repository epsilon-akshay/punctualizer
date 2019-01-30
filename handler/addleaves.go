package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func addLeave(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		stmt, err := db.Prepare("update leaves set leaves=leaves+1 where id = ?")
		checkErr(err)

		res, err := stmt.Exec(id)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
