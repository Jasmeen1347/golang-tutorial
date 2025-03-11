package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")

	content := "This need to go into file"

	file, err := os.Create("./mygo.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length is: ", length)

	defer file.Close()

	readFile("./mygo.txt")
}

func readFile(filepath string) {
	databyte, err := ioutil.ReadFile(filepath)

	checkNilError(err)

	fmt.Println("File content is: ", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}

}
