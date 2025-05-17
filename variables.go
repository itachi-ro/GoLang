package main

import "fmt"

func variable() {
	var b, c int = 1, 2
	fmt.Println(b, c) // b = 1, c = 2

	// string
	var a = "initial"
	fmt.Println(a)

	// Go's short variable declaration operator := works with every type
	// Works only with var, not with const

	x := "initial"
	fmt.Println(x)

	// []string (slice)
	names := []string{"Alice", "Bob"}
	fmt.Println(names) // Output: [Alice Bob]

	// map[string]int
	scores := map[string]int{"Alice": 90, "Bob": 85}
	fmt.Println(scores) // Output: map[Alice:90 Bob:85]
}
