package main

type homePage struct {
	Title           string
	FeaturedPosts   []featuredPostData
	MostRecentPosts []mostRecentPostData
}

type postPage struct {
	PostId     string `db:"id"`
	Img        string `db:"img"`
	ImgAlt     string `db:"img_alt"`
	Title      string `db:"title"`
	Subtitle   string `db:"subtitle"`
	Content    string `db:"content"`
	Paragraphs []string
}

type loginPage struct {
	Title string
}

type featuredPostData struct {
	PostId      string `db:"id"`
	Category    string `db:"category"`
	Title       string `db:"title"`
	URLTitle    string
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"img_modifier"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_img"`
	PublishDate string `db:"publish_date"`
}

type mostRecentPostData struct {
	PostId      string `db:"id"`
	Img         string `db:"img"`
	ImgAlt      string `db:"img_alt"`
	Title       string `db:"title"`
	URLTitle    string
	Subtitle    string `db:"subtitle"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_img"`
	PublishDate string `db:"publish_date"`
}
