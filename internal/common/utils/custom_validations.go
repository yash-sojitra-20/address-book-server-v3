package utils

import (
	"log"
	"regexp"
	"unicode"

	"github.com/go-playground/validator"
)

func registerCustomValidators(validate *validator.Validate) *validator.Validate {

	// Phone number validation
	// This regex checks for a 10-digit phone number starting with digits 6, 7, 8, or 9.
	// It ensures that the phone number consists of exactly 10 digits.
	if err := validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		return regexp.MustCompile(`^[6-9]\d{9}$`).MatchString(fl.Field().String())
	}); err != nil {
		log.Print(err)
	}

	// Email validation
	if err := validate.RegisterValidation("strict_email", func(fl validator.FieldLevel) bool {
		email := fl.Field().String()
		regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		return regex.MatchString(email)
	}); err != nil {
		log.Print(err)
	}

	// Pincode validation
	if err := validate.RegisterValidation("pincode", func(fl validator.FieldLevel) bool {
		pincode := fl.Field().String()

		// Indian PIN code: 6 digits, first digit 1–9
		regex := regexp.MustCompile(`^[1-9][0-9]{5}$`)
		return regex.MatchString(pincode)
	}); err != nil {
		log.Print(err)
	}

	// Password Validation
	if err := validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 8 {
			return false
		}

		var hasUpper, hasLower, hasDigit, hasSpecial bool

		for _, char := range password {
			switch {
			case unicode.IsUpper(char):
				hasUpper = true
			case unicode.IsLower(char):
				hasLower = true
			case unicode.IsDigit(char):
				hasDigit = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true
			}
		}

		return hasUpper && hasLower && hasDigit && hasSpecial

	}); err != nil {
		log.Print(err)
	}

	return validate
}

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}
