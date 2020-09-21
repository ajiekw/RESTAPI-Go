package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Design struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Creator  *Creator `json:"Creator"`
}

type Creator struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var designs []Design

//Get
func getDesigns(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(designs)
}

func getDesign(w http.ResponseWriter, r *http.Request) {

}

func createDesign(w http.ResponseWriter, r *http.Request) {

}

func updateDesign(w http.ResponseWriter, r *http.Request) {

}

func deleteDesign(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	designs = append(designs, Design{ID: "1", Name: "The Fool", Category: "Painting", Creator: &Creator{Firstname: "John", Lastname: "Doe"}})
	designs = append(designs, Design{ID: "2", Name: "Dog Playing Poker", Category: "Painting", Creator: &Creator{Firstname: "Will", Lastname: "Smith"}})

	r.HandleFunc("/api/books", getDesigns).Methods("GET")
	r.HandleFunc("/api/book/{id}", getDesign).Methods("GET")
	r.HandleFunc("/api/books", createDesign).Methods("POST")
	r.HandleFunc("/api/book/{id}", updateDesign).Methods("PUT")
	r.HandleFunc("/api/book/{id}", deleteDesign).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
