package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// Creating a map to store programming languages with their abbreviations as keys
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	languages["GO"] = "Golang"
	languages["CS"] = "C#"

	// Printing the entire map
	fmt.Println("List of all languages: ", languages)

	// Accessing and printing a specific value using its key
	fmt.Println("Language for PY key: ", languages["PY"])

	// Deleting a key-value pair from the map
	delete(languages, "JS")
	fmt.Println("Languages after deleting JS: ", languages)

	// Iterating over the map and printing all key-value pairs
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
