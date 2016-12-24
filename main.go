package main

import (
    "fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

//Product implements a new type of struct for information about VR experiences
type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}
//Create a catalog of VR products slice of Products
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
	r.HandleFunc("/status", StatusHandler).Methods("GET")
	//Producst route for retrieve a list of products user can leave feedback on
	r.HandleFunc("/products", ProductsHandler).Methods("GET")
	//Producst/{slug}/feedback route will capture user feedback on products
	r.HandleFunc("/products/{slug}/feedback", AddFeedbackHandler).Methods("POST")
	//Serve statics assets from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//run the application on port 3000
	http.ListenAndServe(":3000", r)
}

//NotImplemented handler handles whenever an API endpoint is hit it will return a not implemented message
func NotImplemented (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}

//StatusHandler will be inkoked when a user makes a GETY request to the /status endpoint
//it will simply return a string with the message relaying that the API is running
func StatusHandler (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("API is up and running"))
}

//ProductsHandler will be called when the user makes a GET request /products endpoint
//it will return a list of products for the user to review
func ProductsHandler (w http.ResponseWriter, r *http.Request) {
    //conver the slice of Products to json
    payload, _ := json.Marshal(products)

    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(payload))
}

//AddFeedbackHandler will add either positive or negative feedbackto the products
//For later save to database and send an OK status
func AddFeedbackHandler (w http.ResponseWriter, r *http.Request) {
    var product Product
    vars := mux.Vars(r)
    slug := vars["slug"]

    for _, p := range products {
        if p.Slug == slug {
            product = p
        }
    }

    w.Header().Set("Content-Type", "application/json")

    if product.Slug != "" {
        payload, _ := json.Marshal(product)
        w.Write([]byte(payload))
    } else {
        w.Write([]byte("Product Not Found"))
    }
}