// package utils
//
// import (
// 	"strconv"
//
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/google/uuid"
// 	mssql "github.com/microsoft/go-mssqldb"
// )
//
// func StringToFloat64(s string) (float64, error) {
// 	return strconv.ParseFloat(s, 64)
// }
//
// func StringToUuid(strUuid string) (uuid.UUID, error) {
// 	return uuid.Parse(strUuid)
// }
//
// var errCodes = map[int32]int{
// 	2627: fiber.StatusConflict,
//
// 	50000: fiber.StatusNotFound, // Custom error, example: "Data tidak ditemukan"
// 	50001: fiber.StatusConflict,
// 	50002: fiber.StatusBadRequest,          // Invalid input
// 	50003: fiber.StatusInternalServerError, // Unexpected server error
//
// }
//
// func ErrorSpToMessageError(e error) (int, string) {
// 	if e != nil {
// 		if mssqlErr, ok := e.(mssql.Error); ok {
// 			if status, found := errCodes[mssqlErr.Number]; found {
// 				return status, mssqlErr.Message
// 			}
// 			return fiber.StatusInternalServerError, mssqlErr.Message
// 		}
// 		return fiber.StatusInternalServerError, e.Error()
// 	}
// 	return 0, ""
// }
//
// func ErrorSpToErrorFiber(e error) error {
// 	if e != nil {
// 		if mssqlErr, ok := e.(mssql.Error); ok {
//
// 			if status, found := errCodes[mssqlErr.Number]; found {
// 				return fiber.NewError(status, mssqlErr.Message)
// 			}
// 			return fiber.NewError(fiber.StatusInternalServerError, mssqlErr.Message)
// 		}
// 		return fiber.NewError(fiber.StatusInternalServerError, e.Error())
// 	}
// 	return nil
// }

package utils

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	mssql "github.com/microsoft/go-mssqldb"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var errCodes = map[int32]int{
	2627: fiber.StatusConflict,

	50000: fiber.StatusNotFound,            // Custom error, example: "Data tidak ditemukan"
	50001: fiber.StatusConflict,            // Data not found
	50002: fiber.StatusBadRequest,          // Invalid input
	50003: fiber.StatusInternalServerError, // Unexpected server error
	50004: fiber.StatusBadRequest,          // Validation error (e.g., input format issue)
	50006: fiber.StatusForbidden,           // Access denied for requested data
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func Float64ToString(s float64) string {
	return strconv.FormatFloat(s, 'f', 2, 64)
}

func StringToUuid(strUuid string) (uuid.UUID, error) {
	return uuid.Parse(strUuid)
}

func StringToDate(strDate string) (time.Time, error) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, strDate)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}

func StringToInt(value string) int {
	var result int
	fmt.Sscanf(value, "%d", &result)
	return result
}

func StringToDatePointer(strDate *string) (*time.Time, error) {
	layout := "2006-01-02"
	if strDate == nil {
		return nil, nil
	}
	parsedTime, err := time.Parse(layout, *strDate)
	if err != nil {
		return nil, err
	}

	return &parsedTime, nil
}

func DateToString(date time.Time) string {
	layout := "2006-01-02"
	return date.Format(layout)
}

func DateToStringPointer(date *time.Time) *string {
	layout := "2006-01-02"
	v := date.Format(layout)
	return &v
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

func ErrorLocation() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "Unknown location"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func CamelCaseToSnakeCase(s string) string {
	re := regexp.MustCompile(`([a-z\d])([A-Z])`)
	snake := re.ReplaceAllString(s, `${1}_${2}`)

	re2 := regexp.MustCompile(`([A-Z])([A-Z][a-z])`)
	snake = re2.ReplaceAllString(snake, `${1}_${2}`)

	return strings.ToLower(snake)
}

func SplitCamelCase(s string) string {
	re := regexp.MustCompile(`([a-z])([A-Z])|([A-Z]+)([A-Z][a-z])`)
	split := re.ReplaceAllString(s, `$1$3 $2$4`)

	words := strings.Fields(split)
	titleCaser := cases.Title(language.English)

	for i, word := range words {
		if word == strings.ToUpper(word) {
			continue
		}
		words[i] = titleCaser.String(word)
	}

	return strings.Join(words, " ")
}
