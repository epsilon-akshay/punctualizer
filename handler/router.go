package handler

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) http.Handler {
	router := mux.NewRouter()
	addLeaves := addLeave(db)
	router.HandleFunc("/add_leaves", addLeaves)

	showLeaves := showLeave(db)
	router.HandleFunc("/show_leaves", showLeaves)
	return router
}
