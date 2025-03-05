package patients

import (
	"errors"
	"petprojectmed/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ControllerGetPatients(c *fiber.Ctx, mode string) error {
	switch mode {
	case common.LIST_ID:
		paramsID := new(ParamsID)
		err := c.ParamsParser(paramsID)

		if err == nil {
			outputData, status := paramsID.ID.GetList()
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

		outputData, status := queryFilters.GetList()
		if status == common.OK {
			return c.JSON(outputData)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString(status)
		}
	default:
		return errors.New("INVALID_MODE")
	}
}

func ControllerCreatePatient(c *fiber.Ctx) error {
	PatientJson := new(Patient)

	if err := c.BodyParser(PatientJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
	}
	if err, description := PatientJson.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(description)
	}

	status := PatientJson.Create()
	return c.Status(fiber.StatusCreated).SendString(status)
}

func ControllerUpdateDeletePatient(c *fiber.Ctx, mode string) error {
	ID := c.Params("id")

	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_ID_REQUEST)
	}

	PatientID := patientID(intID)
	if PatientID.verify() {

		switch mode {
		case common.UPDATE:
			PatientJson := new(PatientU)
			if err := c.BodyParser(PatientJson); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
			}
			if err, string := PatientJson.validate(); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(string)
			}
			status := PatientJson.Update(intID)
			return c.Status(fiber.StatusAccepted).SendString(status)

		case common.DELETE:
			status, Patient := PatientID.Delete()
			if status == common.OK {
				return c.JSON(Patient)
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
