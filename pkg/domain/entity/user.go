package entity

type User struct {
	ID string `dynamo:"Id,hash"`
	Name   string `dynamo:"Name,range"`
	Text   string `dynamo:"Text"`
}
