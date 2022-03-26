package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hellohandler).Methods("GET")
	r.HandleFunc("/goodbye", goodbyehandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}

func goodbyehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Goodbye world!")
}
