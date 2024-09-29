package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])
	delete(languages, "RB")
	fmt.Println(languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Println(key, value)
	}
}
