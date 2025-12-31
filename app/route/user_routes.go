package route

import (
	"fmt"
	userController "project/app/controller/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store = session.New(session.Config{
	Expiration: 30 * time.Minute,
})

func SetupUserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Use("/js/*", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/javascript;charset=UTF-8")
		return c.SendFile("./public/user/js/" + c.AllParams()["*2"])
	})
	user.Use("/images/*", func(c *fiber.Ctx) error {
		return c.SendFile("./public/user/images/" + c.AllParams()["*2"])
	})
	user.Static("/css", "./public/user/css")

	user.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("user/login", nil)
	})

	user.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("user/register", nil)
	})

	user.Post("/login", func(c *fiber.Ctx) error {
		return userController.NewUserController().Login(c, store)
	})
	user.Post("/register", func(c *fiber.Ctx) error {
		return userController.NewUserController().Register(c, store)
	})

	user.Get("/dashboard", SessionValidator(store), func(c *fiber.Ctx) error {
		return userController.NewUserController().Dashboard(c, store)
	})
	user.Get("/transactions", SessionValidator(store), func(c *fiber.Ctx) error {
		return userController.NewUserController().Transactions(c, store)
	})
	user.Post("/expense/create", func(c *fiber.Ctx) error {
		return userController.NewTransaction().AddExpense(c, store)
	})
	user.Post("/income/create", func(c *fiber.Ctx) error {
		return userController.NewTransaction().AddIncome(c, store)
	})
	user.Post("/budget/create", func(c *fiber.Ctx) error {
		return userController.NewTransaction().AddIncome(c, store)
	})
}

func SessionValidator(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil || sess.Get("isLoggedIn") == nil {
			fmt.Println("Session expired redirecting to login page")
			return c.Render("user/login", userController.StatusData{Msg: "session expired", MsgType: "alert-danger"})
		}

		sess.Save()
		return c.Next()
	}
}
