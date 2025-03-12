package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("WebRequests")
	// perfomGetRequest()
	// performPostJsonRequest()
	performPostFormRequest()
}

func perfomGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code", response.StatusCode)

	// way one
	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))

	// way two
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func performPostJsonRequest() {
	const myurl = "http://localhost:3000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Angular",
			"price": 200
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:3000/postform"

	// formdata

	data := url.Values{}
	data.Add("Name", "Jasmeen")
	data.Add("Age", "27")

	resposne, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(resposne.Body)

	fmt.Println(string(content))
}
