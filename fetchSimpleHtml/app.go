package main

import (
	"fmt"
	"io"
	"net/http"
)

const MyUrl = "https://www.york.ac.uk/teaching/cws/wws/webpage1.html"

func main() {
	res, err := http.Get(MyUrl)
	if err != nil {
		panic(err)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	defer res.Body.Close()

	// url, err := url.Parse(MyUrl)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(url.ForceQuery)

}
