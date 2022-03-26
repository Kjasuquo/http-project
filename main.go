package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

//so that out templat will be accessible to our router, we create a global object
var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templat/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexhandler).Methods("GET")
	// r.HandleFunc("/goodbye", goodbyehandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

//remove the goodbye handler
//func goodbyehandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Goodbye world!")
//}
