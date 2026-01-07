package user

import (
	"time"
	"trackit/app/db"
	"trackit/app/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type BudgetController interface {
	GetBudget(c *fiber.Ctx, store *session.Store) error
	UpdateBudget(c *fiber.Ctx, store *session.Store) error
}

type budget struct{}

func NewBudget() BudgetController {
	return &budget{}
}

func (b *budget) GetBudget(c *fiber.Ctx, store *session.Store) error {
	return sendBudget(c, store, "")
}

func (b *budget) UpdateBudget(c *fiber.Ctx, store *session.Store) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "session error",
		})
	}
	username := sess.Get("username").(string)
	var req model.UpdateBudgetRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	for _, item := range req.Budgets {
		budget := model.Budget{
			Username:     username,
			Category:     item.Category,
			BudgetAmount: item.BudgetAmount,
			Month:        req.Month,
		}
		res := db.NewBudgetRepo().UpdateBudget(budget)
		if res.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update budget",
			})
		}
	}

	message := "Budget updated successfully"
	return sendBudget(c, store, message)
}

func sendBudget(c *fiber.Ctx, store *session.Store, message string) error {
	sess, _ := store.Get(c)
	username := sess.Get("username").(string)
	month := int(time.Now().Month())
	budgets := db.NewBudgetRepo().GetBudget(username, month)
	return c.Render("user/budget", fiber.Map{
		"budgets": budgets,
		"message": message,
	})
}
