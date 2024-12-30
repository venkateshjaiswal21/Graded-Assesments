package services

import (
	"a2_bank-transaction-system/models"
	"errors"
	"fmt"
	"strings"
)

var accounts []models.Account

func AddAccount(id int, name string, balance float64) error {
	if id <= 0 {
		return errors.New("iD must be a positive integer")
	}
	if len(strings.TrimSpace(name)) == 0 || !isAlpha(name) {
		return errors.New("name must be a non-empty string containing only alphabetic characters")
	}
	if balance < 0 {
		return errors.New("initial balance cannot be negative")
	}
	for _, account := range accounts {
		if account.ID == id {
			return errors.New("account ID must be unique")
		}
	}
	accounts = append(accounts, models.Account{
		ID:              id,
		Name:            name,
		Balance:         balance,
		TransactionHist: []string{"Account created with balance: " + fmt.Sprintf("%.2f", balance)},
	})
	return nil
}

func Deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	for i, account := range accounts {
		if account.ID == id {
			accounts[i].Balance += amount
			accounts[i].TransactionHist = append(accounts[i].TransactionHist, fmt.Sprintf("Deposited: %.2f Balance: %.2f", amount, accounts[i].Balance))
			return nil
		}
	}
	return errors.New("account not found")
}

func Withdraw(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	for i, account := range accounts {
		if account.ID == id {
			if account.Balance < amount {
				return errors.New("insufficient balance")
			}
			accounts[i].Balance -= amount
			accounts[i].TransactionHist = append(accounts[i].TransactionHist, fmt.Sprintf("Withdrew: %.2f Balance: %.2f", amount, accounts[i].Balance))
			return nil
		}
	}
	return errors.New("account not found")
}

func ViewTransactionHistory(id int) ([]string, error) {
	for _, account := range accounts {
		if account.ID == id {
			return account.TransactionHist, nil
		}
	}
	return nil, errors.New("account not found")
}

func isAlpha(input string) bool {
	for _, r := range input {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == ' ') {
			return false
		}
	}
	return true
}
