package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rohit23421/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API in Golang")
	//calling the router for rediretion of routes
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening on port 4000...")
}
