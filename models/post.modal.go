package models

type Post struct {
	ID      int
	URL     string
	Caption string
	User    *User
	Likes   int
}
