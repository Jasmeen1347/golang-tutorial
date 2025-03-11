package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://example.com/"

func main() {
	fmt.Println("Web Requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// fmt.Println(response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}
