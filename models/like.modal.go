package models

type Like struct {
	ID   int
	Post *Post
	User *User
}
