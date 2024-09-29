package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch & case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
		//fallthrough // it will let the next conditon run too
	default: 
		fmt.Println("What was that?")
	}
}
