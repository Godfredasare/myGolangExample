package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Print(err)
		return
	}

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Print(err)
		return
	}

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Print(err)
		return
	}
	// ebt := revenue - expenses
	// profit := ebt * (1-taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calcValues(revenue, expenses, taxRate)
	writeCalcValues(ebt, profit, ratio)

	fmt.Printf("Earning before tax: %.1f\n", ebt)

	fmt.Printf("Profit: %.1f\n", profit)

	fmt.Printf("Ratio: %.1f\n", ratio)
}

func getUserInput(text string) (float64, error) {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)

	if input <= 0 {
		return 0, errors.New("Value must be a positive")
	}
	return input, nil
}

func calcValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	p := ebt * (1 - taxRate/100)
	r := ebt / p

	return ebt, p, r
}

func writeCalcValues(ebt, profit, ratio float64) {
	values := fmt.Sprintf("EBT: %.1f,  Profit: %.2f,  Ratio: %.3f", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(values), 0644)
}
