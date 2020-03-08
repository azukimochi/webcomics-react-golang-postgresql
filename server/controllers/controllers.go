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
	var data []models.Comic = []models.Comic{}

	rows, err := DB.Query("SELECT * FROM comics")
	if err != nil {
		errDesc := fmt.Sprintf("Failed querying the database when fetching for all the comics. %s", err.Error())
		ServeError(DBQueryFailed, errDesc, w, r)
		return
	}

	for rows.Next() {
		var id int
		var title, author, status string
		rows.Scan(&id, &title, &author, &status)
		comic := models.Comic{
			ID:     id,
			Title:  title,
			Author: author,
			Status: status,
		}
		data = append(data, comic)
		fmt.Printf("Got: Id: %v, Title: %v, Status: %v\n", id, title, status)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// AddComic is a function that adds a webcomic into the DB
func AddComic(w http.ResponseWriter, r *http.Request) {
	var comic models.Comic
	json.NewDecoder(r.Body).Decode(&comic)

	qr, err := DB.Exec("INSERT INTO comics (title, author, status) VALUES ($1, $2, $3)",
		comic.Title, comic.Author, comic.Status)
	if err != nil {
		errDesc := fmt.Sprintf("Failed adding the comic into the database. %s", err.Error())
		ServeError(DBChangeFailed, errDesc, w, r)
		return
	}

	fmt.Printf("Successfully inserted the record!. Results: %v\n", qr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comic)
}

// UpdateComic is a function that updates a webcomic in the DB
func UpdateComic(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	var updatedComic models.Comic

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errDesc := fmt.Sprintf("Failed converting data type of id. %s", err.Error())
		ServeError(InvalidURLParams, errDesc, w, r)
		return
	}

	json.NewDecoder(r.Body).Decode(&updatedComic)

	qr, err := DB.Exec("UPDATE comics SET title = $1 , author = $2, status = $3 WHERE id = $4",
		updatedComic.Title, updatedComic.Author, updatedComic.Status, id)
	if err != nil {
		errDesc := fmt.Sprintf("Failed updating the comic in the database. %s", err.Error())
		ServeError(DBChangeFailed, errDesc, w, r)
		return
	}

	rows, err := qr.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("Failed getting the number of rows affected. %s", err.Error())
		ServeError(GeneralDBError, errDesc, w, r)
		return
	}
	if rows < 1 {
		errDesc := "The specified record was not found."
		ServeError(DBChangeFailed, errDesc, w, r)
		return
	}
	fmt.Printf("Successfully updated the record! Results: %v", qr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedComic)
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
