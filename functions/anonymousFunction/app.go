package anonymousfunction

import "fmt"

type num []int

//anonymous functions in go

func Main() {
	numbers := num{1, 2, 3}
	double := transformNumbers(&numbers, func(i int) int {
		return i * 2
	})
	fmt.Println(double)
}

func transformNumbers(numbers *num, transform func(int) int) []int {
	dNumbers := []int{}

	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}
