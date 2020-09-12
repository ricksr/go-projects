package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
	fmt.Println("Base Hit 🌟")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go Orm ")

	handleRequests()
}
