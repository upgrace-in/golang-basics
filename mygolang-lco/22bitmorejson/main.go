package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string `json:"-"` // this will never show up in the api
	Tags     []string `json:"tags,omitempty"` // it will remove the null value
}

func main() {
	fmt.Println("Welcome to JSON Video")
	// EncodeJSON()
	DecodeJson()
}

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web", "web-dev", "js"}},
		{"NextJS Bootcamp", 199, "LearnCodeOnline.in", "bcd123", nil},
	}

	// package this data as JSON data

	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")
	// fmt.Println(string(finalJson))
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "tags": [
                        "web",
                        "web-dev",
                        "js"
                ]
        }
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid!")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid!")
	}

	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Println(k, v)
	}
}
