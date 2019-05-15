package model

type Alert struct {
	User    // extend User struct, will have all the fields of User struct
	Section string
	MinRow  int
	MaxRow  int
	Price   float32
}
