package entity

type User struct {
	UserId string `dynamo:"UserId,hash"`
	Name   string `dynamo:"Name,range"`
	Text   string `dynamo:"Text"`
}
