package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"

	// import from local package
	"github.com/paul-tong/buzzer_beater/controller"
	"github.com/paul-tong/buzzer_beater/model"
)

func main() {
	// setup  router
	router := httprouter.New()
	controller.SetupRouter(router)

	// connect to database and setup db instance
	db := model.ConnectToDB()
	defer db.Close() // this will excute at end of main function
	model.SetDB(db)

	// start listening to the port
	http.ListenAndServe(":8888", router)
}
