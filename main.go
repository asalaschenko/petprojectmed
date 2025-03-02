package main

import (
	"errors"
	"os"
	"petprojectmed/common"
	"petprojectmed/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port, exists := os.LookupEnv("PORT_GOLANG")
	if !exists {
		common.CheckErr(errors.New("not found port for current application"))
	}

	app := fiber.New(fiber.Config{
		Prefork:                  true,
		EnableSplittingOnParsers: true})

	app.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	registerRoutes(app, port)

	app.Listen(":" + port)
}

func registerRoutes(app *fiber.App, port string) {
	routes.RegisterRoutesDoctors(app, port)
}
