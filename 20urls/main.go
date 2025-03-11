package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:5000/learn?course=reacrjs&price=300"

func main() {
	fmt.Println("Urls")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams)
	fmt.Println(qparams["course"])

	for _, value := range qparams {
		fmt.Println(value)
	}

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=jasmin",
	}

	fmt.Println(partsofurl.String())
}
