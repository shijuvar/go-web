// Example program with unbuffered chanel
package main

import (
	"fmt"
	"sync"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func main() {

	count := make(chan int)
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	fmt.Println("Start Goroutines")
	//launch a goroutine with label "A"
	go printCounts("A", count)
	//launch a goroutine with label "B"
	go printCounts("B", count)
	fmt.Println("Channel begin")
	count <- 1
	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

func printCounts(label string, count chan int) {
	// Schedule the call to WaitGroup's Done to tell we are done.
	defer wg.Done()
	for {
		//Receives message from Channel
		val, ok := <-count
		if !ok {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel Closed from %s \n", label)
			// Close the channel
			close(count)
			return
		}
		val++
		// Send count back to the other goroutine.
		count <- val
	}
}
