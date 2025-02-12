package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutesDoctors(app *fiber.App) {
	doctors := app.Group("/doctors")
	doctors.Get("/", GetDoctorsListFilter)
	doctors.Get("/:id", GetDoctorsListID)
	doctors.Post("/", CreateDoctor)
	doctors.Delete("/:id", DeleteDoctor)
	doctors.Put("/:id", UpdateDoctor)

	doctors.Get("/schedule", GetAppointments)
	doctors.Post("/schedule", CreateAppointment)
}
