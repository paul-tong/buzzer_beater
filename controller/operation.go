package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/paul-tong/buzzer_beater/model"
)

const layoutPagePath string = "view/templates/index.html"
const contentFolderPath string = "view/templates/contents/"

// combine given data and content page with layout, then render the page to client
func renderTemplates(w http.ResponseWriter, contentPageName string, data interface{}) {

	contentPagePath := contentFolderPath + contentPageName
	tpl, _ := template.ParseFiles(layoutPagePath, contentPagePath)
	tpl.ExecuteTemplate(w, "index", data)
}

func showHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Print("showHomePage\n")
	user := model.User{Username: "Tong"}

	renderTemplates(w, "welcome.html", user)
}

func showAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Print("showAllPosts\n")

	// preapre data
	u1 := model.User{Username: "tong"}
	u2 := model.User{Username: "duo duo"}

	posts := []model.Post{
		model.Post{User: u1, Body: "I like you duo duo!"},
		model.Post{User: u2, Body: "I like you 3000, peng tong"},
	}

	v := model.PostViewModel{Title: "Homepage", User: u1, Posts: posts}

	// send data and corresponding content page for renderring
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
