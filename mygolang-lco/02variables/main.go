package main

import "fmt"

// we cam decalre constants
// using first letter as capital states that it is a public variale
const LoginToken string = "sdadasdasda"

func main() {
	var username string = "Hari"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var smallVal bool = false
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type
	var website = "upgrace.in"
	fmt.Println(website)   
	fmt.Printf("Variable is of type: %T \n", website)

	// no var style
	numberOfUser := 30
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
 }  
