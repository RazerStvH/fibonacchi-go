package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
)

// A function for calculating Fibonacci numbers from 0 to n
func fibonacciSequence(n int) []*big.Int {
	sequence := make([]*big.Int, n+1)
	for i := range sequence {
		sequence[i] = big.NewInt(0)
	}

	if n >= 0 {
		sequence[0] = big.NewInt(0)
	}
	if n >= 1 {
		sequence[1] = big.NewInt(1)
	}

	for i := 2; i <= n; i++ {
		sequence[i].Set(sequence[i-1])
		sequence[i].Add(sequence[i], sequence[i-2])
	}

	return sequence
}

func main() {
	var n int
	var outputFileName string

	// Defining flags
	flag.IntVar(&n, "n", 10, "The number up to which to calculate the Fibonacci sequence")
	flag.StringVar(&outputFileName, "o", "output.txt", "The name of the output file to record the result")
	flag.Parse()

	// Calculating the sequence
	sequence := fibonacciSequence(n)

	// Opening the file for writing
	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	// Writing the sequence to a file
	for i, num := range sequence {
		_, err := file.WriteString(fmt.Sprintf("Fibonacci(%d) = %s\n", i, num.String()))
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return
		}
	}

	fmt.Printf("Fibonacci numbers up to %d are written to the %s file\n", n, outputFileName)
}
