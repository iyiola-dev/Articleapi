package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the Homepage")
	fmt.Println("Endpoint hit: HomePage")

}
func handleRequests() {
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

var Articles []Article

type Article struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Desc    string `json:"Desc"`
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint: Article has been hit")
	json.NewEncoder(w).Encode(Articles)
}
func main() {
	Articles = []Article{
		Article{
			Title:   "Hello",
			Desc:    "the first description",
			Content: "na God help me reach here",
		},
	}
	handleRequests()
}
