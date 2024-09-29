package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("details are: %+v\n", hitesh)
	hitesh.getStatus()
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) getStatus() {
	fmt.Println("Status of this user is: ", u.Status)
}