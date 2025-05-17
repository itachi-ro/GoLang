package main

import "fmt"

func do() {
	var name string
	fmt.Print("Enter your Name\n")
	fmt.Scan(&name)
	fmt.Print("Hello ", name)
}
