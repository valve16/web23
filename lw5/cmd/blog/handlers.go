package main

import (
	"html/template"
	"log"
	"net/http"
)

type indexPage struct {
	Title         string
	FeaturedPosts []featuredPostData
	RecentsPosts []recentsPostData
}

type featuredPostData struct {
	Title    string
	Subtitle string
	Author   string
	Date     string
	AuthorImg string
	PostImg  string
	Commit   bool
}
type recentsPostData struct {
	Title    string
	Subtitle string
	Author   string
	Date     string
	AuthorImg string
	PostImg  string
}

func index(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := indexPage{
		Title:         "Escape",
		FeaturedPosts: featuredPosts(),
		RecentsPosts: recentsPosts(),
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:    "The Road Ahead",
			Subtitle: "The road ahead might be paved - it might not be.",
			Author:   "Mat Vogels",
			AuthorImg:"static/img/mat.jpg",
			Date:     "September 25, 2015",
			PostImg:  "static/img/northlight2.png",
			Commit: false,
		},
		{
			Title:    "From Top Down",
			Subtitle: "Once a year, go someplace you've never been before",
			Author:   "William Wong",
			AuthorImg:"static/img/william.jpg",
			Date:     "September 25, 2015",
			PostImg:  "static/img/chinalight.png",
			Commit: true,
		},
	}
}
func recentsPosts() []recentsPostData {
	return []recentsPostData{
		{
			Title:    "Still Standing Tall",
			Subtitle: "Life begins at the end of your comfort zone.",
			Author:   "William Wong",
			AuthorImg:"static/img/william.jpg",
			Date:     "9/25/2015",
			PostImg:   "static/img/airbaloons.jpg",
		},
		{
			Title:    "Sunny Side Up",
			Subtitle: "No place is ever as bad as they tell you it’s going to be.",
			Author:   "Mat Vogels",
			AuthorImg:"static/img/william.jpg",
			Date:     "9/25/2015",
			PostImg:   "static/img/redbridge.png",
		},
		{
			Title:    "Water Falls",
			Subtitle: "We travel not to escape life, but for life not to escape us.",
			Author:   "Mat Vogels",
			AuthorImg:"static/img/william.jpg",
			Date:     "9/25/2015",
			PostImg:   "static/img/sunnylake.png",
		},
		{
			Title:    "Through the Mist",
			Subtitle: "Travel makes you see what a tiny place you occupy in the world.",
			Author:   "William Wong",
			AuthorImg:"static/img/william.jpg",
			Date:     "9/25/2015",
			PostImg:   "static/img/water.png",
		},
		{
			Title:    "Awaken Early",
			Subtitle: "Not all those who wander are lost.",
			Author:   "Mat Vogels",
			AuthorImg:"static/img/william.jpg",
			Date:     "9/25/2015",
			PostImg:   "static/img/fog.png",
		},
		{
			Title:    "Try it Always",
			Subtitle: "The world is a book, and those who do not travel read only one page.",
			Author:   "Mat Vogels",
			AuthorImg:"static/img/william.jpg",
			Date:     "9/25/2015",
			PostImg:   "static/img/waterfall.png",
		},
	}
}
