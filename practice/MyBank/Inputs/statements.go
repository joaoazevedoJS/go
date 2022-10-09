package Inputs

import (
	"Setup/Bank"
	"Setup/Bank/Currency"
	"Setup/Bank/Statement"
	"errors"
	"fmt"
	"strings"
)

func ValidateYesOrNo(value string) bool {
	if value != "sim" && value != "não" {
		return false
	}

	return true
}

func YesOrNo() (bool, error) {
	var value string = ""

	_, _ = fmt.Scan(&value)

	value = strings.ToLower(value)

	isValid := ValidateYesOrNo(value)

	if !isValid {
		fmt.Println("Operação não reconhecida pelo nosso sistema")

		return false, errors.New("Invalid input")
	}

	if value == "não" {
		return false, nil
	}

	return true, nil
}

func ShowMoreStatements() bool {
	fmt.Println("\nMostrar mais? [Sim/Não]")

	value, err := YesOrNo()

	if err != nil {
		return ShowMoreStatements()
	}

	return value
}

func CloseStatements() bool {
	fmt.Println("\nFechar extrato? [Sim/Não]")

	yes, err := YesOrNo()

	if err != nil || !yes {
		return CloseStatements()
	}

	return true
}

func ShowStatements() {
	limit := 2
	statements := Bank.Statements()

	fmt.Println("Tipo      Valor     Data")

	lengthTotal := len(statements)

	for index, statement := range statements {
		statementFormatted := Statement.FormatStatement(statement)

		fmt.Printf("%s   %s   %s\n", statementFormatted[0], statementFormatted[1], statementFormatted[2])

		if length := index + 1; length >= limit && length < lengthTotal {
			stop := ShowMoreStatements()

			if !stop {
				break
			}

			fmt.Println("\nTipo      Valor     Data")

			limit = limit + 2
		}
	}

	fmt.Printf("\nTotal: %s\n", Currency.FormatBalance(Statement.Balance()))

	CloseStatements()
}
