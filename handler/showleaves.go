package handler

import (
	"database/sql"
	"net/http"

	"html/template"
)

type employeeList struct {
	Employees []employee
}
type employee struct {
	ID     string
	Name   string
	Leaves int
}

func showLeave(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./handler/showleaves.html"))
		rows, err := db.Query("SELECT * FROM leaves")
		checkErr(err)

		emp := employee{}
		var emps []employee

		for rows.Next() {
			err = rows.Scan(&emp.Name, &emp.ID, &emp.Leaves)
			emps = append(emps, emp)
		}

		data := employeeList{
			Employees: emps,
		}

		rows.Close()
		tmpl.Execute(w, data)
	})
}
