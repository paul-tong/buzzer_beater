package model

type Alert struct {
	ID         int
	UserId     int
	EventId    string
	Section    string
	PriceLimit float32
}
