package utils

import (
	"log"
	"regexp"

	"github.com/go-playground/validator"
)

func registerCustomValidators(validate *validator.Validate) *validator.Validate {

	// RegionCode validation
	// This regex checks for a format like "XX-YY" where X and Y are uppercase letters.
	// It ensures that the first two characters are uppercase letters, followed by a hyphen, and then another two uppercase letters.
	if err := validate.RegisterValidation("regioncode", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return regexp.MustCompile(`^[A-Z]{2}-[A-Z]{2}$`).MatchString(value)
	}); err != nil {
		log.Print(err)
	}

	//Phone number validation
	// This regex checks for a 10-digit phone number starting with digits 6, 7, 8, or 9.
	// It ensures that the phone number consists of exactly 10 digits.
	if err := validate.RegisterValidation("phonenum", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[6-9]\d{9}$`).MatchString(fl.Field().String())
	}); err != nil {
		log.Print(err)
	}

	// PAN number validation
	if err := validate.RegisterValidation("pan", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]{1}$`).MatchString(fl.Field().String())
	}); err != nil {
		log.Print(err)
	}

	if err := validate.RegisterValidation("registrationid", func(fl validator.FieldLevel) bool {
		return !regexp.MustCompile(`^\S+@\S+\.\S+$`).MatchString(fl.Field().String())
	}); err != nil {
		log.Print(err)
	}

	return validate
}
