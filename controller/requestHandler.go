package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/paul-tong/buzzer_beater/model"
)

const layoutPagePath string = "view/templates/index.html"
const contentFolderPath string = "view/templates/contents/"

// test user and alert records
var (
	User1  = model.User{ID: 1, Email: "tong.gmail.com", PasswordHash: "tong"}
	User2  = model.User{ID: 2, Email: "duo.gmail.com", PasswordHash: "duoduo"}
	Alert1 = model.Alert{ID: 1, UserId: 1, EventId: "1", Section: "1", PriceLimit: 11.1}
	Alert2 = model.Alert{ID: 2, UserId: 2, EventId: "2", Section: "2", PriceLimit: 22.2}
)

// combine given data and content page with layout, then render the page to client
func renderTemplates(w http.ResponseWriter, contentPageName string, data interface{}) {

	contentPagePath := contentFolderPath + contentPageName
	tpl, _ := template.ParseFiles(layoutPagePath, contentPagePath)
	tpl.ExecuteTemplate(w, "index", data)
}

func showHomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print("showHomePage\n")
	user := model.User{Email: "Tong"}

	renderTemplates(w, "welcome.html", user)
}

func showAllPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print("showAllPosts\n")

	// preapre data
	u1 := model.User{Email: "tong"}
	u2 := model.User{Email: "duo duo"}

	posts := []model.Post{
		model.Post{User: u1, Body: "I like you duo duo!"},
		model.Post{User: u2, Body: "I like you 3000, peng tong"},
	}

	v := model.PostViewModel{Title: "Homepage", User: u1, Posts: posts}

	// send data and corresponding content page for renderring
	renderTemplates(w, "allPosts.html", &v)
}

func sayLove(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print("sayLove\n")
	renderTemplates(w, "love.html", "Tong")
}

func sayBye(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print("sayBye\n")
	renderTemplates(w, "bye.html", "Tong")
}

func showLogInForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print("showLogInPage\n")
	loginModel := model.LoginViewModel{}
	renderTemplates(w, "login.html", loginModel)
}

func handleLogInForm(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print("handleLogInForm\n")

	loginModel := model.LoginViewModel{}

	// show the login page if request type is Get
	/*if r.Method == http.MethodGet {
		renderTemplates(w, "login.html", loginModel)
		return
	}*/

	// verify the login data
	r.ParseForm()
	username := r.Form.Get("email")
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

func checkLogIn(username string, password string) bool {
	return username == "tong" && password == "1234"
}

// serach event information from ticketmaster api based on given keywords
func searchEventByKeywords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print("serach events by keywords\n")

	//loginModel := model.LoginViewModel{}

	// show the events page if request type is Get
	if r.Method == http.MethodGet {
		renderTemplates(w, "events.html", "Get mothod")
		return
	}

	// request data from ticketmaster api if request type is Post
	if r.Method == http.MethodPost {

		r.ParseForm()
		keyWords := r.Form.Get("searchKeyWords")

		// request results from ticktmaster api
		events, err := searchEventFromTicketmaster("keyword", string(keyWords), defaultEventCount)
		if err != nil {
			log.Println("Search Event from tickemaster api error: " + err.Error())
		}
		renderTemplates(w, "events.html", events)
	}
}

// serach event information from ticketmaster api based on given eventId
func searchEventByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Print("serach events by id\n")

	eventId := ps.ByName("eventId")

	// request results from ticktmaster api
	events, err := searchEventFromTicketmaster("eventId", string(eventId), defaultEventCount)
	if err != nil {
		log.Println("Search Event from tickemaster api error: " + err.Error())
	}

	// no event found
	if len(events) == 0 {
		log.Println("Sorry, cannot find this event.")
		renderTemplates(w, "eventDetail.html", nil) // send an nil
	}

	// return the only event to page
	renderTemplates(w, "eventDetail.html", events[0])
	/*if r.Method == http.MethodGet {
		renderTemplates(w, "events.html", "Get mothod")
		return
	}

	// request data from ticketmaster api if request type is Post
	if r.Method == http.MethodPost {

		r.ParseForm()
		keyWords := r.Form.Get("searchKeyWords")

		// request results from ticktmaster api
		events, err := searchEventFromTicketmaster("keyword", string(keyWords), defaultEventCount)
		if err != nil {
			log.Println("Search Event from tickemaster api error: " + err.Error())
		}
		renderTemplates(w, "events.html", events)
	}*/
}
