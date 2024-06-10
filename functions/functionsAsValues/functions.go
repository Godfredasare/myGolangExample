package functionsasvalues

import "fmt"

//functions as values

type numbers []int
type transformFn func(int) int

func Main() {
	numbers := numbers{1, 2, 3, 4, 5}
	double := transformNumbers(&numbers, double)
	trible := transformNumbers(&numbers, trible)
	fmt.Println(double)
	fmt.Println(trible)
}

func transformNumbers(num *numbers, transform transformFn) []int {
	dNumbers := []int{}
	for _, v := range *num {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func trible(number int) int {
	return number * 3
}
