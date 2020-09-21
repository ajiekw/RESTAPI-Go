package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range designs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Design{})
}

func createDesign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var design Design
	_ = json.NewDecoder(r.Body).Decode(&design)
	design.ID = strconv.Itoa(rand.Intn(10000000))
	designs = append(designs, design)
	json.NewEncoder(w).Encode(design)
}

func updateDesign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range designs {
		if item.ID == params["id"] {
			designs = append(designs[:index], designs[index+1:]...)
			var design Design
			_ = json.NewDecoder(r.Body).Decode(&design)
			design.ID = params["id"]
			designs = append(designs, design)
			json.NewEncoder(w).Encode(design)
			return
		}
	}
	json.NewEncoder(w).Encode(designs)
}

func deleteDesign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range designs {
		if item.ID == params["id"] {
			designs = append(designs[:index], designs[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(designs)
}

func main() {
	r := mux.NewRouter()

	designs = append(designs, Design{ID: "1", Name: "The Fool", Category: "Painting", Creator: &Creator{Firstname: "John", Lastname: "Doe"}})
	designs = append(designs, Design{ID: "2", Name: "Dog Playing Poker", Category: "Painting", Creator: &Creator{Firstname: "Will", Lastname: "Smith"}})

	r.HandleFunc("/api/designs", getDesigns).Methods("GET")
	r.HandleFunc("/api/design/{id}", getDesign).Methods("GET")
	r.HandleFunc("/api/designs", createDesign).Methods("POST")
	r.HandleFunc("/api/design/{id}", updateDesign).Methods("PUT")
	r.HandleFunc("/api/design/{id}", deleteDesign).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
