package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const (
	port         = ":3000"
	dbDriverName = "mysql"
)

func openDB() (*sql.DB, error) {
	return sql.Open(dbDriverName, "root:hvY-dpR-CXp-zG4@tcp(localhost:3306)/blog")
}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName)

	r := mux.NewRouter()
	r.HandleFunc("/home", homeHandler(dbx))
	r.HandleFunc("/post/{postId}/{postTitle}", postHandler(dbx))
	r.HandleFunc("/login", loginHandler(dbx))
	r.HandleFunc("/admin", adminHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	r.Handle("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./favicon.ico") }))

	log.Println("Start server http://localhost" + port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
}
