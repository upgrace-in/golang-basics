package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Requests in Golang")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformGetRequest() {
	url := "http://localhost:3000/get"

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

	var responseString strings.Builder
	responseString.Write(body)
	fmt.Println(responseString.String())
	// fmt.Println(string(body))
}

func PerformPostJsonRequest() {
	const url = "http://localhost:3000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform": "learnCodeOnline.in"
		}
	`)

	res, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	
	var resStrings strings.Builder
	body, _ := io.ReadAll(res.Body)
	resStrings.Write(body)
	fmt.Println(resStrings.String())
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postForm"

	// fake json payload
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "xyz")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}

