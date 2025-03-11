package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// Declaring a pointer variable (initially nil)
	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23 // Declare and initialize an integer variable

	// Create a pointer that stores the address of myNumber
	var ptr = &myNumber

	// Print the memory address stored in the pointer
	fmt.Println("Value of actual pointer (memory address) is ", ptr)

	// Print the value stored at the memory address (dereferencing the pointer)
	fmt.Println("Value at the memory address is ", *ptr)

	// Modify the value at the memory address (dereferencing and updating)
	*ptr = *ptr * 2

	// Since ptr points to myNumber, changing *ptr also changes myNumber
	fmt.Println("New value of myNumber is ", myNumber)
}
