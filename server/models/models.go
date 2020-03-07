package models

// Post is a struct that represents a single post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

// User is a struct that represents a user in our app
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}
