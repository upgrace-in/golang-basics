package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	// we have to provide a size
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	// defualt it provides a space in each index
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}
