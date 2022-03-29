package fortesting

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	a := Error("No Eroor")
	fmt.Println(a)
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

//Error handling
func Error(e string) string {
	if e == "" {
		fmt.Println("Error occurred")
	}
	return e
}

//To get index page
func indexhandler(w http.ResponseWriter, r string) string {
	//This points to the html location
	t, _ := template.ParseFiles("templat/index.html")
	//Error(e)

	//This writes whatever is in the DataStructure database to the html file
	_ = t.Execute(w, DataStructure)
	//Error(e)

	return r

}

////To get content page
//func getContenthandler(w http.ResponseWriter, r *http.Request) {
//	//This points to the html location
//	t, e := template.ParseFiles("templat/addpost.html")
//	Error(e)
//
//	//This displays whatever on the html page
//	e = t.ExecuteTemplate(w, "addpost.html", nil)
//	Error(e)
//}
//
////To manipulate content page
//func postContenthandler(w http.ResponseWriter, r *http.Request) {
//	//creating an instance of a data struct
//	f := Data{}
//
//	//This gets/populates the content of the form
//	e := r.ParseForm()
//	Error(e)
//
//	//Gets the id/name of the form components
//	title := r.PostForm.Get("title")
//	content := r.PostForm.Get("content")
//
//	//Filling the data struct instance
//	f.Title = title
//	f.Content = content
//	f.Status = true
//	f.Edit = "Edit"
//	f.Delete = "Delete"
//	f.Id = uuid.NewString() //new id is being populated for an item using google/uuid
//
//	//Attach the f to the Data base so that it can be populated on the index page
//	DataStructure = append(DataStructure, f)
//
//	//redirect your page back to the index/home page when done (on a click)
//	http.Redirect(w, r, "/", 302)
//}
//
////To delete each post
//func deleteByIdhandler(w http.ResponseWriter, r *http.Request) {
//	//Whenever it is clicked, get the id of its element
//	ID := chi.URLParam(r, "Id")
//
//	//Range through the database comparing the gotten id and perform basic algorithm on it
//	for i, _ := range DataStructure {
//		if ID == DataStructure[i].Id {
//
//			//Deletes the item with such index
//			DataStructure = append(DataStructure[:i], DataStructure[i+1:]...)
//		}
//	}
//
//	//redirect your page back to the index/home page when done (on a click)
//	http.Redirect(w, r, "/", 302)
//}
//
////To get the content on a page when Edit is clicked
//func updateByIdhandler(w http.ResponseWriter, r *http.Request) {
//	//Whenever it is clicked, get the id of its element
//	ID := chi.URLParam(r, "Id")
//
//	//Range through the database comparing the gotten id and perform basic algorithm on it
//	for i, _ := range DataStructure {
//		if ID == DataStructure[i].Id {
//
//			//This stores the data gotten in d for easy manipulation/editting
//			d = DataStructure[i]
//
//			//This points to the html location
//			t, e := template.ParseFiles("templat/editpost.html")
//			Error(e)
//
//			//Calls or writes the item inside that database in the html file/template where it is called
//			e = t.Execute(w, DataStructure[i])
//			Error(e)
//
//		}
//	}
//}
//
////After getting the content to be edited in html page, after editing, it will create a new item of it and store in database
//func postupdateByIdhandler(w http.ResponseWriter, r *http.Request) {
//	//This gets/populates the content of the form
//	e := r.ParseForm()
//	Error(e)
//
//	//This is to check the DataStructure for id matching that of data d.id and delete
//	for i, _ := range DataStructure {
//		if DataStructure[i].Id == d.Id {
//			DataStructure = append(DataStructure[:i], DataStructure[i+1:]...)
//		}
//	}
//
//	//Gets the id/name of the form components
//	title := r.PostForm.Get("tit")
//	content := r.PostForm.Get("con")
//
//	//Filling the data struct instance
//	d.Title = title
//	d.Content = content
//
//	//Attach the f to the Data base so that it can be populated on the index page
//	DataStructure = append(DataStructure, d)
//
//	//redirect your page back to the index/home page when done (on a click)
//	http.Redirect(w, r, "/", 302)
//}
