package main

import (
	"Go-Movie-CRUD/controllers"
	"Go-Movie-CRUD/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var movies []models.Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, models.Movie{
		ID:    "1",
		Isbn:  "1234",
		Title: "The call",
		Director: &models.Director{
			Firstname: "Williams",
			Lastname:  "Jackson",
		},
	})

	movies = append(movies, models.Movie{
		ID:    "2",
		Isbn:  "5678",
		Title: "Argument",
		Director: &models.Director{
			Firstname: "Servile",
			Lastname:  "Alex",
		},
	})

	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
	fmt.Printf("starting server on PORT 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
