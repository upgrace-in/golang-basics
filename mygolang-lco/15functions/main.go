package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	// result := adder(2,3);
	result, response := proAdder(2, 3, 4, 5)
	fmt.Println("Result is:", result, response)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hii!"
}

func greeter() {
	fmt.Println("Namastey from golang")
}

// Immediately Invoked & Anonymous function does exists
