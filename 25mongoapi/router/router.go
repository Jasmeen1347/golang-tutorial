package router

import (
	"mongoapis/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/movies", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
