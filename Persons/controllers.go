package persons

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var people []Person

func init() {

	people = append(people, Person{IDLocal: "1", FirstName: "Yeison", LastName: "Gutierrez", Address: Address{City: "Soacha", State: "Olivos"}})
	people = append(people, Person{IDLocal: "2", FirstName: "Leidy", LastName: "Guerrero", Address: Address{City: "Bogotá", State: "Bosa"}})

}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, person := range people {
		if person.IDLocal == params["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(Person{})
	return
}

func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)

	person.IDLocal = params["id"]
	people = append(people, person)

	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i, person := range people {
		if person.IDLocal == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)

}
