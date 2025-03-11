package main

import "fmt"

// Defining a struct named 'User' with four fields
type User struct {
	Name   string // Name of the user
	Email  string // Email of the user
	Status bool   // Active status (true/false)
	Age    int    // Age of the user
}

func main() {
	fmt.Println("Structs")

	// Creating an instance of the User struct with values
	jasmeen := User{"Jasmeen", "jasmeen@go.dev", true, 20}

	// Print the struct directly (default formatting)
	fmt.Println(jasmeen)

	// Print struct with field names for better readability
	fmt.Printf("jasmeen's details are: %+v \n", jasmeen)

	// Accessing and printing a specific field (Name) from the struct
	fmt.Println(jasmeen.Name)
}
