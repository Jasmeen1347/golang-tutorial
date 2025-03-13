package main

import (
	"fmt"
	"log"
	"mongoapis/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo DB API")

	log.Fatal(http.ListenAndServe(":4000", router.Router()))

	fmt.Println("Listing at port: 4000")
}
