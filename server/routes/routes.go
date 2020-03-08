package routes

import (
	"github.com/azukimochi/webcomics-react-golang-postgresql/server/controllers"
	"github.com/gorilla/mux"
)

// Router comment
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/webcomic", controllers.AddComic).Methods("POST")
	router.HandleFunc("/webcomics", controllers.GetAllComics).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.PatchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")

	return router
}
