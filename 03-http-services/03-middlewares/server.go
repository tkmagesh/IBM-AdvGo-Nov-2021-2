package main

import (
	"log"
	"net/http"
	"time"
)

type Handler func(http.ResponseWriter, *http.Request)

func foo(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("bar"))
}

/* func profile(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
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

func chain(handler func(w http.ResponseWriter, r *http.Request), middlewares ...func(func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		handler(w, r)
	}
} */

/*
func profile(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("request took %s", time.Since(start))
		}()
		handler(w, r)
	}
}

func logger(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		handler(w, r)
	}
}

func chain(handler Handler, middlewares ...func(Handler) Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		handler(w, r)
	}
}
*/

func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("request took %s", time.Since(start))
		}()
		handler(w, r)
	}
}

func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		handler(w, r)
	}
}

func chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		handler(w, r)
	}
}

func main() {

	/* http.HandleFunc("/foo", logger(profile(foo)))
	http.HandleFunc("/bar", logger(profile(bar))) */
	http.HandleFunc("/foo", chain(foo, logger, profile))
	http.HandleFunc("/bar", chain(bar, logger, profile))
	http.ListenAndServe(":8080", nil)
}
