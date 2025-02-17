package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func homeHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		featuredPosts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		for index, post := range featuredPosts {
			featuredPosts[index].URLTitle = strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))
		}

		mostRecentPosts, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		for index, post := range mostRecentPosts {
			mostRecentPosts[index].URLTitle = strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))
		}

		ts, err := template.ParseFiles("pages/home.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := homePage{
			Title:           "Escape",
			FeaturedPosts:   featuredPosts,
			MostRecentPosts: mostRecentPosts,
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func postHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := `
			SELECT
				id,
				title,
				subtitle,
				img,
				img_alt,
				content
			FROM
				post
			WHERE id = 
		`
		query += mux.Vars(r)["postId"]

		var content []postPage
		err := db.Select(&content, query)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
		if len(content) > 0 {
			ts, err := template.ParseFiles("pages/post.html")
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				log.Println(err.Error())
				return
			}

			content[0].Paragraphs = strings.Split(content[0].Content, "\n")

			err = ts.Execute(w, content[0])
			if err != nil {
				http.Error(w, "Internal Server Error", 500)
				log.Println(err.Error())
				return
			}
		} else {
			http.Redirect(w, r, "/home", http.StatusNotFound)
		}
	}
}

func loginHandler(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("pages/login.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		data := loginPage{
			Title: "Login",
		}

		err = ts.Execute(w, data)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles("pages/admin.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

	data := loginPage{
		Title: "Admin",
	}

	err = ts.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if stringInSlice(pathSegments[1], []string{"index", ""}) {
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
	}
	ts, err := template.ParseFiles("pages/not-found.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}
