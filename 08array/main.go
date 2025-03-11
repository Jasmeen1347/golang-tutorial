package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitsArray [3]string // must have to define length
	fruitsArray[0] = "apple"
	fruitsArray[1] = "banana"
	fruitsArray[2] = "mango"

	fmt.Println("Fruits are: ", fruitsArray)

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegetables are: ", vegList)
	fmt.Println("Number of  vegetables are: ", len(vegList))
}
