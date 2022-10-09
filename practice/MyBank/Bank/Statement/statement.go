package Statement

import (
	"Setup/Bank/Currency"
	"Setup/CSV"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Statement struct {
	Id       uuid.UUID
	Value    float64
	Type     string
	CreateAt time.Time
}

var statements []Statement

func Balance() float64 {
	balance := 0.0

	for _, statement := range statements {
		if statement.Type == "Deposit" {
			balance = balance + statement.Value
		}

		if statement.Type == "Withdraw" {
			balance = balance - statement.Value
		}
	}

	return balance
}

func New(statement Statement) ([]Statement, error) {
	balance := Balance()

	if statement.Type != "Withdraw" && statement.Type != "Deposit" {
		println("Invalid Type")

		return nil, errors.New("")
	}

	if statement.Type == "Withdraw" && statement.Value > balance {
		println("Você não tem dinheiro suficiente para realizar essa retirada!")

		return nil, errors.New("")
	}

	statements = append(statements, statement)

	return statements, nil
}

func List() []Statement {
	return statements
}

func Download() {
	var rows [][]string

	for _, statement := range statements {
		rows = append(rows, FormatStatement(statement))
	}

	balance := Balance()

	rows = append(rows, []string{"Total:", Currency.FormatBalance(balance)})

	CSV.Create(CSV.CreateCSV{
		Filename: "statements",
		Rows:     rows,
		Titles:   []string{"Tipo", "Valor", "Data"},
	})
}

func FormatStatement(statement Statement) []string {
	statementType := "Depósito"

	if statement.Type == "Withdraw" {
		statementType = "Retirada"
	}

	return []string{
		statementType,
		Currency.FormatBalance(statement.Value),
		statement.CreateAt.Format("02/01/2006 15:04:05"),
	}
}
