package main

import (
	"fmt"
)

func calculator3() {
	var num1, num2 int
	var opt string
	var result int

	fmt.Println("Simple Go Calculator (type 'exit' to quit)")

	// Loop keep executing until exit hits!!!!
	for {
		fmt.Println("\nOperations: sum (+), subtraction (-), division (/), multiply (*), exit")
		fmt.Print("Enter your operation: ")
		fmt.Scan(&opt)

		if opt == "exit" {
			fmt.Println("Exiting Calculator. Goodbye!")
			break
		}

		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		switch opt {
		case "sum", "+":
			result = num1 + num2
		case "subtraction", "-", "sub":
			result = num1 - num2
		case "division", "/":
			if num2 == 0 {
				fmt.Println("Error: Cannot divide by zero.")
				continue
			}
			result = num1 / num2
		case "multiply", "*":
			result = num1 * num2
		default:
			fmt.Printf("Invalid operator: '%v'\n", opt)
			continue
		}

		fmt.Printf("Result: %d\n", result)
	}
}
