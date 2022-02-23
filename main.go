package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "Person type" (tipo um objeto)
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    PhoneNumber string `json:"phonenumber,omitempty"`
    
}

var people []Person

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// UpdatePerson faz alterações no contato
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

// função principal para executar a api
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "Geovane", Lastname: "Filho", PhoneNumber:"(82)99684-8507" })
    people = append(people, Person{ID: "2", Firstname: "Patrick", Lastname: "Brito", PhoneNumber: "(82)99999-9999"})
    router.HandleFunc("/contato", GetPeople).Methods("GET")
    router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/contato", CreatePerson).Methods("POST")
    router.HandleFunc("/contato/{id}", UpdatePerson).Methods("PUT")
    router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}