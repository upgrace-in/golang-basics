package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hare Krsna"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err syntax -> used alternative for trycatch
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)
}
