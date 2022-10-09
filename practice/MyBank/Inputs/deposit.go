package Inputs

import (
	"Setup/Bank"
	"Setup/Bank/Currency"
	"fmt"
)

func NewDeposit() {
	var value string = ""

	Clear()

	fmt.Print("Digite o valor do depósito: ")

	_, _ = fmt.Scan(&value)

	currency, err := FormatCurrencyToFloat(value)

	if err != nil {
		TryAgain(NewDeposit)

		return
	}

	_, depositError := Bank.Deposit(currency)

	if depositError != nil {
		TryAgain(NewDeposit)

		return
	}

	fmt.Printf("O valor de %s foi depositado com sucesso", Currency.FormatBalance(currency))
}
