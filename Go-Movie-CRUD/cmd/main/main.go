package main

import (
	"fmt"
	"log"
	"net/http"
	"Go-Movie-CRUD/controllers"
	"Go-Movie-CRUD/models"

	
	"github.com/gorilla/mux"
)

// type Director struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// type Movie struct {
// 	ID       string    `json:"id"`
// 	Isbn     string    `json:"isbn"`
// 	Title    string    `json:"title"`
// 	Director *Director `json:"director"` // embed Director struct
// }

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
