package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
)

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func StringToUuid(strUuid string) (uuid.UUID, error) {
	return uuid.Parse(strUuid)
}

var errCodes = map[int32]int{
	2627: fiber.StatusConflict,

	50000: fiber.StatusNotFound, // Custom error, example: "Data tidak ditemukan"
	50001: fiber.StatusConflict,
	50002: fiber.StatusBadRequest,          // Invalid input
	50003: fiber.StatusInternalServerError, // Unexpected server error

}

func ErrorSpToMessageError(e error) (int, string) {
	if e != nil {
		if mssqlErr, ok := e.(mssql.Error); ok {
			if status, found := errCodes[mssqlErr.Number]; found {
				return status, mssqlErr.Message
			}
			return fiber.StatusInternalServerError, mssqlErr.Message
		}
		return fiber.StatusInternalServerError, e.Error()
	}
	return 0, ""
}

func ErrorSpToErrorFiber(e error) error {
	if e != nil {
		if mssqlErr, ok := e.(mssql.Error); ok {

			if status, found := errCodes[mssqlErr.Number]; found {
				return fiber.NewError(status, mssqlErr.Message)
			}
			return fiber.NewError(fiber.StatusInternalServerError, mssqlErr.Message)
		}
		return fiber.NewError(fiber.StatusInternalServerError, e.Error())
	}
	return nil
}
