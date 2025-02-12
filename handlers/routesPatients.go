package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutesPatients(app *fiber.App) {

	patients := app.Group("/patients")
	patients.Get("/", GetPatientsListFilter)
	patients.Get("/:id", GetPatientsListID)
	patients.Post("/", CreatePatient)
	patients.Delete("/:id", DeletePatient)
	patients.Put("/:id", UpdatePatient)
}
