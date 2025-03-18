package patients

import (
	"errors"
	"petprojectmed/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ControllerGetPatients struct {
	filterService IFilterService
	listService   IListService
}

func NewControllerGetPatients(filterService IFilterService, listService IListService) *ControllerGetPatients {
	value := new(ControllerGetPatients)
	value.filterService = filterService
	value.listService = listService
	return value
}

func (g *ControllerGetPatients) GetPatients(c *fiber.Ctx, mode string) error {
	switch mode {
	case common.LIST_ID:
		paramsID := new(ParamsID)
		err := c.ParamsParser(paramsID)

		if err == nil {
			outputData := g.listService.GetList(&paramsID.ID)
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
		queryFilters := new(QueryPatientsListFilter)
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

type ControllerCreatePatients struct {
	createService ICreateService
}

func NewControllerCreatePatients(createService ICreateService) *ControllerCreatePatients {
	value := new(ControllerCreatePatients)
	value.createService = createService
	return value
}

func (cc *ControllerCreatePatients) CreatePatient(c *fiber.Ctx) error {
	patientJson := new(Patient)

	if err := c.BodyParser(patientJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
	}
	if description, err := patientJson.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(description)
	}

	cc.createService.Create(patientJson)
	status := cc.createService.ReturnStatus()
	return c.Status(fiber.StatusCreated).SendString(status)
}

type ControllerUpdateDeletePatients struct {
	updateService IUpdateService
	deleteService IDeleteService
}

func NewControllerUpdateDeletePatient(updateService IUpdateService, deleteService IDeleteService) *ControllerUpdateDeletePatients {
	value := new(ControllerUpdateDeletePatients)
	value.updateService = updateService
	value.deleteService = deleteService
	return value
}

func (u *ControllerUpdateDeletePatients) UpdateDeletePatient(c *fiber.Ctx, mode string) error {
	ID := c.Params("id")

	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_ID_REQUEST)
	}

	if Verify(&intID) {

		switch mode {
		case common.UPDATE:
			patientJson := new(PatientU)
			if err := c.BodyParser(patientJson); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
			}
			if string, err := patientJson.validate(); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(string)
			}
			u.updateService.Update(intID, patientJson)
			status := u.updateService.ReturnStatus()
			return c.Status(fiber.StatusAccepted).SendString(status)

		case common.DELETE:
			patient := u.deleteService.Delete(&intID)
			status := u.deleteService.ReturnStatus()
			if status == common.OK {
				return c.JSON(patient)
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
