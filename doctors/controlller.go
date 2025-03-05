package doctors

import (
	"errors"
	"petprojectmed/common"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ControllerGetDoctors(c *fiber.Ctx, mode string) error {
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
		queryFilters := new(QueryDoctorsListFilter)
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

func ControllerCreateDoctor(c *fiber.Ctx) error {
	doctorJson := new(Doctor)

	if err := c.BodyParser(doctorJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
	}
	if err, description := doctorJson.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(description)
	}

	status := doctorJson.Create()
	return c.Status(fiber.StatusCreated).SendString(status)
}

func ControllerUpdateDeleteDoctor(c *fiber.Ctx, mode string) error {
	ID := c.Params("id")

	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(common.INVALID_ID_REQUEST)
	}

	doctorID := doctorID(intID)
	if doctorID.verify() {

		switch mode {
		case common.UPDATE:
			doctorJson := new(DoctorU)
			if err := c.BodyParser(doctorJson); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(common.IVALID_JSON_REQUEST)
			}
			if err, string := doctorJson.validate(); err != nil {
				return c.Status(fiber.StatusBadRequest).SendString(string)
			}
			status := doctorJson.Update(intID)
			return c.Status(fiber.StatusAccepted).SendString(status)

		case common.DELETE:
			status, doctor := doctorID.Delete()
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
