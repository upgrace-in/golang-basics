package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Soemthigne lse"
	}

	// assign on the go
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}

	// if err != nil {

	// }

	fmt.Println(result)
}
