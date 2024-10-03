package main

func busyWork(fib int) {
	// Calculate nth Fibonacci number
	fibonacci(fib)
}

// A recursive mathematical function that is CPU-bound
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
