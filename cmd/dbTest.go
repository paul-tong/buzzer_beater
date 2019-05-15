// test on database operations

package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	// import from local package

	"github.com/paul-tong/buzzer_beater/model"
)

func main() {
	// connect to database and setup db instance
	db := model.ConnectToDB()
	defer db.Close() // this will excute at end of main function
	model.SetDB(db)

	// create table and add data
	// model.CreatePostTable()
	/*post1 := model.PostTest{Title: "morning", Body: "Good morning!"}
	post2 := model.PostTest{Title: "noon", Body: "Good noon!"}
	post3 := model.PostTest{Title: "evening", Body: "Good evening!"}
	model.AddPost(post1)
	model.AddPost(post2)
	model.AddPost(post3)*/

	// add one record
	/*post = model.PostTest{Title: "tong", Body: "Hello, tong!"}
	post = model.AddPost(post) // get new post(id added)
	fmt.Println(post)*/

	// delete one record
	model.DeletePostByID(8)

	// get record
	posts := model.GetAllPosts()
	fmt.Println(posts)

	post := model.GetPostByID(5)
	fmt.Println(post)

	post = model.GetPostByTitle("morning")
	fmt.Println(post)

	// update
	post = model.PostTest{ID: 9, Title: "duo", Body: "I like you"}
	post = model.UpdatePost(post)
	fmt.Println(post)

	post = model.UpdatePostBodyByID(9, "I like you")
	fmt.Println(post)
}
