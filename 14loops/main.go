package main

import "fmt"

func main() {
	fmt.Println("Loops")

	// Defining a slice (dynamic array) of strings
	days := []string{"apple", "banana", "orange", "kiwi"}

	// Using a traditional for loop with an index
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d]) // Accessing elements using index
	}

	// Using a for loop with range (index-based iteration)
	for i := range days {
		fmt.Println(days[i]) // Accessing elements using index
	}

	// Using for loop with range (index and value directly)
	for index, day := range days {
		fmt.Println(index, day) // Printing both index and value
	}

	tmp := 1 // Initializing a variable

	// Loop similar to a while loop (runs while condition is true)
	for tmp < 10 {

		if tmp == 2 {
			goto ico // Jump to the label "ico" when tmp is 2
		}
		if tmp == 5 {
			break // Exit the loop when tmp is 5
		}
		if tmp == 5 { // This condition will never be reached due to 'break'
			tmp++
			continue // Skip this iteration (not needed here)
		}

		fmt.Println(tmp) // Print the value of tmp
		tmp++            // Increment tmp
	}

	// Label for the goto statement
ico:
	fmt.Println("Hello World") // This will print when tmp == 2
}
