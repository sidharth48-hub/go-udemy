package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------")
	}
	fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	presentOptions()
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	for choice != 4 {

		switch choice {
		case 1:
			balance := checkBalance()
			fmt.Printf("Your account balance: %0.2f\n", balance)
		case 2:
			var deposit float64
			fmt.Print("Amount to deposit: ")
			fmt.Scan(&deposit)
			depositMoney(deposit)
			balance := checkBalance()
			fmt.Printf("Your account balance: %0.2f\n", balance)
		case 3:
			var withdrawMoneyAmount float64
			fmt.Print("Amount to withdraw: ")
			fmt.Scan(&withdrawMoneyAmount)
			withdrawMoney(withdrawMoneyAmount)
			balance := checkBalance()
			fmt.Printf("Your account balance: %0.2f\n", balance)
		}

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
	}

	println("Thank you for using the GO Bank")
}

func checkBalance() float64 {
	balance, _ := fileops.GetFloatFromFile(accountBalanceFile)
	return balance
}

func depositMoney(depositAmount float64) {

	balance := checkBalance()
	balance += depositAmount
	fileops.WriteFloatToFile(balance, accountBalanceFile)
	fmt.Println("Amount successfully deposited")
}

func withdrawMoney(withdrawMoneyAmount float64) {
	balance := checkBalance()

	if balance < withdrawMoneyAmount {
		fmt.Println("You do not have the required balance")
		//fmt.Printf("Your account balance: %0.2f\n", balance)
		return
	}

	balance -= withdrawMoneyAmount
	fileops.WriteFloatToFile(balance, accountBalanceFile)
	fmt.Printf("%0.2f Amount successfully withdrawn\n", withdrawMoneyAmount)
}
