package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Handling")

	presentTime := time.Now()

	fmt.Println("current time: ", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdData := time.Date(2012, time.August, 15, 12, 15, 0, 0, time.UTC)

	fmt.Println(createdData)

}
