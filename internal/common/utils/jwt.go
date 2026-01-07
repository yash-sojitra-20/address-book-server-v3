package utils

import (
	"address-book-server-v3/internal/common/fault"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/samber/mo"
)

// func GenerateToken(jwtSecret string, userID uint64, userEmail string) (string, error) {

// 	claims := jwt.MapClaims{
// 		"user_id": userID,
// 		"user_email": userEmail,
// 		"exp": time.Now().Add(24 * time.Hour).Unix(),
// 	}

// 	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(jwtSecret))
// }

func GenerateToken(jwtSecret string, userID uuid.UUID, userEmail string) mo.Result[*string] {
	
	claims := jwt.MapClaims{
		"user_id": userID,
		"user_email": userEmail,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tk, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return mo.Err[*string](fault.FailedTokenGeneration(err))
	}

	return mo.Ok(&tk)
}