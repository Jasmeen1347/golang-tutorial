package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Global slice to store results from our HTTP requests
var signals = []string{"test"}

// WaitGroup is used to wait for all goroutines to finish
// Think of it as a counter: we add to it when starting work and subtract when work is done
var wg sync.WaitGroup

// Mutex (mutual exclusion) is used to protect shared resources from concurrent access
// It works like a lock on a door - only one goroutine can access the protected data at a time
var mut sync.Mutex

func main() {
	fmt.Println("Goroutines")

	// Commented out code showing a simple example of goroutines
	// go greeter("Hello") // This would run in a separate goroutine (background)
	// greeter("World")    // This would run in the main goroutine

	// List of websites we want to check
	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://facebook.com",
		"https://github.dev",
		"https://amazon.com",
	}

	// Loop through each website and check its status code
	for _, website := range websiteList {
		// Launch a new goroutine for each website request
		// This allows multiple HTTP requests to run concurrently (in parallel)
		go getStatusCode(website)

		// Increment the WaitGroup counter for each goroutine we start
		// This tells the program "we have one more task that needs to finish"
		wg.Add(1)
	}

	// Wait for all goroutines to complete
	// This blocks the main function from ending until all goroutines call wg.Done()
	wg.Wait()

	// Print the final results
	fmt.Println(signals)
}

// Commented out example function to demonstrate goroutines
// func greeter(s string) {
//     for i := 0; i < 6; i++ {
//         fmt.Println(s)
//     }
// }

// Function to get HTTP status code from a website
func getStatusCode(endpoint string) {
	// This ensures wg.Done() is called when this function completes
	// Decrements the WaitGroup counter, signaling "this task is finished"
	defer wg.Done()

	// Make HTTP GET request to the website
	result, err := http.Get(endpoint)

	if err != nil {
		// If there's an error, print it
		fmt.Println("OOPS", err)
	} else {
		// Before modifying the shared 'signals' slice, we need to lock the mutex
		// This prevents other goroutines from accessing 'signals' at the same time
		mut.Lock()

		// Add the website to our results slice
		signals = append(signals, endpoint)

		// Print the status code and website
		fmt.Printf("%d status code for website %s \n", result.StatusCode, endpoint)

		// Unlock the mutex so other goroutines can access 'signals'
		mut.Unlock()
	}
}

/*
EXPLANATION OF KEY CONCEPTS:

1. Goroutines:
   - Goroutines are lightweight threads managed by the Go runtime
   - They allow functions to run concurrently (at the same time)
   - You create a goroutine by adding the 'go' keyword before a function call
   - Use goroutines when you want to do multiple tasks at once (like checking multiple websites)
   - Perfect for I/O operations that involve waiting (network requests, file operations)

2. WaitGroup (wg):
   - Used to wait for multiple goroutines to finish
   - wg.Add(n): Adds n to the counter
   - wg.Done(): Decrements counter by 1
   - wg.Wait(): Blocks until counter becomes 0
   - Use WaitGroup when you need to ensure all concurrent tasks complete before proceeding

3. Mutex (mut):
   - Protects shared data from being accessed by multiple goroutines at once
   - mut.Lock(): Acquire the lock (other goroutines must wait)
   - mut.Unlock(): Release the lock (other goroutines can proceed)
   - Use mutex when multiple goroutines need to modify shared data (like our 'signals' slice)
   - Without mutex, concurrent modifications can lead to race conditions and data corruption

When to use goroutines:
- For tasks that can run independently
- For operations that involve waiting (I/O, network)
- To utilize multiple CPU cores for better performance
- To keep your application responsive while doing heavy work

When to use mutex:
- When multiple goroutines need to access/modify the same data
- To prevent race conditions
- Only lock the mutex for the shortest time necessary
- Be careful to always unlock after locking to avoid deadlocks
*/
