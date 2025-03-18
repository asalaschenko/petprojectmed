package schedule

import (
	"petprojectmed/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ControllerGetAppointment struct {
	filterService IFilterService
}

func NewControllerGetAppointment(filterService IFilterService) *ControllerGetAppointment {
	value := new(ControllerGetAppointment)
	value.filterService = filterService
	return value
}

func (cg *ControllerGetAppointment) GetAppointment(c *fiber.Ctx) error {
	queryFilters := new(QuerySheduleListFilter)
	err := c.QueryParser(queryFilters)

	if err == nil {
		outputData := cg.filterService.GetList(queryFilters)
		status := cg.filterService.ReturnStatus()
		if status == common.OK {
			return c.JSON(outputData)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString(status)
		}
	} else {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_REQUEST)
	}
}

type ControllerCreateAppointment struct {
	createService ICreateService
}

func NewControllerCreateAppointment(createService ICreateService) *ControllerCreateAppointment {
	value := new(ControllerCreateAppointment)
	value.createService = createService
	return value
}

func (cg *ControllerCreateAppointment) CreateAppointment(c *fiber.Ctx) error {
	appointmentJson := new(Appointment)

	if err := c.BodyParser(appointmentJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
	}
	if description, err := appointmentJson.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(description)
	}

	cg.createService.Create(appointmentJson)
	status := cg.createService.ReturnStatus()
	return c.Status(fiber.StatusCreated).SendString(status)
}

type ControllerDeleteAppointment struct {
	deleteService IDeleteService
}

func NewControllerDeleteAppointment(deleteService IDeleteService) *ControllerDeleteAppointment {
	value := new(ControllerDeleteAppointment)
	value.deleteService = deleteService
	return value
}

func (cd *ControllerDeleteAppointment) DeleteAppointment(c *fiber.Ctx) error {
	ID := c.Params("id")

	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_ID_REQUEST)
	}

	if verify(&intID) {
		appointment := cd.deleteService.Delete(&intID)
		status := cd.deleteService.ReturnStatus()
		if status == common.OK {
			return c.JSON(appointment)
		} else {
			return c.Status(fiber.StatusForbidden).SendString(status)
		}
	} else {
		return c.Status(fiber.StatusBadRequest).SendString(common.NOT_FOUND_ID)
	}
}
