package db

import (
	"log"
	"trackit/app/db/conn"
	"trackit/app/model"
)

type BudgetRepo interface {
	AddBudget(budget model.Budget) conn.ExecResult
	UpdateBudget(budget model.Budget) conn.ExecResult
	GetBudget(username string, month int) []model.Budget
}

type budget struct{}

func NewBudgetRepo() BudgetRepo {
	return &budget{}
}

func (b *budget) AddBudget(budget model.Budget) conn.ExecResult {
	execRes := conn.NewSqlDb().ExecSql(
		"INSERT INTO budgets (username, category, budget_amount, month) VALUES (?,?,?,?)",
		budget.Username,
		budget.Category,
		budget.BudgetAmount,
		budget.Month,
	)
	if execRes.Error != nil {
		log.Printf("error in query: %v", execRes.Error)
	}
	return execRes
}

func (b *budget) UpdateBudget(budget model.Budget) conn.ExecResult {
	execRes := conn.NewSqlDb().ExecSql(
		"UPDATE budgets SET budget_amount=? WHERE username=? AND category=? AND month=?",
		budget.BudgetAmount,
		budget.Username,
		budget.Category,
		budget.Month,
	)
	if execRes.Error != nil {
		log.Printf("error in query: %v", execRes.Error)
	}
	return execRes
}
func (b *budget) GetBudget(username string, month int) []model.Budget {
	qryRow := conn.NewSqlDb().QuerySql(
		`SELECT 
			b.category,
			b.budget_amount,
			COALESCE(SUM(t.amount), 0) AS spent
		FROM budgets b
		LEFT JOIN transactions t
			ON b.username = t.username
			AND b.category = t.category
			AND MONTH(t.trans_date) = ?
		WHERE b.username = ?
		AND b.month = ?
		GROUP BY b.category, b.budget_amount`,
		month, username, month,
	)

	var budgets []model.Budget
	if qryRow.Error == nil {
		rows := qryRow.Rows
		if rows != nil {
			for rows.Next() {
				var budget model.Budget

				// ✅ scan order matches SELECT
				rows.Scan(
					&budget.Category,
					&budget.BudgetAmount,
					&budget.Spent,
				)

				budget.Username = username
				budget.Month = month
				budget.Status = calculateStatus(budget.BudgetAmount, budget.Spent)

				budgets = append(budgets, budget)
			}
		}
	}
	return budgets
}

func calculateStatus(budget, spent int) string {
	if spent >= budget {
		return "OVER"
	}
	if spent >= int(float64(budget)*0.8) {
		return "WARNING"
	}
	return "SAFE"
}
