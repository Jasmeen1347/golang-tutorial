package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	greeter()

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey go mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang</h1>"))
}

// commands
// go get -u github.com/gorilla/mux
// go build .
//  go run main.go
// go mod tidy
// go mod verify
// go list
// go list all
// go list -m all
// go mod why github.com/gorilla/mux
// go mod graph
//  go mod vendor
