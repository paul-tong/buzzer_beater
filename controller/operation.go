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

func handleLogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Print("handleLogIn\n")

	// show the login page if request type is Get
	if r.Method == http.MethodGet {
		renderTemplates(w, "login.html", "")
		return
	}

	// verify the login data if request type is Post
	if r.Method == http.MethodPost {
		loginModel := model.LoginViewModel{}

		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// check the validation of input format (usually used for register)
		if len(username) < 3 {
			loginModel.AddError("username must be longer than 3")
		}

		if len(password) < 4 {
			loginModel.AddError("password must be longer than 6")
		}

		if !checkLogIn(username, password) {
			loginModel.AddError("username or password not correct")
		}

		// check whether username and possword are correct
		if checkLogIn(username, password) {
			fmt.Fprintf(w, "Login success! Username:%s Password:%s", username, password)
		} else {
			renderTemplates(w, "login.html", loginModel)
		}
	}
}

func checkLogIn(username string, password string) bool {
	return username == "tong" && password == "1234"
}
