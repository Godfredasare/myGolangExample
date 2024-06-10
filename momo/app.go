package main

import (
	"fmt"
	"os"
)

func main() {
	var shortCode string
	var choice int
	var firstChoice int
	var number string
	var reEnterNumber string
	var amount float64
	var ref string
	var pin string

	fmt.Print("Enter Short code: ")
	fmt.Scan(&shortCode)

	for {
		if shortCode == "*170#" {
			fmt.Println("Menu")
			fmt.Println("1) Transfer Money")
			fmt.Println("2) Momo & Pay Bill")
			fmt.Println("3) Airtime & bundles")
			fmt.Println("4) Financial Services")
			fmt.Println("5) My Wallet")
			fmt.Println("6) Allow Cash Out")

			fmt.Print("Your choice: ")
			fmt.Scan(&choice)

			fmt.Println(" ")

			if choice == 1 {
				fmt.Println("Transfer Money")
				fmt.Println("1) MoMo User")
				fmt.Println("2) Non MoMo User")
				fmt.Println("3) Send with Care")
				fmt.Println("4) Favorite")
				fmt.Println("5) Other Networks")
				fmt.Println("6) Bank Account")
				fmt.Println("7) Back")

				fmt.Print("Your choice: ")
				fmt.Scan(&firstChoice)

				fmt.Println(" ")

				if firstChoice == 1 {
					fmt.Print("Enter mobile number: ")
					fmt.Scan(&number)

					fmt.Println(" ")

					fmt.Print("Confirm number: ")
					fmt.Scan(&reEnterNumber)

					fmt.Println(" ")

					if number == reEnterNumber {
						fmt.Print("Enter Amount: ")
						fmt.Scan(&amount)
					}

					fmt.Println(" ")

					fmt.Print("Enter Reference: ")
					fmt.Scan(&ref)

					fmt.Printf(`
				  Transfer to Godfred Asare for GHS
				  %.1f with Reference: %v. Fee is GHS 0.00,
				  Tax amount is GHS 0.00, Total Amount 
				  is GHS %.1f, Enter MM PIN
				  # for next 
				`, amount, ref, amount)
					fmt.Scan(&pin)

					if pin == "0000" {
						CheckBalanceFile(amount, ref)
					}

					break

				} else if choice == 7 {
					continue
				}
				}
				 if choice == 3 {
					fmt.Print("Hello")
			}

		}
	}
}

func CheckBalanceFile(amount float64, ref string) {
	text := fmt.Sprintf(`
	Payment received for GHS %.1f
	from Godfred Asare 
	Current Balance: GHS 21.00.
	reference: %v `, amount, ref)
	os.WriteFile("Message.txt", []byte(text), 0644)
}
