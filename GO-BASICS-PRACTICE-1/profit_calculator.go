package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	taxRate := 0.4

	revenue = getUserInput("Revenue: ")

	expenses = getUserInput("Expenses: ")

	taxRate = getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateEBTProfitAndRatio(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInput(output string) (input float64) {
	fmt.Print(output)
	fmt.Scan(&input)
	return input
}

func calculateEBTProfitAndRatio(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate)

	ratio = ebt / profit

	return ebt, profit, ratio
}
