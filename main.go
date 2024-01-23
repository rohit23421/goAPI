package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//if the url path doesnt mathers then return a error
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	//if the route is any other thing other than GET request
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// we try to parse the form if any error than return it
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParsseForm() error: %v", err)
	}
	fmt.Fprintf(w, "POST request successful")
	//taking the values from the form to these variables
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Address -%s\n", address)
}

func main() {
	fmt.Println("Web server from Golang")

	//asking the golanf httpdir method to check the gowebserver directory
	fileServer := http.FileServer(http.Dir("./static"))

	//handling the root route that is the "/" route by sending the index.html
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler) // for routing/handling the form
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server running on port 8080")
	//listening and serving on port 8080 and handling the error if any
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
