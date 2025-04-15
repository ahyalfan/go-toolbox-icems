package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Validate[T any](validate *validator.Validate, data T) map[string]string {
	err := validate.Struct(data)
	res := make(map[string]string)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			switch v.ActualTag() {
			case "required":
				res[v.StructField()] = fmt.Sprintf("%s is required", v.Field())
				continue

			case "len":
				res[v.StructField()] = fmt.Sprintf("%s must have length between %s and %s", v.Field(), v.Param(), v.Value())
				continue
			case "email":
				res[v.StructField()] = fmt.Sprintf("%s must be a valid email address", v.Field())
				continue

			case "min":
				res[v.StructField()] = fmt.Sprintf("%s must be between %s and %s", v.Field(), v.Param(), v.Value())
				continue
			case "max":
				res[v.StructField()] = fmt.Sprintf("%s must be less than or equal to %s", v.Field(), v.Value())
				continue
			case "eq":
				res[v.StructField()] = fmt.Sprintf("%s must be equal to %s", v.Field(), v.Param())
				continue
			case "gt":
				res[v.StructField()] = fmt.Sprintf("%s must be greater than %s", v.Field(), v.Param())
				continue
			case "lt":
				res[v.StructField()] = fmt.Sprintf("%s must be less than %s", v.Field(), v.Value())
				continue
			case "contains":
				res[v.StructField()] = fmt.Sprintf("%s must contain %s", v.Field(), v.Param())
				continue
			case "isdivisibleby":
				res[v.StructField()] = fmt.Sprintf("%s must be divisible by %s", v.Field(), v.Param())
				continue
			case "numeric":
				res[v.StructField()] = fmt.Sprintf("%s must be a numeric value", v.Field())
				continue
			case "isalpha":
				res[v.StructField()] = fmt.Sprintf("%s must be a alphabetic string", v.Field())
				continue
			default:
				res[v.StructField()] = fmt.Sprintf("Validation error for field %s: %s", CamelCaseToReadable(v.Field()), v.ActualTag())
				continue
			}
		}
	}
	return res
}

func CamelCaseToReadable(input string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")

	output := re.ReplaceAllString(input, "${1} ${2}")

	words := strings.Split(cases.Lower(language.Tag{}).String(output), " ")

	if len(words) > 0 {
		words[0] = cases.Title(language.Tag{}).String(words[0])
	}

	return strings.Join(words, " ")
}

func ValidatePhoneNumber(phone string) (bool, error) {
	pattern := `^(\+62|62|0)(\d{3,4})(\d{3,4})(\d{4})$`
	return regexp.MatchString(pattern, phone)
}
