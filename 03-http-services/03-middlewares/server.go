package main

import (
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
}

func logger(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		handler(w, r)
	}
}

func main() {

	http.HandleFunc("/foo", logger(foo))
	http.HandleFunc("/bar", logger(bar))
	http.ListenAndServe(":8080", nil)
}
