package controller

import "net/http"

// setup routers, like a signal-slot, create a mapping relationship
// the route will trigger corresponding function
func SetupRouter() {
	http.HandleFunc("/", showHomePage)
	http.HandleFunc("/posts", showAllPosts)
	http.HandleFunc("/love", sayLove)
	http.HandleFunc("/bye", sayBye)

}
