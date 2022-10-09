package Inputs

import (
	"fmt"
	"strings"
)

func ValidateTryAgain(value string) bool {
	if value != "sim" && value != "não" {
		return false
	}

	return true
}

func TryAgain(callback func()) {
	var value string = ""

	fmt.Print("Deseja tentar novamente? [Sim/Não]")

	_, _ = fmt.Scan(&value)

	value = strings.ToLower(value)

	isValid := ValidateTryAgain(value)

	if !isValid {
		fmt.Println("Operação não reconhecida pelo nosso sistema")

		TryAgain(callback)

		return
	}

	if value == "sim" {
		callback()
	}
}
