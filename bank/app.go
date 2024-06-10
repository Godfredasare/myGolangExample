package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func readBalanceFile() (float64, error){
	data, err:= os.ReadFile(fileName)
	if err != nil {
		return 100, errors.New("fialed reading file!, file not found")
	}
	dataText := string(data)
	balance, _ := strconv.ParseFloat(dataText, 64)
	return balance, nil
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func main() {

	var choice int
	accountBalance, err := readBalanceFile()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	var deposit float64
	var withdraw float64

	fmt.Println("Welcome To Go Bank")
	fmt.Println("__________________")

	for {
		fmt.Println("What do you what to do")
		fmt.Println("1: Check balance")
		fmt.Println("2: Deposit money")
		fmt.Println("3: Withdraw money")
		fmt.Println("4: Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your account balance is: %0.1f\n", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				continue
			}

			accountBalance += deposit
			fmt.Printf("Balance updated, New amount: %0.1f\n", accountBalance)
			writeBalanceToFile(accountBalance)

		} else if choice == 3 {
			fmt.Print("How much do you want to withdraw: ")
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				continue
			}

			if withdraw > accountBalance {
				fmt.Println("Invalid amount, You can`t withdraw more than you have")
				continue
			}

			accountBalance -= withdraw
			fmt.Printf("Balance updated, New amount: %0.1f\n", accountBalance)
			writeBalanceToFile(accountBalance)

		} else {
			fmt.Println("Goodbye")
			break
		}
		fmt.Print("Thanks for choosing our bank")
	}
}
