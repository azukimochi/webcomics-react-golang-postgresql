package routes

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/azukimochi/webcomics-react-golang-postgresql/server/controllers"
	"github.com/gorilla/mux"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// Router represents the routes for the API calls this app handles
func Router() *mux.Router {
	connStr := "user=azukimochi dbname=test_db host=localhost sslmode=disable"
	DBConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	c := controllers.CtrlDep{DB: DBConn}
	fmt.Println("Started postgreSQL server successfully!")

	router := mux.NewRouter()
	router.HandleFunc("/webcomic", c.AddComic).Methods("POST")
	router.HandleFunc("/webcomics", c.GetComics).Methods("GET")
	router.HandleFunc("/webcomic/{id}", c.UpdateComic).Methods("PUT")
	router.HandleFunc("/webcomic/{id}", c.DeleteComic).Methods("DELETE")

	return router
}
