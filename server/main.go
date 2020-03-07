package main

import (
	"fmt"
	"net/http"

	"github.com/azukimochi/webcomics-react-golang-postgresql/server/controllers"
	"github.com/azukimochi/webcomics-react-golang-postgresql/server/routes"
)

func main() {
	// connStr := "user=azukimochi dbname=test_db host=localhost sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Started postgreSQL server successfully!")
	// defer db.Close()

	// var id int
	// var username string
	// name := "juny"
	// qr, err := db.Exec("DELETE FROM users where name = $1", name)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("Inserted successfully.  The results: %v \n", qr)

	// rows, err := db.Query("SELECT id, name, username FROM users")
	// for rows.Next() {
	// 	rows.Scan(&id, &name, &username)
	// 	fmt.Printf("Got: Id: %v, Name: %v, Username: %v\n", id, name, username)
	// }

	r := routes.Router()
	fmt.Println("Starting listening on port 5000!")
	defer controllers.DB.Close()
	http.ListenAndServe(":5000", r)
}
