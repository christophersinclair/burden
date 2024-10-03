package main

import (
	"fmt"
	"os"
	"strconv"
)

func getFib() int {
	fib := os.Getenv("FIB")
	if fib == "" {
		fmt.Printf("No FIB environment variable found, using default value of 40\n")
		fib = "40"
	}

	fibInt, err := strconv.Atoi(fib)
	if err != nil {
		fmt.Printf("Error converting FIB value to integer: %v\nUsing default value of 40\n", err)
	}

	return fibInt
}

func getMaxThreads() int {
	maxThreads := os.Getenv("MAX_THREADS")
	if maxThreads == "" {
		fmt.Printf("No MAX_THREADS environment variable found, using default value of 1000\n")
		maxThreads = "1000"
	}

	maxThreadsInt, err := strconv.Atoi(maxThreads)
	if err != nil {
		fmt.Printf("Error converting MAX_THREADS value to integer: %v\nUsing default value of 1000\n", err)
		maxThreadsInt = 1000
	}

	return maxThreadsInt
}
