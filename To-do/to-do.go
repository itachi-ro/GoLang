package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toDo() {

	reader := bufio.NewReader(os.Stdin)

	var todoList []string

	for {
		fmt.Println("<-------To-Do List (In Memory)-------->")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Task")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Exit")
		fmt.Println("Choose Option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("Enter task :")
			todo, _ := reader.ReadString('\n')
			todo = strings.TrimSpace(todo)
			todoList = append(todoList, todo)
			fmt.Println("Task Added Successfully!")

		case "2":
			if len(todoList) == 0 {
				fmt.Println("No tasks found")
			} else {
				fmt.Println("Your task")
				for i, task := range todoList {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
		case "3":
			if len(todoList) == 0 {
				fmt.Println("No tasks to delete")
			}

			fmt.Println("Enter task number to delete: ")
			indexValue, _ := reader.ReadString('\n')
			index, err := strconv.Atoi(strings.TrimSpace(indexValue))

			if err != nil || index < 1 || index > len(todoList) {
				fmt.Println("Invalid task Number!")
			} else {
				// fmt.Println(todoList[:index-1]) // If (index = 2) todoList[:1] element before index 1 i.e -> "Learning"
				// fmt.Println(todoList[index:])   // todoList[2:] element after index 2 i.e -> "learn go" "golang"
				todoList = append(todoList[:index-1], todoList[index:]...)
				fmt.Println("Task deleted.")
			}
		case "4":
			fmt.Println("Goodbye, Exiting")
			return
		default:
			fmt.Println("Invalid Choice.")
		}
	}
}

/*

In Todo List, have learned about ASCII to integer Conversion by strconv.Atoi() and
how element is deleted from an slice.


-----------Use of strconv.Atoi()-----------

	1. Converts a string to an integer

	indexInput is a string (read from user input using bufio.Reader).
    strconv.Atoi() stands for ASCII to Integer — it tries to convert a string like "3" to the integer 3.

	2. Handles possible errors

	If the input can't be converted to an integer (e.g. "abc"), strconv.Atoi() returns an error.
    err will be nil if conversion is successful, or an error object if it fails.

-----------Explaination of Deletion in Slice----------

Let's say you have a slice: todoList := []string{"Task1", "Task2", "Task3", "Task4"}
Suppose you want to delete the element at index = 2

What this does:
todoList = append(todoList[:index-1], todoList[index:]...)

Let’s say index = 2, so:

    todoList[:1] → elements before index 1 ("Task1")

    todoList[2:] → elements after index 1 ("Task3", "Task4")

Then append combines them, skipping over "Task2".

IMP to note --
todoList[index:] in Go means:

    "From index index to the end of the slice."

It includes the element at index, not the next one.
todoList[index:] = it means “current onward.”

*/
