package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArcticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Text_1", Desc: "Text_description", Content: "Hi there 🚀🌟🌈"},
	}

	fmt.Println("8081/articles || Endpoint Hit 🌟")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
	fmt.Println("Base Hit 🌟")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/article", allArcticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	fmt.Println("🚀 Server started at\nlocalhost:8081")
	handleRequests()
}
