package user

import (
	"trackit/app/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Settings interface {
	UpdateSettings(c *fiber.Ctx, store *session.Store) error
	GetSettings(c *fiber.Ctx, store *session.Store) error
}
type settings struct{}

func NewSettings() Settings {
	return &settings{}
}

func (s *settings) GetSettings(c *fiber.Ctx, store *session.Store) error {
	sess, _ := store.Get(c)
	username := sess.Get("username").(string)
	return sendSettings(c, username, "")
}
func (s *settings) UpdateSettings(c *fiber.Ctx, store *session.Store) error {
	return sendSettings(c, "", "")
}
func sendSettings(c *fiber.Ctx, username, message string) error {
	user, _ := db.NewUser().GetUser(username)
	return c.Render("user/settings", fiber.Map{
		"user":user,
	})
}
