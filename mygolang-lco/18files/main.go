package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file = LearnCodeOnline.in"

	fileName := "./myfile.txt"
	file, err := os.Create(fileName)
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length returned: ", length)
	defer file.Close()

	readFile(fileName)
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println(string(dataByte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err) // it will just shutdown the execution of program
	}
}
