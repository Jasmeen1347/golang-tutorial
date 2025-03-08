package main

import "fmt"

const PI float32 = 3.14

func main() {
	var username string = "Jasmeen"
	fmt.Println(username)
	fmt.Printf("Variable is Type of: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is Type of: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is Type of: %T \n", smallValue)

	var acceptAllInt int = 255555
	fmt.Println(acceptAllInt)
	fmt.Printf("Variable is Type of: %T \n", acceptAllInt)

	var smallFloat float32 = 255.555789531845161
	fmt.Println(smallFloat)
	fmt.Printf("Variable is Type of: %T \n", smallFloat)

	var bigFloat float64 = 255.555789531845161
	fmt.Println(bigFloat)
	fmt.Printf("Variable is Type of: %T \n", bigFloat)

	// impicit
	var name = "man"
	fmt.Println(name)

	// no var style
	numberOfUsers := 1000
	fmt.Println(numberOfUsers)

	fmt.Println("PI: ", PI)
}
