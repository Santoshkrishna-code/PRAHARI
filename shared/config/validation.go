package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	validate    *validator.Validate
	arnRegex    = regexp.MustCompile(`^arn:aws[a-zA-Z-]*:[a-zA-Z0-9_-]+:[a-z0-9-]*:[0-9]{12}:[a-zA-Z0-9_.-]+(:[a-zA-Z0-9_.-]+)?$`)
	regionRegex = regexp.MustCompile(`^[a-z]{2}-[a-z]+-[0-9]$`)
)

func init() {
	validate = validator.New()

	// 1. Port range validation (1-65535)
	_ = validate.RegisterValidation("port", func(fl validator.FieldLevel) bool {
		val := fl.Field().Int()
		return val >= 1 && val <= 65535
	})

	// 2. AWS Region validation (e.g. us-east-1)
	_ = validate.RegisterValidation("aws_region", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		return regionRegex.MatchString(val)
	})

	// 3. AWS Resource ARN validation
	_ = validate.RegisterValidation("aws_arn", func(fl validator.FieldLevel) bool {
		val := fl.Field().String()
		return arnRegex.MatchString(val)
	})
}

// ValidateStruct executes play-ground validations across the fields of a configuration struct.
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
