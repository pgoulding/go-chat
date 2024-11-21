package main

import (
	"fmt"
	"log"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Home")
		fmt.Fprintln(w, "Simple Server")
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Test")
		fmt.Fprintln(w, "Test Page")
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("API")
		fmt.Fprintln(w, "Logic handling goes here")
	})

}

func main() {

	fmt.Println("Chat App V0.01")
	setupRoutes()
	log.Println("Starting server on :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
