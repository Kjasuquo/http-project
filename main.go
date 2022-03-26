package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Data struct {
	Status  bool
	Title   string
	Content string
}

var DataStructure []Data

//so that out templat will be accessible to our router, we create a global object
//var templates *template.Template

func main() {
	//templates = template.Must(template.ParseGlob("templat/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/index.html", indexhandler).Methods("GET")
	r.HandleFunc("/addpost.html", addContenthandler).Methods("GET")
	r.HandleFunc("/addpost.html", postContenthandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func indexhandler(w http.ResponseWriter, r *http.Request) {
	//p := []Data{{Status: true, Title: "The Best", Content: "Joseph is the best by God's grace"}, {Status: true, Title: "The Other", Content: "No other than God"}}
	t, _ := template.ParseFiles("templat/index.html")
	t.Execute(w, DataStructure)
	//templates.ExecuteTemplate(w, "index.html", nil)
}

func addContenthandler(w http.ResponseWriter, r *http.Request) {
	//p := []Data{{Status: true, Title: "The Best", Content: "Joseph is the best by God's grace"}, {Status: true, Title: "The Other", Content: "No other than God"}}
	t, _ := template.ParseFiles("templat/addpost.html")
	//t.Execute(w, DataStructure)
	t.ExecuteTemplate(w, "addpost.html", nil)
}

func postContenthandler(w http.ResponseWriter, r *http.Request) {
	f := Data{}
	r.ParseForm()
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	f.Title = title
	f.Content = content
	f.Status = true
	DataStructure = append(DataStructure, f)
	http.Redirect(w, r, "/index.html", 302)
}
