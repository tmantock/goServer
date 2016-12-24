package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
    //Initialize mux router
    r := mux.NewRouter()

    //Serve static index page on root route
    r.Handle("/", http.FileServer(http.Dir("/views/")))
    //Serve statics assets from the /static/{file} route
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    //run the application on port 3000
    http.ListenAndServe(":3000", r)
}