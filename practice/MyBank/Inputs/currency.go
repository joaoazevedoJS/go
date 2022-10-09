package Inputs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func FormatCurrencyToFloat(value string) (float64, error) {
	currency, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if err != nil || currency == 0.0 {
		fmt.Println("Valor Inválido!")

		return 0.0, errors.New("Valor Inválido!")
	}

	return currency, nil
}
