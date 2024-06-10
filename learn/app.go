package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := `https://pkg.go.dev`
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
