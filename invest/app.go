package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Rate: ")
	fmt.Scan(&expectedRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

     futureValue, futureRealValue:= calcFutureValues(investmentAmount,expectedRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedRate/100, years)
	//futureRealValue := futureValue * math.Pow(1+ inflationRate / 100, years)

	fmt.Printf("future Value: %.0f\n", futureValue)

	fmt.Printf("future Real Value: %.1f", futureRealValue)
}

//returning multiple values
func calcFutureValues(investmentAmount, expectedRate, years float64) (float64, float64){
	fv := investmentAmount * math.Pow(1+expectedRate/100, years)
	rfv := fv * math.Pow(1+ inflationRate / 100, years)

	return fv, rfv
}
