package doctors

import (
	"context"
	"errors"
	"petprojectmed/common"
	"petprojectmed/storage"
	"slices"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

const (
	FILTER              = "FILTER"
	LIST_ID             = "LIST_ID"
	IVALID_JSON_REQUEST = "IVALID_JSON_REQUEST"
	INVALID_ID_REQUEST  = "IVALID_ID_REQUEST"
	NOT_FOUND_ID        = "NOT_FOUND_ID"
)

func ControllerGetDoctors(c *fiber.Ctx, mode string) error {
	switch mode {
	case LIST_ID:
		paramsID := new(ParamsID)
		err := c.ParamsParser(paramsID)

		if err == nil {
			outputData, status := paramsID.ID.GetList()
			if status == OK {
				return c.JSON(outputData)
			} else {
				return c.Status(fiber.StatusBadRequest).SendString(status)
			}
		} else {
			return c.Status(fiber.StatusBadRequest).SendString(INVALID_REQUEST)
		}
	case FILTER:
		queryFilters := new(QueryDoctorsListFilter)
		err := c.QueryParser(queryFilters)
		common.CheckErr(err)

		outputData, status := queryFilters.GetList()
		if status == OK {
			return c.JSON(outputData)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString(status)
		}
	}
	return errors.New("INVALID_MODE")
}

func ControllerCreateDoctor(c *fiber.Ctx) error {
	doctorJson := new(Doctor)

	if err := c.BodyParser(doctorJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(IVALID_JSON_REQUEST)
	}

	if err, string := doctorJson.validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(string)
	}

	status := doctorJson.Create()
	return c.Status(fiber.StatusCreated).SendString(status)
}

func ControllerUpdateDoctor(c *fiber.Ctx) error {
	ID := c.Params("id")

	intID, err := strconv.Atoi(ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(INVALID_ID_REQUEST)
	}

	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	if values := storage.GetIDofDoctors(conn); slices.Contains(*values, intID) {
		doctorJson := new(Doctor)
		if err := c.BodyParser(doctorJson); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(IVALID_JSON_REQUEST)
		}

		if err, string := doctorJson.validate(); err != nil {
			log.Debug(err)
			return c.Status(fiber.StatusBadRequest).SendString(string)
		}
		status := doctorJson.Update(intID)
		return c.Status(fiber.StatusCreated).SendString(status)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString(NOT_FOUND_ID)
	}
}
