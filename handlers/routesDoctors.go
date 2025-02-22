package handlers

import (
	"petprojectmed/jsonFormatDB"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutesDoctors(app *fiber.App, port string) {
	if port == "3000" {
		doctors := app.Group("/doctors")
		schedule := doctors.Group("/schedule")
		schedule.Get("/", GetAppointments)
		schedule.Post("/", CreateAppointment)
		schedule.Delete("/:id", DeleteAppointment)
		doctors.Get("/", GetDoctorsListFilter)
		doctors.Get("/:id", GetDoctorsListID)
		doctors.Post("/", CreateDoctor)
		doctors.Delete("/:id", DeleteDoctor)
		doctors.Put("/:id", UpdateDoctor)
	} else {
		doctors := app.Group("/doctors")
		doctors.Get("/", jsonFormatDB.GetDoctorsListFilter)
		doctors.Get("/:id", jsonFormatDB.GetDoctorsListID)
		doctors.Post("/", jsonFormatDB.CreateDoctor)
		doctors.Delete("/:id", jsonFormatDB.DeleteDoctor)
		doctors.Put("/:id", jsonFormatDB.UpdateDoctor)

		doctors.Get("/schedule", jsonFormatDB.GetAppointments)
		doctors.Post("/schedule", jsonFormatDB.CreateAppointment)
	}
}
