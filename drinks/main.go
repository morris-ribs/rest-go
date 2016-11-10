package main

import (
	"log"
	"net/http"

	"github.com/rest-go/drinks/routes"
)

func main() {
	router := routes.NewRouter()

	// launch server
	log.Fatal(http.ListenAndServe(":9000", router))
}
