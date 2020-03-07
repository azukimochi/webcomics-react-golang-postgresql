package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/azukimochi/webcomics-react-golang-postgresql/server/models"
	"github.com/gorilla/mux"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// DB is the database connection
var DB *sql.DB

func init() {
	var err error
	connStr := "user=azukimochi dbname=test_db host=localhost sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Started postgreSQL server successfully!")
}

// GetAllComics is a function that fetches all the webcomics from the DB
func GetAllComics(w http.ResponseWriter, r *http.Request) {
	var id int
	var title, author, status string
	rows, err := DB.Query("SELECT * FROM comics;")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Error occurred while trying to fetch the data from the database"))
		return
	}
	for rows.Next() {
		rows.Scan(&id, &title, &author, &status)
		fmt.Printf("Got: Id: %v, Title: %v, Author: %v, Status: %v\n", id, title, author, status)
	}
	w.WriteHeader(200)
	w.Write([]byte("Found something"))
}

////////////////// Code to Take Out Starts Here ///////////////////////////

var posts []models.Post = []models.Post{}

// GetPost is a function for getting a single post
func GetPost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	post := posts[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// // GetAllPosts is a function for getting all the posts
// func GetAllPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(posts)
// }

// // AddPost is a function for adding a single post
// func AddPost(w http.ResponseWriter, r *http.Request) {
// 	var newPost models.Post
// 	json.NewDecoder(r.Body).Decode(&newPost)
// 	posts = append(posts, newPost)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(posts)
// }

// UpdatePost is a function for ovewritting a single post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted into an integer"))
		return
	}

	if id > len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	var updatedPost models.Post
	json.NewDecoder(r.Body).Decode(&updatedPost)
	posts[id] = updatedPost

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPost)
}

// PatchPost is a function for updating a part of a single post
func PatchPost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted into an integer"))
		return
	}

	if id > len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// DeletePost is a function for deleting a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted into an integer"))
		return
	}

	if id > len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}
	// Delete the post from the slice. This is from the wiki for slice tricks, as slice doesn't have a delete method
	// https://github.com/golang/go/wiki/SliceTricks
	posts = append(posts[:id], posts[id+1:]...)

	w.WriteHeader(200)
}
