package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"http/myBlog"
	"net/http"
	"time"
)

//Main is where my program runs
func main() {
	//chi is an external router package imported, and it serves as my router
	router := chi.NewRouter()

	myBlog.Register(router) // function of all to be routed
	fmt.Println("Server working ", time.Now())

	//Listening on my system port 8080 and routing it through the chi router
	e := http.ListenAndServe(":8080", router)
	myBlog.Error(e)

}
