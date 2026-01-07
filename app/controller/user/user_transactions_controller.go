package user

import (
	"strconv"
	"trackit/app/model"

	"trackit/app/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Transaction interface {
	AddExpense(c *fiber.Ctx, store *session.Store) error
	AddIncome(c *fiber.Ctx, store *session.Store) error
}

type transaction struct{}

func NewTransaction() Transaction {
	return &transaction{}
}

func (t *transaction) AddExpense(c *fiber.Ctx, store *session.Store) error {
	description := c.FormValue("description")
	category := c.FormValue("category")
	amount := c.FormValue("amount")

	sess, _ := store.Get(c)
	username := sess.Get("username").(string)

	amt, err := strconv.Atoi(amount)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid amount",
		})
	}

	transaction := model.Transaction{
		Username:    username,
		Description: description,
		Category:    category,
		Type:        "Expense",
		Amount:      amt,
	}

	if err := t.AddTransaction(transaction); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to add expense",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Expense added successfully!",
	})
}

func (t *transaction) AddIncome(c *fiber.Ctx, store *session.Store) error {
	description := c.FormValue("description")
	category := c.FormValue("source")
	amount := c.FormValue("amount")

	sess, _ := store.Get(c)
	username := sess.Get("username").(string)

	amt, err := strconv.Atoi(amount)
	if err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid amount",
		})
	}

	transaction := model.Transaction{
		Username:    username,
		Description: description,
		Category:    category,
		Type:        "Income",
		Amount:      amt,
	}

	if err := t.AddTransaction(transaction); err != nil {
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to add income",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Income added successfully!",
	})
}

func (t *transaction) AddTransaction(transaction model.Transaction) error {
	res := db.NewTransaction().AddTransaction(transaction)
	return res.Error
}
