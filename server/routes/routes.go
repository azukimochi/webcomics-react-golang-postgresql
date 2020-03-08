package routes

import (
	"github.com/azukimochi/webcomics-react-golang-postgresql/server/controllers"
	"github.com/gorilla/mux"
)

// Router comment
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/webcomic", controllers.AddComic).Methods("POST")
	router.HandleFunc("/webcomics", controllers.GetComics).Methods("GET")
	router.HandleFunc("/webcomic/{id}", controllers.UpdateComic).Methods("PUT")
	router.HandleFunc("/webcomic/{id}", controllers.DeleteComic).Methods("DELETE")

	return router
}
