package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	sayGreet() // call the function defined in hello-world.go
	variable()
}

// In a Go main package (an executable program),
// only one file contains the main() function — the program’s entry point.
