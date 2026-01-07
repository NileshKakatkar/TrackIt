package db

import (
	"fmt"
	"trackit/app/db/conn"
	"trackit/app/model"
)

type Billing interface {
	AddBillingInfo(billing model.Billing) conn.ExecResult
	UpdateIncome(billing model.Billing) conn.ExecResult
	GetBilling(username string) (model.Billing, error)
	GetMonthlyIncome(username string) int
	GetTotalIncome(username string) int
	GetMonthlyExpense(username string) int
	GetUserBalance(username string) model.UserBalance
}

type billing struct{}

func NewBillingRepo() Billing {
	return &billing{}
}

func (b *billing) AddBillingInfo(billing model.Billing) conn.ExecResult {
	execRes := conn.NewSqlDb().ExecSql("insert into billing_info (username, total_balance, month_income, expenses, upcoming_bills) values (?,?,?,?,?)", billing.Username, billing.TotalBalance, billing.MonthIncome, billing.Expenses, billing.UpcomingBills)
	if execRes.Error != nil {
		fmt.Printf("error in the query: %v\n", execRes.Error)
	}
	return execRes
}

func (b *billing) UpdateIncome(billing model.Billing) conn.ExecResult {
	execRes := conn.NewSqlDb().ExecSql("update billing_info set total_balance=?, month_income=?, expenses=?, upcoming_bills=? where username = ?", billing.TotalBalance, billing.MonthIncome, billing.Expenses, billing.UpcomingBills, billing.Username)

	if execRes.Error != nil {
		fmt.Printf("Error: %v\n", execRes.Error)
	}
	return execRes
}

func (b *billing) GetBilling(username string) (model.Billing, error) {
	queryRes := conn.NewSqlDb().QuerySql("select username, total_balance, month_income, expenses, upcoming_bills from billing_info where username=?", username)
	var bill model.Billing
	var err error

	if queryRes.Error != nil {
		fmt.Printf("Error in db: %v\n", queryRes.Error)
		err = fmt.Errorf("db operation failed")
	} else {
		rows := queryRes.Rows
		if rows != nil {
			for rows.Next() {
				rows.Scan(&bill.Username, &bill.TotalBalance, &bill.MonthIncome, &bill.Expenses, &bill.UpcomingBills)
			}
			rows.Close()
		}
	}
	return bill, err
}

func (b *billing) GetMonthlyIncome(username string) int {
	qryRes := conn.NewSqlDb().QuerySql(
		"SELECT IFNULL(SUM(amount), 0) AS monthly_income FROM transactions WHERE username=? AND expense_type=? AND trans_date >= DATE_FORMAT(CURRENT_DATE, '%Y-%m-01') AND trans_date < DATE_ADD(DATE_FORMAT(CURRENT_DATE, '%Y-%m-01'), INTERVAL 1 MONTH)",
		username,
		"Income",
	)
	var income int
	if qryRes.Error != nil {
		fmt.Printf("Error in the query: %v\n", qryRes.Error)
	} else {
		rows := qryRes.Rows
		if rows != nil {
			for rows.Next() {
				rows.Scan(&income)
			}
		}
	}
	return income
}

func (b *billing) GetMonthlyExpense(username string) int {
	qryRes := conn.NewSqlDb().QuerySql(
		"SELECT IFNULL(SUM(amount), 0) AS monthly_expense FROM transactions WHERE username=? AND expense_type=? AND trans_date >= DATE_FORMAT(CURRENT_DATE, '%Y-%m-01') AND trans_date < DATE_ADD(DATE_FORMAT(CURRENT_DATE, '%Y-%m-01'), INTERVAL 1 MONTH)",
		username,
		"Expense",
	)
	var expense int
	if qryRes.Error != nil {
		fmt.Printf("Error in the query: %v", qryRes.Error)
	} else {
		rows := qryRes.Rows
		if rows != nil {
			for rows.Next() {
				rows.Scan(&expense)
			}
		}
	}
	return expense
}

func (b *billing) GetTotalIncome(username string) int {
	qryRes := conn.NewSqlDb().QuerySql(`SELECT IFNULL(SUM(amount), 0) AS total_income FROM transactions  WHERE username=? AND expense_type=? `, username, "Income")
	var income int

	if qryRes.Error != nil {
		fmt.Printf("Error in the total income query: %v", qryRes.Error)
	} else {
		rows := qryRes.Rows
		if rows != nil {
			for rows.Next() {
				rows.Scan(&income)
			}
			rows.Close()
		}
	}

	return income
}

func (b *billing) GetUserBalance(username string) model.UserBalance {
	qryRes := conn.NewSqlDb().QuerySql(`
        SELECT
            IFNULL(SUM(CASE WHEN expense_type='Income' THEN amount END), 0) AS total_income,
            IFNULL(SUM(CASE WHEN expense_type='Expense' THEN amount END), 0) AS total_expense,
            IFNULL(SUM(CASE WHEN expense_type='Income' THEN amount END), 0) -
            IFNULL(SUM(CASE WHEN expense_type='Expense' THEN amount END), 0) AS total_balance
        FROM transactions
        WHERE username = ?
    `, username)

	var balance model.UserBalance

	if qryRes.Error != nil {
		fmt.Printf("Error in the balance query: %v\n", qryRes.Error)
		return balance
	}

	rows := qryRes.Rows
	if rows != nil {
		defer rows.Close()
		if rows.Next() {
			rows.Scan(&balance.TotalIncome, &balance.TotalExpense, &balance.TotalBalance)
		}
	}

	return balance
}
