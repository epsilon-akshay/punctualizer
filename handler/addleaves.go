package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

func addLeave(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT COUNT(*) FROM leaves")
		checkErr(err)

		var count = 0
		for rows.Next() {
			count = count + 1
		}
		for i := 0; i <= count; i++ {
			id := r.FormValue(fmt.Sprintf("id%d", i))
			stmt, err := db.Prepare("update leaves set leaves=leaves+1 where id = ?")
			checkErr(err)

			res, err := stmt.Exec(id)
			checkErr(err)

			affect, err := res.RowsAffected()
			checkErr(err)
			fmt.Println(affect)
		}
	})
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
