package main

import (
	"a2_bank-transaction-system/handlers"
	"fmt"
)

func main() {
	for {
		fmt.Println("\nBank Transaction System")
		fmt.Println("1. Add Account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. View Transaction History")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handlers.AddAccountHandler()
		case 2:
			handlers.DepositHandler()
		case 3:
			handlers.WithdrawHandler()
		case 4:
			handlers.ViewTransactionHistoryHandler()
		case 5:
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
