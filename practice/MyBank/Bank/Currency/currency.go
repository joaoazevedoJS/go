package Currency

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func FormatBalance(balance float64) string {
	parse := currency.MustParseISO("BRL")
	scale, _ := currency.Cash.Rounding(parse)
	decimal := number.Decimal(balance, number.Scale(scale))
	printing := message.NewPrinter(language.BrazilianPortuguese)

	return printing.Sprintf("%v %v", currency.Symbol(parse), decimal)
}
