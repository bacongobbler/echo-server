package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for key, value := range r.Header {
			log.Printf("%s = %s", key, value)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
