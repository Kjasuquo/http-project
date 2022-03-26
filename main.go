package main

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world!")
}
