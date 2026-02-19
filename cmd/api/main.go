package main

import (
	"log"
	"net/http"
)

func homeRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome! My name is Alex Guerrra.\nProject: E-Commerce Inventory & Order System\nWhy: To help MSME manage their inventory and orders efficiently.\n"))
}

func main() {

	mux := http.NewServeMux()

	// the route patterns
	mux.HandleFunc("/", homeRouteHandler)

	// start a local web server
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
