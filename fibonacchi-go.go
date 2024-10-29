package main

import (
	"fmt"
)

// Recursive function for calculating the n-th Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Enter the number of the Fibonacci number: ")
	fmt.Scanln(&n)

	result := fibonacci(n)
	fmt.Printf("Fibonacci(%d) = %d\n", n, result)
}
