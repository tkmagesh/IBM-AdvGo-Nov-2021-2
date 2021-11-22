package main

import (
	"log"
	"net/http"
	"time"
)

func foo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("bar"))
}

func profile(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("request took %s", time.Since(start))
		}()
		handler(w, r)
	}
}

func logger(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		handler(w, r)
	}
}

func main() {

	http.HandleFunc("/foo", logger(profile(foo)))
	http.HandleFunc("/bar", logger(profile(bar)))
	http.ListenAndServe(":8080", nil)
}
