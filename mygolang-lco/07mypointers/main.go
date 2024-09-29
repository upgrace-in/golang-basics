package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Vlaue of pointer is: ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Address of this value is: ", ptr)
	fmt.Println("Address of this value is: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber)
}
