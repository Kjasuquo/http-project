package myBlog

import (
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

//Error handling
func Error(e error) {
	if e != nil {
		return
	}
}

// Data Structure containing everything I need to manipulate my data/content
type Data struct {
	Id      string
	Title   string
	Content string
	Status  bool
	Edit    string
	Delete  string
}

//DataStructure is a database in which my data are stored both before and after manipulation
var DataStructure []Data

//For holding data to be edited
var d Data

//Register initializes all my commands and html pages and it is called in the main
func Register(router *chi.Mux) {
	router.Get("/addpost", getContenthandler)
	router.Get("/del/{Id}", deleteByIdhandler)
	router.Get("/", indexhandler)
	router.Post("/addpost", postContenthandler)
	router.Get("/update/{Id}", updateByIdhandler)
	router.Post("/update/{Id}", postupdateByIdhandler)
}

//To get index page
func indexhandler(w http.ResponseWriter, r *http.Request) {
	//This points to the html location
	t, e := template.ParseFiles("templat/index.html")
	Error(e)

	//This writes whatever is in the DataStructure database to the html file
	e = t.Execute(w, DataStructure)
	Error(e)

}

//To get content page
func getContenthandler(w http.ResponseWriter, r *http.Request) {
	//This points to the html location
	// /post?url=12345
	t, e := template.ParseFiles("templat/addpost.html")
	Error(e)

	//This displays whatever on the html page
	e = t.ExecuteTemplate(w, "addpost.html", nil)
	Error(e)
}

//To manipulate content page
func postContenthandler(w http.ResponseWriter, r *http.Request) {
	//creating an instance of a data struct
	f := Data{}

	//This gets/populates the content of the form
	e := r.ParseForm()
	Error(e)

	//Gets the id/name of the form components
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	//Filling the data struct instance
	f.Title = title
	f.Content = content
	f.Status = true
	f.Edit = "Edit"
	f.Delete = "Delete"
	f.Id = uuid.NewString() //new id is being populated for an item using google/uuid

	if f.Content != "" {
		//Attach the f to the Data base so that it can be populated on the index page
		DataStructure = append(DataStructure, f)
	}

	//redirect your page back to the index/home page when done (on a click)
	http.Redirect(w, r, "/", 302)
}

//To delete each post
func deleteByIdhandler(w http.ResponseWriter, r *http.Request) {
	//Whenever it is clicked, get the id of its element
	ID := chi.URLParam(r, "Id")

	//Range through the database comparing the gotten id and perform basic algorithm on it
	for i, _ := range DataStructure {
		if ID == DataStructure[i].Id {

			//Deletes the item with such index
			DataStructure = append(DataStructure[:i], DataStructure[i+1:]...)
		}
	}

	//redirect your page back to the index/home page when done (on a click)
	http.Redirect(w, r, "/", 302)
}

//To get the content on a page when Edit is clicked
func updateByIdhandler(w http.ResponseWriter, r *http.Request) {
	//Whenever it is clicked, get the id of its element
	ID := chi.URLParam(r, "Id")

	//Range through the database comparing the gotten id and perform basic algorithm on it
	for i, _ := range DataStructure {
		if ID == DataStructure[i].Id {

			//This stores the data gotten in d for easy manipulation/editting
			d = DataStructure[i]

			//This points to the html location
			t, e := template.ParseFiles("templat/editpost.html")
			Error(e)

			//Calls or writes the item inside that database in the html file/template where it is called
			e = t.Execute(w, DataStructure[i])
			Error(e)

		}
	}
}

//After getting the content to be edited in html page, after editing, it will create a new item of it and store in database
func postupdateByIdhandler(w http.ResponseWriter, r *http.Request) {
	//This gets/populates the content of the form
	e := r.ParseForm()
	Error(e)

	//Gets the id/name of the form components
	title := r.PostForm.Get("tit")
	content := r.PostForm.Get("con")

	//Filling the data struct instance
	d.Title = title
	d.Content = content

	if d.Content != "" {

		//This is to check the DataStructure for id matching that of data d.id and delete
		for i, _ := range DataStructure {
			if DataStructure[i].Id == d.Id {
				DataStructure = append(DataStructure[:i], DataStructure[i+1:]...)
			}
		}

		//Attach the f to the Data base so that it can be populated on the index page
		DataStructure = append(DataStructure, d)
	}

	//redirect your page back to the index/home page when done (on a click)
	http.Redirect(w, r, "/", 302)
}
