package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://upgrace.in"

func main() {
	fmt.Println("Web Requests in Golang")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)

	// it never gets close so we have to do it
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}	

	fmt.Println(string(body))
}
