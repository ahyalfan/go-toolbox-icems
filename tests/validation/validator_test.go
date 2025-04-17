package validation_test

import (
	"fmt"
	"testing"

	"github.com/ahyalfan/go-toolbox-icems/validation"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type RequestTest struct {
	Username string `validate:"numeric"`
}

type RequestCustomValidateTest struct {
	Username string `validate:"numeric"`

	Yoi string `validate:"trueFalse"`

	Yoa string `validate:"required"`
}

func TestValidationUtils_AddCustomNameAndMessage(t *testing.T) {
	req := RequestTest{
		Username: "yogi",
	}
	validate := validator.New()

	validation.ValidatorUtils.RegisterCustomNameValidatorAndMessage(
		"numeric", func(vf validator.FieldError) string {
			return fmt.Sprintf("%s must be a numeric value", vf.Field())
		},
	)

	fails := validation.Validate(validate, req)
	assert.Equal(t, fails["Username"], "Username must be a numeric value")
}

func TestValidationUtils_AddValidatorCustomAndMessage(t *testing.T) {
	req := RequestCustomValidateTest{
		Username: "yogi",
		Yoi:      "falsee",
		Yoa:      "",
	}
	validate := validator.New()

	customVali := func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return value == "true" || value == "false"
	}

	validation.ValidatorUtils.RegisterCustomValidatorAndMessage(
		validate,
		customVali,
		"trueFalse", func(vf validator.FieldError) string {
			return fmt.Sprintf("%s must be a true or false value", vf.Field())
		},
	)

	fails := validation.Validate(validate, req)
	fmt.Println(fails)
	assert.Equal(t, fails["Yoi"], "Yoi must be a true or false value")
}
