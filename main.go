package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	Articles = []Article{
		Article{
			Title:   "Hello",
			desc:    "the first description",
			content: "na God help me reach here",
		},
	}
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the Homepage")
	fmt.Println("Endpoint hit: HomePage")

}
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

type Article struct {
	Title   string `json:"Title"`
	content string `json:"content"`
	desc    string `json:"desc"`
}

var Articles []Article
