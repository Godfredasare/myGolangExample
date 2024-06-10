package main

import "fmt"

//recursion function

func main() {
	fact := factoria(5)
	fmt.Println(fact)

	result :=addMore(10, 89,1)
	fmt.Print(result)
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

//accepting many values
func addMore(values ...int) int{
	total := 0
	for _, v := range values{
		total += v
	}
	return total
}
