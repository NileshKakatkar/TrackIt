package db

import (
	"fmt"
	"project/app/db/conn"
	"project/app/model"
	"time"
)

type Transaction interface {
	AddTransaction(transaction model.Transaction) conn.ExecResult
	GetAllTransaction(username string) ([]model.Transaction, error)
	RecentTransactions(username string) ([]model.Transaction, error)
}

type transaction struct{}

func NewTransaction() Transaction {
	return &transaction{}
}

func (t *transaction) AddTransaction(transaction model.Transaction) conn.ExecResult {
	currTime := time.Now()
	execRes := conn.NewSqlDb().ExecSql("insert into transactions (username, description, category, expense_type, amount, trans_date) values ($1, $2, $3, $4, $5, $6)", transaction.Username, transaction.Description, transaction.Category, transaction.Type, transaction.Amount, currTime)

	if execRes.Error != nil {
		fmt.Printf("Error in query: %v", execRes.Error)
	}
	return execRes
}

func (t *transaction) GetAllTransaction(username string) ([]model.Transaction, error) {
	qryRes := conn.NewSqlDb().QuerySql("select username,description, category, expense_type, amount, trans_date from transactions where username= $1 order by trans_date DESC", username)
	var transactions []model.Transaction
	var err error
	if qryRes.Error != nil {
		fmt.Printf("Error: %v", qryRes.Error)
		err = fmt.Errorf("Failed to get data")
	} else {
		rows := qryRes.Rows
		if rows != nil {
			for rows.Next() {
				var tran model.Transaction
				var date time.Time
				rows.Scan(&tran.Username, &tran.Description, &tran.Category, &tran.Type, &tran.Amount, &date)
				tran.TransactionDate = date.Format("2006-01-02 15:04:05")
				transactions = append(transactions, tran)
			}
			rows.Close()
		}
	}

	return transactions, err
}

func (t *transaction) RecentTransactions(username string) ([]model.Transaction, error) {
	qryRes := conn.NewSqlDb().QuerySql("select username,description, category, expense_type, amount, trans_date from transactions where username= $1 order by trans_date DESC limit 5", username)
	var transactions []model.Transaction
	var err error
	if qryRes.Error != nil {
		fmt.Printf("Error: %v", qryRes.Error)
		err = fmt.Errorf("Failed to get data")
	} else {
		rows := qryRes.Rows
		if rows != nil {
			for rows.Next() {
				var tran model.Transaction
				var date time.Time
				rows.Scan(&tran.Username, &tran.Description, &tran.Category, &tran.Type, &tran.Amount, &date)
				tran.TransactionDate = date.Format("2006-01-02 15:04:05")
				transactions = append(transactions, tran)
			}
			rows.Close()
		}
	}

	return transactions, err
}
