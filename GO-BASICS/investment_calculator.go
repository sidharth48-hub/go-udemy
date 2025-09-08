package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the Years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the Expected Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Real Value: %.2f", futureRealValue)
}
