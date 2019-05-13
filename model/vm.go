// view models: contains models that used for render view templates

package model

type PostViewModel struct {
	Title string
	User  // extend User struct, it will have all the fields of User struct
	Posts []Post
}
