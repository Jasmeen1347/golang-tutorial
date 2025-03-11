package main

import "fmt"

func main() {
	fmt.Println("If Else")

	loginCount := 25
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if num := 10; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Greater than 10")
	}
}
