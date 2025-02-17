package main

import (
	"database/sql"
	"log"
	"net/http" // Импортируем для возможности подключения к MySQL

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

const (
	port         = ":3000"
	dbDriverName = "mysql"
)

func main() {
	db, err := openDB() // Открываем соединение к базе данных в самом начале
	if err != nil {
		log.Fatal(err)
	}
	dbx := sqlx.NewDb(db, dbDriverName) // Расширяем стандартный клиент к базе

	mux := mux.NewRouter()
	mux.HandleFunc("/home", index(dbx)) // Передаём клиент к базе данных в ф-ию обработчик запроса

	// Указывем postID поста в URL для перехода на конкретный пост
	mux.HandleFunc("/post/{id}", post(dbx))
	// Реализуем отдачу статики
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Start server " + port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() (*sql.DB, error) {
	// Здесь прописываем соединение к базе данных
	return sql.Open(dbDriverName, "root:knight@tcp(localhost:3306)/blogoleg?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}
