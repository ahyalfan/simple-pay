package api

import "github.com/gofiber/fiber/v2"

type notFound struct {
}

func NewNotFound(app *fiber.App) {
	n := notFound{}
	app.All("*", n.Handle)
}

func (n notFound) Handle(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return c.SendString("Custom 404 error: Page not found.")
}
