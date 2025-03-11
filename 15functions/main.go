package main

import "fmt"

func main() {
	fmt.Println("functions")
	greeter()

	result := adder(3, 4)

	fmt.Println("result: ", result)

	complexAdder := proAdder(5, 6, 4, 3)

	fmt.Println("result: ", complexAdder)

	addition, msg := multiValueReturn(5, 6, 4, 3)
	fmt.Println("result: ", addition, msg)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for value := range values {
		total = total + values[value]
	}

	return total
}

func multiValueReturn(values ...int) (int, string) {
	total := 0

	for value := range values {
		total = total + values[value]
	}

	return total, "hey bro"
}

func greeter() {
	fmt.Println("Namstey form golanag")
}
