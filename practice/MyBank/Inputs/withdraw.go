package Inputs

import (
	"Setup/Bank"
	"Setup/Bank/Currency"
	"fmt"
)

func NewWithdraw() {
	var value string = ""

	Clear()

	fmt.Print("Digite o valor que deseja retirar: ")

	_, _ = fmt.Scan(&value)

	currency, err := FormatCurrencyToFloat(value)

	if err != nil {
		TryAgain(NewWithdraw)

		return
	}

	_, withdrawError := Bank.Withdraw(currency)

	if withdrawError != nil {
		TryAgain(NewWithdraw)

		return
	}

	fmt.Printf("O valor de %s foi retirado com sucesso", Currency.FormatBalance(currency))
}
