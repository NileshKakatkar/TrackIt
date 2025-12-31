package model

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Age      string `josn:"age" form:"age"`
}

type Transaction struct {
	Username        string `json:"username"`
	Description     string `json:"desc"`
	Category        string `json:"category"`
	Type            string `json:"type"`
	Amount          int    `json:"amount"`
	TransactionDate string `json:"trans_date"`
}
type Billing struct {
	Username      string `json:"username"`
	TotalBalance  int    `json:"total_balance"`
	MonthIncome   int    `json:"month_income"`
	Expenses      int    `json:"expenses"`
	UpcomingBills int    `json:"upcoming_bills"`
}
type UserSummary struct {
	MonthlyIncome  int
	MonthlyExpense int
	TotalBalance   int
}

type UserBalance struct {
	TotalIncome  int
	TotalExpense int
	TotalBalance int
}
