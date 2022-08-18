package main

import (
	"encoding/json"
	"net/http"
)

var consumers = map[string]*Consumer{}

type Consumer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/consumers", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(rw).Encode(consumers)
		case http.MethodPost:
			var consumer Consumer
			json.NewDecoder(r.Body).Decode(&consumer)

			consumers[consumer.Email] = &consumer

			json.NewEncoder(rw).Encode(consumer)
		}
	})
	http.ListenAndServe(":9000", nil)
}
