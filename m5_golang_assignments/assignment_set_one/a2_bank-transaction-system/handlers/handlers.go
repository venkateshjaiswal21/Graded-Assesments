package handlers

import (
	"a2_bank-transaction-system/services"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddAccountHandler() {
	reader := bufio.NewReader(os.Stdin)
	var idStr, name, balanceStr string
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: ID must be a positive integer")
		return
	}
	fmt.Print("Enter Account Holder Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Enter Initial Balance: ")
	balanceStr, _ = reader.ReadString('\n')
	balanceStr = strings.TrimSpace(balanceStr)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		fmt.Println("Error: Balance must be a valid number")
		return
	}

	if err := services.AddAccount(id, name, balance); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Account added successfully.")
	}
}

func DepositHandler() {
	reader := bufio.NewReader(os.Stdin)
	var idStr, amountStr string
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: ID must be a positive integer")
		return
	}
	fmt.Print("Enter Deposit Amount: ")
	amountStr, _ = reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Error: Amount must be a valid number")
		return
	}

	if err := services.Deposit(id, amount); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Deposit successful.")
	}
}

func WithdrawHandler() {
	reader := bufio.NewReader(os.Stdin)
	var idStr, amountStr string
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: ID must be a positive integer")
		return
	}
	fmt.Print("Enter Withdrawal Amount: ")
	amountStr, _ = reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Error: Amount must be a valid number")
		return
	}

	if err := services.Withdraw(id, amount); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Withdrawal successful.")
	}
}

func ViewTransactionHistoryHandler() {
	// reader := bufio.NewReader(os.Stdin)
	var idStr string
	fmt.Print("Enter Account ID: ")
	fmt.Scanln(&idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error: ID must be a positive integer")
		return
	}

	history, err := services.ViewTransactionHistory(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Transaction History:")
		for _, record := range history {
			fmt.Println(record)
		}
	}
}
