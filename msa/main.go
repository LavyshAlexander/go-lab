package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error with request", http.StatusBadRequest)
			return
		}

		log.Printf("Request body: %s\n", body)
		fmt.Fprintf(w, "Hello %s", body)
	})

	http.ListenAndServe(":9090", nil)
}
