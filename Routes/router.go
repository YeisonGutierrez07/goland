package routes

import (
	persons "github.com/goland-project/Persons"
	"github.com/gorilla/mux"
)

// GetRouters Rutas para el proyecto
func GetRouters() *mux.Router {
	router := mux.NewRouter()

	// endPoints
	router.HandleFunc("/people", persons.GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", persons.GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", persons.CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", persons.DeletePersonEndPoint).Methods("DELETE")

	return router
}
