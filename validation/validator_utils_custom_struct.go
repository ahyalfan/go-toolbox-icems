package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// use github.com/go-playground/validator/v10

type ValidatorUtilsStruct struct {
	CheckValidator   map[string]string
	ValidatorMessage map[string]func(vf validator.FieldError) string
}

func (v *ValidatorUtilsStruct) RegisterCustomNameValidatorAndMessage(
	nameValidator string,
	customMessage func(vf validator.FieldError) string,
) error {
	v.CheckValidator[nameValidator] = nameValidator
	v.ValidatorMessage[nameValidator] = customMessage
	return nil
}

func (v *ValidatorUtilsStruct) RegisterCustomValidatorAndMessage(
	validate *validator.Validate,
	customValidate func(fl validator.FieldLevel) bool,
	nameValidator string,
	customMessage func(vf validator.FieldError) string,
) error {
	v.CheckValidator[nameValidator] = nameValidator
	v.ValidatorMessage[nameValidator] = customMessage
	validate.RegisterValidation(nameValidator, customValidate)
	return nil
}

func RegisterAllCustomValidators(validatorUtils ValidatorUtilsStruct) {
	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"required", func(v validator.FieldError) string {
			return fmt.Sprintf("%s is required", v.Field())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"len", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must have length between %s and %s", v.Field(), v.Param(), v.Value())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"email", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be a valid email address", v.Field())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"min", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be between %s and %s", v.Field(), v.Param(), v.Value())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"max", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be less than or equal to %s", v.Field(), v.Value())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"eq", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be equal to %s", v.Field(), v.Param())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"gt", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be greater than %s", v.Field(), v.Param())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"lt", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be less than %s", v.Field(), v.Value())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"contains", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must contain %s", v.Field(), v.Param())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"isdivisibleby", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be divisible by %s", v.Field(), v.Param())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"numeric", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be a numeric value", v.Field())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"isalpha", func(v validator.FieldError) string {
			return fmt.Sprintf("%s must be a alphabetic string", v.Field())
		},
	)

	validatorUtils.RegisterCustomNameValidatorAndMessage(
		"uuid", func(vf validator.FieldError) string {
			return fmt.Sprintf("%s Ensures the string is a valid UUID format", vf.Field())
		})
}
