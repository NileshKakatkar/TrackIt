package user

import (
	"fmt"
	"time"
	"trackit/app/db"
	"trackit/app/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type UserController interface {
	Login(c *fiber.Ctx, store *session.Store) error
	Register(c *fiber.Ctx, store *session.Store) error
	Dashboard(c *fiber.Ctx, store *session.Store) error
	Transactions(c *fiber.Ctx, store *session.Store) error
}
type userController struct{}

type StatusData struct {
	Status    string
	Msg       string
	MsgType   string
	OrgErrMsg string
}

func NewUserController() UserController {
	return &userController{}
}

func (u *userController) Login(c *fiber.Ctx, store *session.Store) error {
	username := c.FormValue("username")
	userPwd := c.FormValue("password")

	user, _ := db.NewUser().GetUser(username)
	if userPwd != user.Password {
		return c.Render("user/login", fiber.Map{
			"statusMsg": StatusData{Msg: "Invalid credentials", MsgType: "alert-danger"},
		})
	}
	sess, err := store.Get(c)
	if err != nil {
		fmt.Printf("Error in session creation: %v", err)
	}
	sess.Set("username", username)
	sess.Set("isLoggedIn", true)
	if err := sess.Save(); err != nil {
		return err
	}
	return c.Redirect("/user/dashboard", fiber.StatusSeeOther)
}

func (u *userController) Register(c *fiber.Ctx, store *session.Store) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		fmt.Printf("ERROR: %v", err)
		_ = fmt.Errorf("error while reading the data: %v", err)
		return c.Render("user/register", fiber.Map{
			"message": "Invalid data",
		})
	}
	message, exist := db.NewUser().IsUserExist(user)
	if exist {
		fmt.Println("User exist")
		return c.Render("user/register", fiber.Map{
			"statusMsg": StatusData{Msg: message, MsgType: "alert-danger"},
		})
	}
	db.NewUser().Register(user)
	fmt.Println("Rendering to dashboard")
	// return c.Render("user/dashboard", fiber.Map{
	// 	"message": "registration successful",
	// })
	sess, err := store.Get(c)
	if err != nil {
		fmt.Printf("Error in session creation: %v", err)
	}
	sess.Set("username", user.Username)
	sess.Set("isLoggedIn", true)
	if err := sess.Save(); err != nil {
		return err
	}
	return c.Redirect("/user/dashboard", fiber.StatusSeeOther)
}

func (u *userController) Dashboard(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	userId := sess.Get("username").(string)
	// userId := "user01"
	// // bills, err := db.NewBillingRepo().GetBilling(userId)
	// if err != nil {
	// 	fmt.Println("Error while extracting the data")
	// 	fmt.Printf("%v", err)
	// }
	transactions, err := db.NewTransaction().RecentTransactions(userId)
	if err != nil {
		fmt.Println("Error in repo")
		fmt.Printf("%v", err)
	}

	Categories := GetCategories()
	Sources := GetSources()
	income := GetIncome(userId)
	budget := db.NewBudgetRepo().GetBudget(userId, int(time.Now().Month()))
	return c.Render("user/dashboard", fiber.Map{
		"ActivePage": "dashboard",
		// "bills":        bills,
		"transactions": transactions,
		"categories":   Categories,
		"sources":      Sources,
		"income":       income,
		"budget":       budget,
	})
}

func GetCategories() []string {
	return []string{
		"food",
		"transport",
		"entertainment",
		"health",
		"education",
		"rent",
		"others",
	}
}

func GetSources() []string {
	return []string{
		"Salary",
		"Business",
		"Freelance",
		"Interest",
		"Rental",
		"Bonus",
		"Gift",
		"Refund",
		"Other",
	}
}

func GetIncome(username string) model.UserSummary {
	monthIncome := db.NewBillingRepo().GetMonthlyIncome(username)
	monthExpense := db.NewBillingRepo().GetMonthlyExpense(username)
	// totalIncome := db.NewBillingRepo().GetTotalIncome(username)
	balance := db.NewBillingRepo().GetUserBalance(username)
	fmt.Printf("balance: %v\n", balance)
	return model.UserSummary{
		MonthlyIncome:  monthIncome,
		MonthlyExpense: monthExpense,
		TotalBalance:   balance.TotalBalance,
	}
}

func (u *userController) Transactions(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	userId := sess.Get("username").(string)
	// userId = "user01"
	transactions, err := db.NewTransaction().GetAllTransaction(userId)
	if err != nil {
		fmt.Println("Error in repo")
		fmt.Printf("%v", err)
	}
	return c.Render("user/transactions", fiber.Map{
		"ActivePage":   "transactions",
		"transactions": transactions,
	})
}
