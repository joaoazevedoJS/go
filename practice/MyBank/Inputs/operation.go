package Inputs

import (
	"Setup/Bank"
	"fmt"
)

func ValidOperation(value string) {
	if value != "1" && value != "2" && value != "3" && value != "4" && value != "5" {
		fmt.Println("Operação não reconhecida pelo nosso sistema")

		NewOperation()

		return
	}
}

func NewOperation() {
	var operation string

	Clear()

	fmt.Println("Deseja fazer qual tipo de operação?")
	fmt.Println("[1]: Depósito")
	fmt.Println("[2]: Retirada")
	fmt.Println("[3]: Ver Extrato")
	fmt.Println("[4]: Download Extrato")
	fmt.Println("[5]: Sair")
	fmt.Print("Digite o número da operação: ")

	_, _ = fmt.Scan(&operation)

	ValidOperation(operation)

	if operation == "1" {
		NewDeposit()
	}

	if operation == "2" {
		NewWithdraw()
	}

	if operation == "3" {
		ShowStatements()
	}

	if operation == "4" {
		Bank.DownloadStatement()
	}

	if operation == "5" {
		return
	}

	NewOperation()
}
