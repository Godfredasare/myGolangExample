package main

import "fmt"

func main() {
	var choice int
	var balance float64 = 200

	fmt.Println("Welcome To GH Bank")
	fmt.Println("__________________")
	println(" ")

	for {
		fmt.Println(" ")
		fmt.Println("1: Check Balance")
		fmt.Println("2: Deposit Money")
		fmt.Println("3: Withdraw Money")
		fmt.Println("4: Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

        
		if choice == 1 {
			fmt.Printf("Your account balance is: %.1f", balance)
			fmt.Println(" ")
		} else if choice == 2 {
			Deposit := getInput("Enter deposite amount: ")
			balance += Deposit
			fmt.Printf("You have deposited: %.2f into your account, your new balance is: %.2f", Deposit, balance)
			fmt.Println(" ")
		} else if choice == 3 {
			withdraw := getInput("Enter amount to withdraw: ")
			balance -= withdraw
			fmt.Printf("You have withdraw: %.2f from your account, your new balance is: %.2f\n", withdraw, balance)
			fmt.Println(" ")
		} else {
			return
		}
	}
	// fmt.Println(balance)

}

func getInput(text string) float64 {
	fmt.Print(text)
	var value float64
	fmt.Scan(&value)
	return value
}
