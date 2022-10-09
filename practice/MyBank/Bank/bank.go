package Bank

import (
	"Setup/Bank/Statement"
	"github.com/google/uuid"
	"time"
)

func Deposit(value float64) ([]Statement.Statement, error) {
	return Statement.New(Statement.Statement{
		Id:       uuid.New(),
		Value:    value,
		Type:     "Deposit",
		CreateAt: time.Now(),
	})
}

func Withdraw(value float64) ([]Statement.Statement, error) {
	return Statement.New(Statement.Statement{
		Id:       uuid.New(),
		Value:    value,
		Type:     "Withdraw",
		CreateAt: time.Now(),
	})
}

func Statements() []Statement.Statement {
	return Statement.List()
}

func DownloadStatement() {
	Statement.Download()
}
