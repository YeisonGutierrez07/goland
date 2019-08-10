package main

import (
	"fmt"
	"net/http"

	"os"

	routes "github.com/goland-project/Routes"
	"github.com/goland-project/shared"
)

func main() {

	shared.Init()
	// shared.GetDb().AutoMigrate(persons.Person{})

	routes := routes.GetRouters()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	fmt.Println("Corriendo servicios en el puerto: ", port)

	http.ListenAndServe(":"+port, routes)
}
