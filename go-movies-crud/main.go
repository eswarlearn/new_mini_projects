package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// struct is like a java/js object have key value pairs with their data types
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // * is a pointer to the Director struct
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie //slice of Movie structs

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json") // set the content type to application/json
	json.NewEncoder(w).Encode(movies)                  // encode the movies slice to json and write it to the response writer
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contentent-type", "application/json")
	params := mux.Vars(r) // get the route variables
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // remove the movie from the slice
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // encode the updated movies slice to json and write it to the response writer
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	params := mux.Vars(r) // get the route variables
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // encode the movie to json and write it to the response writer
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/Json")
	var movie Movie                            // create a new movie variable
	_ = json.NewDecoder(r.Body).Decode(&movie) // decode the request body to the movie variable
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) // encode the movie to json and write it to the response writer
}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	// set the content type to application/json
	w.Header().Set("Content-Type", "")

	params := mux.Vars(r) // get the route variables
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // remove the movie from the slice
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // encode the updated movie to json and write it to the response writer
			return
		}
	}
}
func main() {
	r := mux.NewRouter() // create a new router using muc library

	// seed some data
	movies = append(movies, Movie{ID: "1", Isbn: "438743", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}}) //& to use the reference of the Director struct and * is used to dereference it

	movies = append(movies, Movie{ID: "2", Isbn: "438744", Title: "Movie Two", Director: &Director{FirstName: "Jane", LastName: "Doe"}})

	movies = append(movies, Movie{ID: "3", Isbn: "438745", Title: "Movie Three", Director: &Director{FirstName: "Jim", LastName: "Beam"}})

	// route handlers
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
