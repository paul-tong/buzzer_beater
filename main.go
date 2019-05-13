package main

import (
	"net/http"

	// import from local package
	"github.com/paul-tong/buzzer_beater/controller"
)

func main() {
	// setup  router
	controller.SetupRouter()

	// start listening to the port
	http.ListenAndServe(":8888", nil)
}
