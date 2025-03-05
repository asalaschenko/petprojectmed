package schedule

import (
	"petprojectmed/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ControllerGetAppointment(c *fiber.Ctx) error {
	queryFilters := new(QuerySheduleListFilter)
	err := c.QueryParser(queryFilters)

	if err == nil {
		outputData, status := queryFilters.GetList()
		if status == common.OK {
			return c.JSON(outputData)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString(status)
		}
	} else {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_REQUEST)
	}
}

func ControllerCreateAppointment(c *fiber.Ctx) error {
	appointmentJson := new(Appointment)

	if err := c.BodyParser(appointmentJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
	}
	if description, err := appointmentJson.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(description)
	}

	status := appointmentJson.Create()
	return c.Status(fiber.StatusCreated).SendString(status)
}

func ControllerDeleteAppointment(c *fiber.Ctx) error {
	ID := c.Params("id")

	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_ID_REQUEST)
	}

	AppointmentID := appointmentID(intID)
	if AppointmentID.verify() {
		status, appointment := AppointmentID.Delete()
		if status == common.OK {
			return c.JSON(appointment)
		} else {
			return c.Status(fiber.StatusForbidden).SendString(status)
		}
	} else {
		return c.Status(fiber.StatusBadRequest).SendString(common.NOT_FOUND_ID)
	}
}
