package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // embed Director struct
}

var movies []Movie

//func HandleFunc() {}

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// deleted a particular movie from the array of movies struct literals
func deleteMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "1234",
		Title: "The call",
		Director: &Director{
			Firstname: "Williams",
			Lastname:  "Jackson",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "5678",
		Title: "Argument",
		Director: &Director{
			Firstname: "Servile",
			Lastname:  "Alex",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("starting server on PORT 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
