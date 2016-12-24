package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

//Product implements a new type of struct for information about VR experiences
type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

var products = []Product{
    Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description : "Shoot your way to the top on 14 different hoverboards"},
    Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description : "Explore the depths of the sea in this one of a kind underwater experience"},
    Product{Id: 3, Name: "Dinosaur Park", Slug : "dinosaur-park", Description : "Go back 65 million years in the past and ride a T-Rex"},
    Product{Id: 4, Name: "Cars VR", Slug : "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
    Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description : "Pick up the bow and arrow and master the art of archery"},
    Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description : "Explore the seven wonders of the world in VR"},
}

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
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
})
