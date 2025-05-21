package main

import "fmt"

func calulator2() {
	var num1, num2 int
	var opt string
	var result int

	fmt.Println("Choose Operation to Perform:")

	// Array for Options
	options := []string{"Sum", "Subtraction", "Division", "Multiply"}

	for i := range len(options) {
		fmt.Printf("%v\n", options[i])
	}

	fmt.Println("Enter your Operation")
	fmt.Scan(&opt)

	fmt.Print("Enter First Number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter Second Number: ")
	fmt.Scan(&num2)

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
	case "multiply", "*":
		result = num1 * num2
	default:
		fmt.Printf("Invalid Operator -> '%v' ", opt)
		return

	}

	fmt.Printf("Result is: %d\n", result)
}
