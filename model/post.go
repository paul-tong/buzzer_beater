package model

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Post struct {
	User // extend User struct, will have all the fields of User struct
	Body string
}

type PostTest struct {
	ID    int    `gorm:"primary_key"`
	Title string `gorm:"type:varchar(64)"`
	Body  string `gorm:"type:varchar(64)"`
}

func CreatePostTable() {
	db.DropTableIfExists(PostTest{})
	db.CreateTable(PostTest{})
}

// create one record
func AddPost(post PostTest) PostTest {
	db.Create(&post) // need to use &, why? => can modify the content of post
	return post
}

// get one record
func GetPostByID(id int) PostTest {
	var post PostTest
	db.First(&post, id)
	return post
}

func GetPostByTitle(title string) PostTest {
	var post PostTest
	db.Where(&PostTest{Title: "morning"}).First(&post)
	return post
}

// get all records
func GetAllPosts() []PostTest {
	var posts []PostTest
	db.Find(&posts)
	return posts
}

// delete one record
func DeletePostByID(id int) {
	post := PostTest{ID: id}
	db.Delete(&post)
}

// update, pass the item with new id and attributes
func UpdatePost(newPost PostTest) PostTest {
	db.Model(&newPost).Update(newPost)
	return newPost
}

// update part of item by id
func UpdatePostBodyByID(id int, body string) PostTest {
	post := PostTest{ID: id}
	db.Model(&post).Update("Body", body)
	return GetPostByID(id) // cannot just return post, it only contains the id and changed fields
}
