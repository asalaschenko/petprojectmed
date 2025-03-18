package main

import (
	"context"
	"errors"
	"os"
	"petprojectmed/common"
	"petprojectmed/routes"
	"petprojectmed/storage"

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

	connD := storage.GetConnectionDB()
	defer connD.Close(context.Background())
	routes.RegisterRoutesDoctors(app, port, connD)

	connP := storage.GetConnectionDB()
	defer connP.Close(context.Background())
	routes.RegisterRoutesPatients(app, port, connP)

	connSch := storage.GetConnectionDB()
	defer connSch.Close(context.Background())
	routes.RegisterRoutesSchedule(app, port, connSch)

	app.Listen(":" + port)
}
