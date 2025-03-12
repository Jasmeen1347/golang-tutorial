package main

import (
	"encoding/json" // Import JSON package for encoding and decoding
	"fmt"           // Import fmt package for printing output
)

// Define a struct to represent a course
// Struct field tags specify how fields are encoded/decoded in JSON
// `json:"-"` means the field won't be included in JSON
// `json:",omitempty"` means the field will be omitted if it's empty or nil
type course struct {
	Name     string   `json:"coursename"` // JSON key will be "coursename"
	Price    int      // Default JSON key will be "Price"
	Platform string   `json:"website"`        // JSON key will be "website"
	Password string   `json:"-"`              // This field is ignored in JSON
	Tags     []string `json:"tags,omitempty"` // Omits field if nil or empty
}

func main() {
	fmt.Println("handling json data")
	// EncodeJson() // Uncomment to test JSON encoding
	DecodeJson() // Calls function to decode JSON
}

// Function to convert Go struct data into JSON
func EncodeJson() {
	// Creating a slice (list) of courses
	courses := []course{
		{"Reactjs-bootcamp", 299, "lco.in", "avcsfs123", []string{"web-dev", "frontend"}},
		{"MERN-bootcamp", 499, "lco.in", "avcsfs123", []string{"web-dev", "fullstack"}},
		{"J#-bootcamp", 599, "lco.in", "avcsfs123", nil}, // "Tags" is nil
	}

	// Convert the struct slice to JSON with indentation for readability
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	// Check for encoding errors
	if err != nil {
		panic(err) // Stop execution and print error if encoding fails
	}

	// Print the formatted JSON output
	fmt.Printf("%s \n", finalJson)
}

// Function to parse JSON into a Go struct
func DecodeJson() {
	// Simulating JSON data received from the web (as a byte slice)
	jsonDataFromWeb := []byte(`
		{
			"coursename": "Reactjs-bootcamp",
			"Price": 299,
			"website": "lco.in",
			"tags": ["web-dev", "frontend"]
		}
	`)

	var courses course // Define a variable to store decoded data

	// Check if the JSON data is valid
	isVaildJson := json.Valid(jsonDataFromWeb)
	fmt.Println("Is JSON valid: ", isVaildJson)

	if isVaildJson {
		// Decode JSON data into the course struct
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v\n", courses) // Print struct with field names
	} else {
		fmt.Println("Json was not valid")
	}

	// Decoding JSON into a map (key-value pair storage)
	var onlineJsondata map[string]interface{} // Interface allows any value type
	json.Unmarshal(jsonDataFromWeb, &onlineJsondata)
	fmt.Printf("%#v\n", onlineJsondata) // Print JSON as a map

	// Loop through the map and print each key-value pair
	for k, v := range onlineJsondata {
		fmt.Printf("key is %v and value is %v \n", k, v)
	}
}
