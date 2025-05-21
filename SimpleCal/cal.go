package main

import "fmt"

func calulator() {
	var num1, num2 int

	fmt.Println("What Operation do you want to Perform")
	fmt.Print("1. Sum\n2. Subtraction\n3. Division\n4. Multiply\n")

	var opt string
	fmt.Println("Enter your Operation")
	fmt.Scan(&opt)

	fmt.Print("Enter First Number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number: ")
	fmt.Scan(&num2)

	var result int
	switch opt {
	case "sum", "+":
		result = num1 + num2
	case "subtraction", "sub", "-":
		result = num1 - num2
	case "division", "/":
		if num2 == 0 {
			fmt.Println("Cannot Divide by zero")
			return
		}
		result = num1 / num2
	case "multiply":
		result = num1 * num2

	default:
		fmt.Printf("Invalid Operator -> '%v' ", opt)
		return

	}
	fmt.Printf("Result: %d", result)
}
