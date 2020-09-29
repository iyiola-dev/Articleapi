package main

import (
	"encoding/json"
	"net/http"

	"frest/structs"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{}

			json.NewEncoder(w).encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)
}
