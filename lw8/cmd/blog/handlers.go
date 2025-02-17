package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPageData struct {
	Title         string
	FeaturedPosts []featuredPostData
	RecentsPosts  []recentsPostData
}

type featuredPostData struct {
	Title     string `db:"title"`
	Subtitle  string `db:"subtitle"`
	Author    string `db:"author"`
	Date      string `db:"publish_date"`
	AuthorImg string `db:"author_url"`
	PostImg   string `db:"image_url"`
	PostId    string `db:"id"`
	Adventure int    `db:"adventure"`
}
type recentsPostData struct {
	Title     string `db:"title"`
	Subtitle  string `db:"subtitle"`
	Author    string `db:"author"`
	Date      string `db:"publish_date"`
	AuthorImg string `db:"author_url"`
	PostImg   string `db:"image_url"`
	PostId    string `db:"id"`
	Adventure int    `db:"adventure"`
}

type postData struct {
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	Contents string `db:"content"`
	PostImg  string `db:"image_url"`
	PostId   string `db:"id"`
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		featuredPosts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
			return                                      // Не забываем завершить выполнение ф-ии
		}
		recentsPosts, err := recentsPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		data := indexPageData{
			Title:         "Escape",
			FeaturedPosts: featuredPosts,
			RecentsPosts:  recentsPosts,
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
		log.Println("Request completed successfully")
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["id"] // Получаем postID в виде строки из параметров урла

		postID, err := strconv.Atoi(postIDStr) // Конвертируем строку postID в число
		if err != nil {
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author, 
			author_url, 
			publish_date, 
			image_url,
			id,
			adventure
		FROM
		    post 
		WHERE featured = 1`
	// Такое объединение строк делается только для таблицы order, т.к. это зарезерированное слово в SQL, наряду с SELECT, поэтому его нужно заключить в ``
	// Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []featuredPostData    // Заранее объявляем массив с результирующей информацией
	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func recentsPosts(db *sqlx.DB) ([]recentsPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author, 
			author_url, 
			publish_date, 
			image_url,
			id,
			adventure 
		FROM
		    post
		WHERE featured = 0` // Составляем SQL-запрос для получения записей для секции recents-posts

	var posts []recentsPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			image_url,
			content
		FROM
			post
		WHERE
			id = ?
	`
	// В SQL-запросе добавились параметры, как в шаблоне. ? означает параметр, который мы передаем в запрос ниже

	var post postData

	// Обязательно нужно передать в параметрах postID
	err := db.Get(&post, query, postID)
	if err != nil {
		return postData{}, err
	}

	return post, nil
}
