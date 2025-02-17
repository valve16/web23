package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"
	"strings"

	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title           string
	SubTitle        string
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type featuredPostData struct {
	PostID      string `db:"post_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
	ImgModifier string `db:"image_mod"`
	PostImg     string `db:"image_url"`
}

type mostRecentPostData struct {
	PostID      string `db:"post_id"`
	PostImg     string `db:"image_url"`
	Title       string `db:"title"`
	Description string `db:"description"`
	AuthorImg   string `db:"author_url"`
	Author      string `db:"author"`
	PublishDate string `db:"publish_date"`
}

type postContent struct {
	Title    string `db:"title"`
	SubTitle string `db:"description"`
	ImgPost  string `db:"image_url"`
	Content  string `db:"content"`
}

type createPostRequest struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	AuthorName      string `json:"author"`
	AuthorPhoto     string `json:"avatar"`
	AuthorPhotoName string `json:"avatar_name"`
	Date            string `json:"date"`
	Image           string `json:"hero"`
	ImageName       string `json:"hero_name"`
	Content         string `json:"content"`
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		featuredPostsData, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		mostRecentPostsData, err := mostRecentPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		ts, err := template.ParseFiles("pages/index.html")
		if err != nil {
			http.Error(w, "Internal error", 500)
			log.Println(err)
			return
		}

		data := indexPage{
			Title:           "Let's do it together.",
			SubTitle:        "We travel the world in search of stories. Come along for the ride.",
			FeaturedPosts:   featuredPostsData,
			MostRecentPosts: mostRecentPostsData,
		}
		err = ts.Execute(w, data)

		if err != nil {
			http.Error(w, "Internal Server error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		postIDStr := mux.Vars(r)["postID"]

		postID, err := strconv.Atoi(postIDStr)
		if err != nil {
			http.Error(w, "Internal post id", http.StatusForbidden)
			log.Println(err.Error())
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Post not found", 404)
				log.Println(err.Error())
				return
			}

			http.Error(w, "Internal server error", 500)
			log.Println(err.Error())
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server error", 500)
			log.Println(err.Error())
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
		    post_id,
			  title,
			  description,
			  image_url,
			  author,
			  author_url,
			  publish_date,
			  image_mod
		FROM
			post
		WHERE featured = 1
	`

	var posts []featuredPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func mostRecentPosts(db *sqlx.DB) ([]mostRecentPostData, error) {
	const query = `
		SELECT
		    post_id,
		    image_url,
		    title,
		    description,
		    author,
		    author_url,
		    publish_date
		FROM
			post
		WHERE featured = 0
	`

	var posts []mostRecentPostData

	err := db.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func postByID(db *sqlx.DB, postID int) (postContent, error) {
	const query = `
	SELECT
		title,
		description,
		image_url,
		content
	FROM
		post
	WHERE
		  post_id = ?
	`

	var post postContent

	err := db.Get(&post, query, postID)
	if err != nil {
		return postContent{}, err
	}

	return post, nil
}

func admin(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		ts, err := template.ParseFiles("pages/admin.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		reqData, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to process request", 500)
			log.Println(err.Error())
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req)
		if err != nil {
			http.Error(w, "JSON Unmarshal failed", 500)
			log.Println(err.Error())
			return
		}

		req.AuthorPhotoName = "static/image/" + req.AuthorPhotoName
		req.ImageName = "static/image/" + req.ImageName

		err = saveFromBase64(req.AuthorPhoto, req.AuthorPhotoName)
		if err != nil {
			http.Error(w, "Unable to save Image", 500)
			log.Println(err.Error())
			return
		}
		err = saveFromBase64(req.Image, req.ImageName)
		if err != nil {
			http.Error(w, "Unable to save Image", 500)
			log.Println(err.Error())
			return
		}

		req.AuthorPhotoName = "/" + req.AuthorPhotoName
		req.ImageName = "/" + req.ImageName

		req.Date = formatDate(req.Date)

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "Unable to save post", 500)
			log.Println(err.Error())
			return
		}

		return
	}
}

func saveFromBase64(img string, path string) error {

	image, err := base64.StdEncoding.DecodeString(img)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	fileImage, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer fileImage.Close()

	_, err = fileImage.Write(image)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println("Image Saved")

	return nil
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	const query = `
		INSERT INTO
			post
		(
			title,
			description,
			author,
			author_url,
			publish_date,
			image_url,
			content,
			featured,
			image_mod
		)
		VALUES
		(
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)
	`

	_, err := db.Exec(query, req.Title, req.Description, req.AuthorName, req.AuthorPhotoName, req.Date, req.ImageName, req.Content, 0, 0)
	return err
}

func formatDate(oldDate string) string {
	dateStr := strings.Split(oldDate, "-")
	newDateStr := dateStr[2] + "/" + dateStr[1] + "/" + dateStr[0]
	return newDateStr
}

func login(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		ts, err := template.ParseFiles("pages/login.html")
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
}
