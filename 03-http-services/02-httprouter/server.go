package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome!"))
}

func Greet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!", ps.ByName("name"))
}

func GreetPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("/Users/tkmagesh77/Documents/Training/IBM-AdvGo-Nov-2021-2/03-http-services/02-httprouter/static"))
	router.GET("/", Index)
	router.GET("/greet/:name", Greet)
	router.POST("/greet/:name", GreetPost)
	http.ListenAndServe(":8080", router)
}
