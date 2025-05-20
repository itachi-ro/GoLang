package main

import (
	"fmt"
)

func do() {
	var name string

	fmt.Print("Enter your Name: \n")
	fmt.Scan(&name) // Rohan Nautiyal

	fmt.Printf("Hello, %v\n", name) // Hello, Rohan
}

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// reader := bufio.NewReader(os.Stdin)
// var name string

// fmt.Print("Enter your Name: \n")
// name, _ = reader.ReadString('\n')
// name = strings.TrimSpace(name)

// fmt.Printf("Hello, %v\n", name)
