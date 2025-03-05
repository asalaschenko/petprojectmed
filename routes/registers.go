package routes

import (
	"petprojectmed/common"
	"petprojectmed/doctors"
	"petprojectmed/patients"
	"petprojectmed/schedule"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutesDoctors(app *fiber.App, port string) {
	if port == "3000" {
		doctor := app.Group("/doctor")
		doctor.Get("/", func(c *fiber.Ctx) error {
			return doctors.ControllerGetDoctors(c, common.FILTER)
		})
		doctor.Get("/:id", func(c *fiber.Ctx) error {
			return doctors.ControllerGetDoctors(c, common.LIST_ID)
		})
		doctor.Post("/", doctors.ControllerCreateDoctor)
		doctor.Delete("/:id", func(c *fiber.Ctx) error {
			return doctors.ControllerUpdateDeleteDoctor(c, common.DELETE)
		})
		doctor.Put("/:id", func(c *fiber.Ctx) error {
			return doctors.ControllerUpdateDeleteDoctor(c, common.UPDATE)
		})
	}
}

func RegisterRoutesPatients(app *fiber.App, port string) {
	if port == "3000" {
		patient := app.Group("/patient")
		patient.Get("/", func(c *fiber.Ctx) error {
			return patients.ControllerGetPatients(c, common.FILTER)
		})
		patient.Get("/:id", func(c *fiber.Ctx) error {
			return patients.ControllerGetPatients(c, common.LIST_ID)
		})
		patient.Post("/", patients.ControllerCreatePatient)
		patient.Delete("/:id", func(c *fiber.Ctx) error {
			return patients.ControllerUpdateDeletePatient(c, common.DELETE)
		})
		patient.Put("/:id", func(c *fiber.Ctx) error {
			return patients.ControllerUpdateDeletePatient(c, common.UPDATE)
		})
	}
}

func RegisterRoutesSchedule(app *fiber.App, port string) {
	if port == "3000" {
		Schedule := app.Group("/schedule")
		Schedule.Get("/", schedule.ControllerGetAppointment)
		Schedule.Post("/", schedule.ControllerCreateAppointment)
		Schedule.Delete("/:id", schedule.ControllerDeleteAppointment)
	}
}
