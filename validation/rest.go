package validation

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func ValidateApiIsNilOrNotNil(data map[string]any) error {
	for key, v := range data {
		if reflect.ValueOf(v).IsNil() {
			return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("id %s not found in data referensi", key))
		}
	}
	return nil
}
