package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to userinput program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating for our pizza: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)

}
