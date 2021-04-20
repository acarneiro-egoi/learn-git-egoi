package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		return c.SendString(fmt.Sprintf("Hello, %s!", name))
	})

	app.Post("/hello-person", func(c *fiber.Ctx) error {
		person := new(Person)
		if err := c.BodyParser(person); err == nil {
			return c.Status(http.StatusBadRequest).SendString("invalid json")
		}

		return c.SendString(fmt.Sprintf("%s, %d", person.FullName(), person.Age))
	})

	app.Listen(":3000")
}

type Person struct {
	Name     string `json: "name"`
	LastName string `json: "name"`
	Age      int8   `json: "age"`
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.LastName)
}
