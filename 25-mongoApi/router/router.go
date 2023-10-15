package router

import (
	"github.com/gorilla/mux"
	controller "github.com/kmabhinai/go-mongoapi/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/deleteAll", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")

	return router
}
