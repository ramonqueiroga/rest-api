package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address, omitempty"`
}

type Address struct {
	city  string `json:city,omitempty`
	state string `json:state,omitempty`
}

var people []Person

func main() {
	people = append(people, Person{ID: "1", Firstname: "Felipe", Lastname: "Queiroga", Address: &Address{city: "São Bernardo do Campo", state: "São Paulo"}})
	people = append(people, Person{ID: "2", Firstname: "Ramon", Lastname: "Queiroga", Address: &Address{city: "Santo André", state: "São Paulo"}})
	people = append(people, Person{ID: "3", Firstname: "Fernanda", Lastname: "Toledo", Address: &Address{city: "Santo André", state: "São Paulo"}})
	createRouter()

}

func createRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
		}
	}
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = parametros["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, person := range people {
		if person.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
