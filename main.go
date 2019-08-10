package main

import (
	"fmt"
	"net/http"

	"os"

	routes "github.com/goland-project/Routes"
)

func main() {
	fmt.Println("EMPEZO LA vaina")

	routes := routes.GetRouters()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	http.ListenAndServe(":"+port, routes)
}
