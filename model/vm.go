// view models: contains models that used for render view templates

package model

type PostViewModel struct {
	Title string
	User  // extend User struct, it will have all the fields of User struct
	Posts []Post
}

type LoginViewModel struct {
	Errs []string // contains error information of this login
}

func (v *LoginViewModel) AddError(err string) {
	v.Errs = append(v.Errs, err)
}
