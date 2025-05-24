package main

import "fmt"

func main() {
	s := []string{"apple", "orange", "banana"}

	fmt.Println("Slice contains:", s)
	s[1] = "Rohan"

	fmt.Println("Slice contains:", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
}
