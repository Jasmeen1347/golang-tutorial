package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	// Creating a slice of strings with different fruit names
	var fruitList = []string{"apple", "banana", "mango", "pineapple", "grapes"}
	fmt.Printf("Type of fruitList is: %T \n", fruitList)

	// Adding more elements to the slice using append
	fruitList = append(fruitList, "kiwi", "orange")
	fmt.Println("Fruits are: ", fruitList)

	// Slicing the slice to exclude the first element (index 0)
	fruitList = append(fruitList[1:])
	fmt.Println("Fruits after slicing from index 1: ", fruitList)

	// Slicing again to include only elements from index 1 to 2 (excluding index 3)
	fruitList = append(fruitList[1:3])
	fmt.Println("Fruits after further slicing: ", fruitList)

	// Creating a slice of integers with predefined length (4 elements)
	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 565
	highScore[3] = 12
	// highScore[4] = 99 // Uncommenting this line will cause an error since index 4 does not exist

	fmt.Println("High scores: ", highScore)

	// Appending more values to the slice, dynamically increasing its size
	highScore = append(highScore, 555, 666)
	fmt.Println("High scores after appending: ", highScore)

	// Checking if the slice is sorted
	fmt.Println("Is sorted: ", sort.IntsAreSorted(highScore))

	// Sorting the slice in ascending order
	sort.Ints(highScore)
	fmt.Println("High scores after sorting: ", highScore)
	fmt.Println("Is sorted: ", sort.IntsAreSorted(highScore))

	// Creating a slice of course names
	var courses = []string{"reactjs", "javascript", "python", "java", "c++"}
	fmt.Println("List of courses: ", courses)

	// Removing an element from the slice based on index
	var index int = 2                                       // Removing "python" (index 2)
	courses = append(courses[:index], courses[index+1:]...) // Concatenating slices before and after the index
	fmt.Println("List of courses after removal: ", courses)
}
