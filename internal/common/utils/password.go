package utils

import (
	"address-book-server-v3/internal/common/fault"

	"github.com/samber/lo"
	"github.com/samber/mo"
	"golang.org/x/crypto/bcrypt"
)

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// func ComparePassword(hashedPassword, password string) error {
// 	return bcrypt.CompareHashAndPassword(
// 		[]byte(hashedPassword),
// 		[]byte(password),
// 	)
// }

func HashPassword(password string) mo.Result[*string] {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return mo.Err[*string](fault.InternalServerError(err))
	}
	return mo.Ok(lo.ToPtr(string(bytes)))
}

func ComparePassword(hashedPassword, password string) mo.Result[*bool] {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	if err != nil {
		return mo.Err[*bool](fault.InternalServerError(err))
	}

	return mo.Ok(lo.ToPtr(err == nil))
}
