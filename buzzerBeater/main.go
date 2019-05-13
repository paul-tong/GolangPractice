package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Username string
}

type Post struct {
	User // extend User struct, will have all the fields of User struct
	Body string
}

type IndexViewModel struct {
	Title string
	User  // extend User struct
	Posts []Post
}

func main() {
	// router, like a signal-slot, create a mapping relationship
	// the route will trigger corresponding function
	http.HandleFunc("/", showHomePage)
	http.HandleFunc("/posts", showAllPosts)
	http.HandleFunc("/love", sayLove)
	http.HandleFunc("/bye", sayBye)

	// start listening to the port
	http.ListenAndServe(":8888", nil)
}

// combine given data and content page with layout, then render the page to client
func renderTemplates(w http.ResponseWriter, contentPagePath string, data interface{}) {
	/*var allFiles []string{
		"templates/index.html"
		"templates/header.html"
	}*/

	tpl, _ := template.ParseFiles("templates/layout.html", contentPagePath)
	tpl.ExecuteTemplate(w, "layout", data)
}

func showHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("showHomePage\n")
	user := User{Username: "Tong"}

	renderTemplates(w, "templates/index.html", user)
}

func showAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Print("showAllPosts\n")

	u1 := User{Username: "tong"}
	u2 := User{Username: "duo duo"}

	posts := []Post{
		Post{User: u1, Body: "I like you duo duo!"},
		Post{User: u2, Body: "I like you 3000, peng tong"},
	}

	v := IndexViewModel{Title: "Homepage", User: u1, Posts: posts}

	renderTemplates(w, "templates/allPosts.html", &v)
}

func sayLove(w http.ResponseWriter, r *http.Request) {
	fmt.Print("sayLove\n")
	renderTemplates(w, "templates/love.html", "Tong")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Print("sayBye\n")
	renderTemplates(w, "templates/bye.html", "Tong")
}
