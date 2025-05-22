package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	var rohan Person
	fmt.Println("Rohan Person: ", rohan)

	rohan.FirstName = "Rohan"
	rohan.LastName = "Nautiyal"
	rohan.Age = 25

	fmt.Println("Rohan Details: ", rohan)

	person1 := Person{
		FirstName: "Ron",
		LastName:  "Luther",
		Age:       25,
	}

	fmt.Println("Ron Details: ", person1)

	// new keyword
	var person2 = new(Person)

	person2.FirstName = "Simran"
	person2.LastName = "Singh"
	person2.Age = 25

	// Print pointer (shows memory address)
	fmt.Println("Person2 Details: ", person2) // &{Simran Singh 25}

	/*
		& → shows that it's a pointer
		{Simran Singh 25} → the actual values inside the struct that the
		pointer is pointing to.
		When we write this and gets that ouput it basically
		means that this is the pointer to struct Person which has these values.
	*/

	// Dereference pointer (shows actual struct values)
	fmt.Println("Person2 Details: ", *person2) // {Simran Singh 25}

	// Address of the pointer itself (pointer to a pointer)
	fmt.Println("Person2 Details: ", &person2) //  0xc000048068

	fmt.Println("Person2 Details: ", person2.FirstName) // {Simran Singh 25}

}

/*

 Code	   					Meaning	     			      Output Example

person2				Pointer to the struct				&{Simran Singh 25}
*person2		Dereferenced pointer, actual values		{Simran Singh 25}
&person2		Address of the pointer itself			0xc00000e028

*/
