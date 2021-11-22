package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!", r.URL.Query().Get("name"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("All the products will be served"))
			return
		}
		if r.Method == "POST" {
			w.Write([]byte("New product will be created"))
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	http.ListenAndServe(":8080", nil)
}
