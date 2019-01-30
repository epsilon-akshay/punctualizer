package main

import (
	"database/sql"
	"go-to-office-v2/handler"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./icecreammeter.db")
	if err != nil {
		log.Fatal(err)
	}
	router := handler.NewRouter(db)
	log.Fatal(http.ListenAndServe(":7000", router))
}
