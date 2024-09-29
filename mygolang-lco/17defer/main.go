package main

import "fmt"

func main() {
	// LIFO principle followed
	defer fmt.Println("Bol")
	defer fmt.Println("Krsna")
	defer fmt.Println("Hare")
	fmt.Println("Oye")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
