package main

import "fmt"

func main() {
	numbers := [3]int{1, 2, 3}
	numbers[0]= 5
	fmt.Println(numbers)
	
	n := numbers[0:]
	fmt.Println(n)

	sliceNnumbers := []int{5,10,15,20,30}

	sliceNnumbers = append(sliceNnumbers, 90,100)

	fmt.Println(sliceNnumbers)

	//how to remove a value from slices base on the index
	fruits := []string{"Apple", "Orange","Banana","Mango","Pawpaw"}
	fmt.Println(fruits)
	index :=2
	fruits = append(fruits[:index], fruits[index+1:]...)
	fmt.Println(fruits)
}