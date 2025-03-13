package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	// Create a buffered channel that can hold up to 2 integers
	// Buffered means it can store values without a receiver being ready
	myCh := make(chan int, 2)

	// Create a WaitGroup pointer to coordinate goroutines
	wg := &sync.WaitGroup{}

	// These lines are commented out to show simple channel operations:
	// myCh <- 23         // Send value 23 to channel
	// fmt.Println(<-myCh) // Receive value from channel and print it

	// Add 2 to WaitGroup counter (for 2 goroutines)
	wg.Add(2)

	// First goroutine - RECEIVER
	// <-chan int means this goroutine can only receive from the channel (read-only)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// Receive value from channel and check if channel is still open
		value, isChannelOpen := <-myCh

		// Try to read another value from channel
		// This will either get the second value or zero value if channel is closed
		fmt.Println(<-myCh)

		// Print the first value received
		fmt.Println(value)

		// Print whether the channel is open (true) or closed (false)
		fmt.Println(isChannelOpen)

		// Mark this goroutine as done
		wg.Done()
	}(myCh, wg)

	// Second goroutine - SENDER
	// chan<- int means this goroutine can only send to the channel (write-only)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// Send value 23 to the channel
		myCh <- 23

		// Close the channel - no more values can be sent
		// Receivers can still read until channel is empty
		close(myCh)

		// This line is commented out because sending on closed channel causes panic:
		// myCh <- 24

		// Mark this goroutine as done
		wg.Done()
	}(myCh, wg)

	// Wait for both goroutines to finish before exiting
	wg.Wait()
}

/*
EXPLANATION OF CHANNELS:

1. What are Channels?
   - Channels are typed conduits that allow goroutines to communicate with each other
   - They provide a way to send and receive values between goroutines safely
   - Think of them as pipes where one goroutine puts data in, and another takes data out

2. Types of Channels:
   - Unbuffered channel: make(chan int) - no capacity, sends block until someone receives
   - Buffered channel: make(chan int, N) - can hold N values before blocking
   - Read-only channel: <-chan int - can only receive values
   - Write-only channel: chan<- int - can only send values

3. Channel Operations:
   - Send: ch <- value
   - Receive: value := <-ch
   - Close: close(ch)
   - Check if closed: value, ok := <-ch (ok is false if channel closed)
   - Range over channel: for value := range ch

4. Program Execution Flow:
   - Main creates a buffered channel and starts two goroutines
   - Sender goroutine puts 23 into channel and closes it
   - Receiver goroutine gets 23 and then tries to read again
   - Since channel is closed, second read returns zero value
   - When reading from closed channel, ok value is false

5. When to Use Channels:
   - For communication between goroutines
   - For synchronization (as alternative to mutex in some cases)
   - For signaling completion (instead of WaitGroup sometimes)
   - For controlling concurrency (limit number of concurrent operations)
   - For implementing pipelines (chain of processing steps)
   - When data needs to flow between concurrent tasks

6. Channel Best Practices:
   - Use directional channels (chan<-, <-chan) when possible for type safety
   - Close channels from the sender side, not receiver
   - Don't close a channel if there are multiple senders
   - Check if a channel is closed before receiving
   - Use select statement for handling multiple channels
   - Remember: "Don't communicate by sharing memory; share memory by communicating"
*/
