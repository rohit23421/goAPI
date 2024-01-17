package router

import (
	"github.com/gorilla/mux"
	"github.com/rohit23421/mongoapi/controller"
)

//we return the refernce of the router
func Router() *mux.Router {
	router := mux.NewRouter()

	//route hadnler function for handling each route
	router.HandleFunc("/api/movies", controller.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")
	//themux router needs to be exported to be used in main file
	return router
}
