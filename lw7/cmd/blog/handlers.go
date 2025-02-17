package main

import (
	"html/template"
	"log"
	"net/http"

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
}
type recentsPostData struct {
	Title     string `db:"title"`
	Subtitle  string `db:"subtitle"`
	Author    string `db:"author"`
	Date      string `db:"publish_date"`
	AuthorImg string `db:"author_url"`
	PostImg   string `db:"image_url"`
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

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			author, 
			author_url, 
			publish_date, 
			image_url 
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

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
			image_url 
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции recents-posts

	var posts []recentsPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}
