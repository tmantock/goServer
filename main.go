package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
    //Initialize mux router
    r := mux.NewRouter()

    //Serve static index page on root route
    r.Handle("/", http.FileServer(http.Dir("./views/")))
    //Status route will be called to make sure the API is up and running
    r.Handle("/status", NotImplemented).Methods("GET")
    //Producst route for retrieve a list of products user can leave feedback on
    r.Handle("/products", NotImplemented).Methods("GET")
    //Producst/{slug}/feedback route will capture user feedback on products
    r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")
    //Serve statics assets from the /static/{file} route
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    //run the application on port 3000
    http.ListenAndServe(":3000", r)
}

//NotImplemented handler handles whenever an API endpoint is hit it will return a not implemented message
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Not implemented"))
})