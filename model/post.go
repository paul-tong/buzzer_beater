package model

type Post struct {
	User // extend User struct, will have all the fields of User struct
	Body string
}
