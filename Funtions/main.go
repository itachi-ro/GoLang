package main

import "fmt"

func main() {
	sq := square(5)
	sq7 := square(7)
	fmt.Printf("Square of 7 is: %v\n", sq7)
	fmt.Print("Square of 6 is: ", square(6), "\n")
	fmt.Printf("Square of 5 is: %v\n", sq)

	sum, multiply := addAndMultiply(2, 3)
	fmt.Printf("Sum is %v and Multiply is %v\n", sum, multiply)

	// fucntiosn with unlimited args
	fmt.Println("Sum all: ", sumAll(1, 2, 3, 4, 5))
}

// Rule in Go:
// Function names are immutable identifiers — you cannot assign values to them.
