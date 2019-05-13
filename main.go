package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
func renderTemplates(w http.ResponseWriter, contentPageName string, data interface{}) {
	/*var allFiles []string{
		"templates/index.html"
		"templates/header.html"
	}*/
	contentPagePath := "view/templates/contents/" + contentPageName
	tpl, _ := template.ParseFiles("view/templates/index.html", contentPagePath)
	tpl.ExecuteTemplate(w, "index", data)
}

func showHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("showHomePage\n")
	user := User{model.Username: "Tong"}

	renderTemplates(w, "welcome.html", user)
}

func showAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Print("showAllPosts\n")

	u1 := User{Username: "tong"}
	u2 := User{Username: "duo duo"}

	posts := []Post{
		Post{User: u1, Body: "I like you duo duo!"},
		Post{User: u2, Body: "I like you 3000, peng tong"},
	}

	v := PostViewModel{Title: "Homepage", User: u1, Posts: posts}

	renderTemplates(w, "allPosts.html", &v)
}

func sayLove(w http.ResponseWriter, r *http.Request) {
	fmt.Print("sayLove\n")
	renderTemplates(w, "love.html", "Tong")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Print("sayBye\n")
	renderTemplates(w, "bye.html", "Tong")
}
