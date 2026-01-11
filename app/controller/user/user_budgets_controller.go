package user

import (
	"fmt"
	"time"
	"trackit/app/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type BudgetController interface {
	GetBudget(c *fiber.Ctx, store *session.Store) error
}

type budget struct{}

func NewBudget() BudgetController {
	return &budget{}
}

func (b *budget) GetBudget(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	username := sess.Get("username").(string)
	month := int(time.Now().Month())
	budgets := db.NewBudgetRepo().GetBudget(username, month)
	fmt.Println("Total budget")
	fmt.Printf("%v", budgets)
	message := ""
	return c.Render("user/budget", fiber.Map{
		"budgets": budgets,
		"message": message,
	})
}
