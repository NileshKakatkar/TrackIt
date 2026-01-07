package service

import "trackit/app/db"

type Income interface {
	UpdateIncome(income int, username string)
}

type income struct{}

func NewIncome() Income {
	return &income{}
}

func (i *income) UpdateIncome(income int, username string) {
	bill, _ := db.NewBillingRepo().GetBilling(username)
	bill.TotalBalance += income
	db.NewBillingRepo().UpdateIncome(bill)
}
