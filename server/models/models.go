package models

// Post is a struct that represents a single post
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

// Comic is a struct that represents a single comic
type Comic struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"`
}

// User is a struct that represents a user in our app
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

// ErrResponse is a struct represents error messages sent back to the client
type ErrResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
