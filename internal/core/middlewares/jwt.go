package middlewares

import (
	// "fmt"
	"strings"

	"address-book-server-v3/internal/common/fault"
	"address-book-server-v3/internal/core/application"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/mo"
)

func AuthMiddleware(application application.Application, c *gin.Context) mo.Result[*bool] {
	appCfg := application.GetConfig()
	jwtSecret := appCfg.GetSecretKey()
	
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return mo.Err[*bool](fault.AuthTokenNotFoundError())
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return mo.Err[*bool](fault.AuthTokenNotFoundError())
	}
	
	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		
		return mo.Err[*bool](fault.AuthTokenInvalidError(err))
	}

	claims := token.Claims.(jwt.MapClaims)
	// fmt.Println("===============> Auth Md : ==================> user_id: *Before String conv: ", claims["user_id"])
	userID := string(claims["user_id"].(string))
	// fmt.Println("===============> Auth Md : ==================> user_id: *After String conv: ", userID)
	// fmt.Println("===============> Auth Md : ==================> user_email: *Before String conv: ", claims["user_email"])
	userEmail := string(claims["user_email"].(string))
	// fmt.Println("===============> Auth Md : ==================> user_email: *After String conv: ", userEmail)

	// must pass pointer in : bitbucket.org/vayana/walt-gin-gonic v1.0.1 or older
	// c.Set("user_id", &userID)

	// can pass pointer or value in : bitbucket.org/vayana/walt-gin-gonic v1.0.2
	c.Set("user_id", userID)
	c.Set("user_email", userEmail)

	validToken := true
	return mo.Ok(&validToken)
}
