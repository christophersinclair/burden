package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	greeting := `
      ..-::::::-..          
  .:-::::::::::::::-:.
  ._:::    ::    :::_.
   .:( ^   :: ^   ):.             __                   __                
   .:::   (..)   :::.            / /_  __  ___________/ /__  ____          
   .:::::::UU:::::::.           / __ \/ / / / ___/ __  / _ \/ __ \
   .::::::::::::::::.          / /_/ / /_/ / /  / /_/ /  __/ / / /
   -::::::::::::::::-         /_.___/\__,_/_/   \__,_/\___/_/ /_/ 
   .::::::::::::::::.
    .::::::::::::::.
      oO:::::::Oo
      
  Welcome to burden! Author: @christophersinclair
      `

	fmt.Println(greeting)

	// Get environment variables
	fib := getFib()
	maxThreads := getMaxThreads()

	// Print the number of CPUs available
	fmt.Printf("Running on %d CPU cores...\n", runtime.NumCPU())

	// Create a channel to control the number of active goroutines
	// This is used to limit the number of goroutines that are running at any given time
	// An empty struct in Go is ZERO bytes - so by using this channel, we're not using any memory
	threadLimiter := make(chan struct{}, maxThreads)

	fmt.Printf("Starting %d goroutines, each calculating the %dth Fibonacci number...\n", maxThreads, fib)
	fmt.Printf("%d goroutines will be maintained until the Pod is destroyed (or program is halted).\n", maxThreads)

	for {
		// Add a placeholder to the threadLimiter channel
		// When the channel is full, this will block until a goroutine is removed
		threadLimiter <- struct{}{}

		go func(fib int) {
			busyWork(fib)   // spawn a new goroutine that calculates the nth Fibonacci number
			<-threadLimiter // remove a goroutine from the blocking channel when it's done
		}(fib)

		time.Sleep(100 * time.Millisecond)
	}
}
