package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Go Buffio")
	fmt.Println("Movies Rating")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input= strings.TrimSpace(input)

	fmt.Println("My rating is: ", input)

	newInput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print("Rating plus 1: ", newInput+1)
}