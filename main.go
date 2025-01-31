package main

import (
	"errors"
	"os"
	"petprojectmed/handlers"
	"petprojectmed/utils"

	"github.com/gofiber/fiber/v2"
)

var err_env error = errors.New("not found port for current application")

func main() {
	port, exists := os.LookupEnv("PORT_GOLANG")
	if !exists {
		utils.CheckErr(err_env)
	}

	app := fiber.New(fiber.Config{
		Prefork:                  true,
		EnableSplittingOnParsers: true})

	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	handlers.RegisterRoutes(app)

	app.Listen(":" + port)
}
