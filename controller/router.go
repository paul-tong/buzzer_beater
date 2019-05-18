package controller

import (
	"github.com/julienschmidt/httprouter"
)

/*setup routers, like a signal-slot, create a mapping relationship
  the route will trigger corresponding function
  use julienschmidt/httprouter to pass parameters and implement REST API
*/
func SetupRouter(router *httprouter.Router) {
	router.GET("/", showHomePage)
	router.GET("/posts", showAllPosts)
	router.GET("/love", sayLove)
	router.GET("/bye", sayBye)
	router.GET("/login", showLogInForm)
	router.POST("/login", handleLogInForm)
	router.POST("/events", searchEventByKeywords)
	router.GET("/event/:eventId", searchEventByID)
	// http.HandleFunc("/event/:eventId", searchEventById)
	/*http.HandleFunc("/", showHomePage)
	http.HandleFunc("/posts", showAllPosts)
	http.HandleFunc("/love", sayLove)
	http.HandleFunc("/bye", sayBye)
	http.HandleFunc("/login", handleLogIn)
	http.HandleFunc("/events", searchEventByKeywords)*/
	// http.HandleFunc("/event/:eventId", searchEventById)
}
