package routes

import (
	"petprojectmed/doctors"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutesDoctors(app *fiber.App, port string) {
	if port == "3000" {
		doctor := app.Group("/doctor")
		doctor.Get("/", func(c *fiber.Ctx) error {
			return doctors.ControllerGetDoctors(c, doctors.FILTER)
		})
		doctor.Get("/:id", func(c *fiber.Ctx) error {
			return doctors.ControllerGetDoctors(c, doctors.LIST_ID)
		})
		doctor.Post("/", doctors.ControllerCreateDoctor)
		//doctors.Delete("/:id", )
		doctor.Put("/:id", doctors.ControllerUpdateDoctor)
	}
}
