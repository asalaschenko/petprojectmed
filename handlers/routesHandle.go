package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	doctors := app.Group("/doctors")
	doctors.Get("/list", GetDoctorsList)
	doctors.Post("/new", CreateDoctor)
	doctors.Delete("/delete/:id", DeleteDoctor)
	doctors.Put("/update/:id", UpdateDoctor)

	patients := app.Group("/patients")
	patients.Get("/list", GetPatientsList)
	patients.Post("/new", CreatePatient)
	patients.Delete("/delete/:id", DeletePatient)
	patients.Put("/update/:id", UpdatePatient)

}
