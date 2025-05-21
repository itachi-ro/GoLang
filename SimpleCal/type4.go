package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator4() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first number: ")
	input1, _ := reader.ReadString('\n')
	num1, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	fmt.Print("Enter operator (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	operator := strings.TrimSpace(op)

	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')
	num2, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
