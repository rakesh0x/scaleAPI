package main

import (
	"fmt"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the landing Page!")
	fmt.Println("Endpoint hit: Homepage")
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Welcome to Dashboard!")
	fmt.Println("Endpoint hit: Dashboard")
}

func handleRequest() {
	http.HandleFunc("/homepage", homepage)
	http.HandleFunc("/dashboard", dashboard)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequest()
}
