package main

import "fmt"

func main() {
	web := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}

	fmt.Println(web)
	
	//add
	web["twitter"] = "https://twitter.com"
	fmt.Println(web)
	
	delete(web, "Google")
	fmt.Println(web)
}
