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

// GetComics is a function that passes the request to either SearchForComics or GetAllComics
func GetComics(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	value, ok := queryParams["keywords"]
	if ok {
		SearchForComics(value[0], w, r)
		return
	}
	GetAllComics(w, r)
}

// GetAllComics ia  function that fetches all the webcomics from the DB
func GetAllComics(w http.ResponseWriter, r *http.Request) {
	var data []models.Comic

	rows, err := DB.Query("SELECT * FROM comics")
	if err != nil {
		errDesc := fmt.Sprintf("Failed querying the database when fetching for all the comics. %s", err.Error())
		ServeError(DBQueryFailed, errDesc, 403, w, r)
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
		fmt.Printf("Results of Global Search: Id: %v, Title: %v, Author: %v, Status: %v\n", id, title, author, status)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// SearchForComics is a function that gets all the webcomics by filtering the titles by keywords
func SearchForComics(v string, w http.ResponseWriter, r *http.Request) {
	var data []models.Comic
	iLikeOperatorVar := "%" + v + "%"

	rows, err := DB.Query("SELECT * FROM comics WHERE title ILIKE $1", iLikeOperatorVar)
	if err != nil {
		errDesc := fmt.Sprintf("Failed querying the database when fetching for comics based on the keywords. %s", err.Error())
		ServeError(DBQueryFailed, errDesc, 403, w, r)
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
		fmt.Printf("Filtered Search Results: Id: %v, Title: %v, Author: %v, Status: %v\n", id, title, author, status)
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
		ServeError(DBChangeFailed, errDesc, 403, w, r)
		return
	}

	fmt.Printf("Successfully inserted the record!. Results: %v\n", qr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(comic)
}

// UpdateComic is a function that updates a webcomic in the DB
func UpdateComic(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	var updatedComic models.Comic

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errDesc := fmt.Sprintf("Failed converting data type of id. %s", err.Error())
		ServeError(InvalidURLParams, errDesc, 422, w, r)
		return
	}

	json.NewDecoder(r.Body).Decode(&updatedComic)

	qr, err := DB.Exec("UPDATE comics SET title = $1 , author = $2, status = $3 WHERE id = $4",
		updatedComic.Title, updatedComic.Author, updatedComic.Status, id)
	if err != nil {
		errDesc := fmt.Sprintf("Failed updating the comic in the database. %s", err.Error())
		ServeError(DBChangeFailed, errDesc, 403, w, r)
		return
	}

	rows, err := qr.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("Failed getting the number of rows affected. %s", err.Error())
		ServeError(GeneralDBError, errDesc, 400, w, r)
		return
	}
	if rows < 1 {
		errDesc := "The specified record was not found."
		ServeError(DBChangeFailed, errDesc, 404, w, r)
		return
	}

	fmt.Printf("Successfully updated the record! Results: %v", qr)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedComic)
}

// DeleteComic is a function to delete a comic from the DB
func DeleteComic(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errDesc := fmt.Sprintf("Failed converting data type of id. %s", err.Error())
		ServeError(InvalidURLParams, errDesc, 422, w, r)
		return
	}

	qr, err := DB.Exec("DELETE FROM comics WHERE id = $1", id)
	if err != nil {
		errDesc := fmt.Sprintf("Failed to delete the comic from the database. %s", err.Error())
		ServeError(DBChangeFailed, errDesc, 403, w, r)
	}

	rows, err := qr.RowsAffected()
	if err != nil {
		errDesc := fmt.Sprintf("Failed getting the number of rows affected. %s", err.Error())
		ServeError(GeneralDBError, errDesc, 400, w, r)
		return
	}
	if rows < 1 {
		errDesc := "The specified record was not found."
		ServeError(DBChangeFailed, errDesc, 404, w, r)
		return
	}

	fmt.Printf("Successfully deleted the record with id %d! %v\n", id, qr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
