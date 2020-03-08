package models

// Comic is a struct that represents a single comic
type Comic struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Status string `json:"status,omitempty"`
}

// ErrResponse is a struct represents error messages sent back to the client
type ErrResponse struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}
