package main

import "fmt"

//recursion function

func main() {
	fact := factoria(5)
	fmt.Println(fact)
}

func factoria(number int) int {
	if number == 0 {
		return 1
	}
	return number * factoria(number-1)

	//normal way of achieving this
	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }
	// return result
}
