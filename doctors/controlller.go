package doctors

import (
	"errors"
	"petprojectmed/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ControllerGetDoctors struct {
	filterService IFilterService
	listService   IListService
}

func NewControllerGetDoctors(filterService IFilterService, listService IListService) *ControllerGetDoctors {
	value := new(ControllerGetDoctors)
	value.filterService = filterService
	value.listService = listService
	return value
}

func (g *ControllerGetDoctors) GetDoctors(c *fiber.Ctx, mode string) error {
	switch mode {

	case common.LIST_ID:
		paramsID := new(ParamsID)
		err := c.ParamsParser(paramsID)

		if err == nil {
			outputData := g.listService.GetList(&paramsID.doctorID)
			status := g.listService.ReturnStatus()
			if status == common.OK {
				return c.JSON(outputData)
			} else {
				return c.Status(fiber.StatusBadRequest).SendString(status)
			}
		} else {
			return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_REQUEST)
		}

	case common.FILTER:
		queryFilters := new(QueryDoctorsListFilter)
		err := c.QueryParser(queryFilters)
		common.CheckErr(err)

		outputData := g.filterService.GetList(queryFilters)
		status := g.filterService.ReturnStatus()
		if status == common.OK {
			return c.JSON(outputData)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString(status)
		}
	default:
		return errors.New("INVALID_MODE")
	}
}

type ControllerCreateDoctors struct {
	createService ICreateService
}

func NewControllerCreateDoctor(createService ICreateService) *ControllerCreateDoctors {
	value := new(ControllerCreateDoctors)
	value.createService = createService
	return value
}

func (cr *ControllerCreateDoctors) CreateDoctor(c *fiber.Ctx) error {
	doctorJson := new(Doctor)

	if err := c.BodyParser(doctorJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
	}
	if description, err := doctorJson.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(description)
	}

	cr.createService.Create(doctorJson)
	return c.Status(fiber.StatusCreated).SendString(cr.createService.ReturnStatus())
}

type ControllerUpdateDeleteDoctors struct {
	updateService IUpdateService
	deleteService IDeleteService
}

func NewControllerUpdateDeleteDoctor(updateService IUpdateService, deleteService IDeleteService) *ControllerUpdateDeleteDoctors {
	value := new(ControllerUpdateDeleteDoctors)
	value.updateService = updateService
	value.deleteService = deleteService
	return value
}

func (u *ControllerUpdateDeleteDoctors) UpdateDeleteDoctor(c *fiber.Ctx, mode string) error {
	ID := c.Params("id")

	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_ID_REQUEST)
	}

	if Verify(&intID) {

		switch mode {
		case common.UPDATE:
			doctorJson := new(DoctorU)
			if err := c.BodyParser(doctorJson); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
			}
			if string, err := doctorJson.validate(); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(string)
			}
			u.updateService.Update(intID, doctorJson)
			return c.Status(fiber.StatusAccepted).SendString(u.updateService.ReturnStatus())

		case common.DELETE:
			doctor := u.deleteService.Delete(&intID)
			status := u.deleteService.ReturnStatus()
			if status == common.OK {
				return c.JSON(doctor)
			} else {
				return c.Status(fiber.StatusForbidden).SendString(status)
			}

		default:
			return errors.New("INVALID_MODE")
		}
	} else {
		return c.Status(fiber.StatusBadRequest).SendString(common.NOT_FOUND_ID)
	}
}
