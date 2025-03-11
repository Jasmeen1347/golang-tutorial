package main

import "fmt"

func main() {
	// Defer statement delays execution until the surrounding function exits
	defer fmt.Println("Hello World") // Prints 4th (last)

	fmt.Println("Defer") // Prints 1st (normal execution)

	defer fmt.Println("Hello World2") // Prints 3rd (deferred execution)

	myDefer() // Executes before deferred statements in main(), prints 2nd
}

// Function demonstrating defer in a loop
func myDefer() {
	for i := 0; i < 5; i++ {
		// Deferred statements are pushed onto a stack, so they execute in LIFO order
		defer fmt.Println(i) // Prints in reverse order: 4, 3, 2, 1, 0
	}
}
