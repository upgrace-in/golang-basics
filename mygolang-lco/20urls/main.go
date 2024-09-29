package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://upgrace.in:3000/buy?product=FSWD&paymentID=1g0g1"

func main() {
	fmt.Println("Welcoem to handling URLs in Golang")
	fmt.Println(myurl)

	//parsing
	res, _ := url.Parse(myurl)
	fmt.Println(res.Scheme, res.Host, res.Path, res.Port(), res.RawQuery)

	qparams := res.Query()
	fmt.Println(qparams["product"])

	for key, val := range qparams {
		fmt.Println(key, "is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "upgrace.in",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
