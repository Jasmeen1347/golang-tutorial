package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Maths")

	var numberOne int = 4
	var numberTwo float64 = 3.2

	fmt.Println("sum of this two: ", numberOne+int(numberTwo))

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5))
}
