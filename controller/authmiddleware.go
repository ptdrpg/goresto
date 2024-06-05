package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/lib"
)

func (c *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "you'r not authorized",
		})
		return
	}

	jwtKey := []byte(os.Getenv("SECRET_KEY"))

	token = strings.TrimPrefix(token, "Bearer ")
	claims := &lib.AccesClaims{}

	verifToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    ctx.Abort()
    return
	}

	if !verifToken.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
    ctx.Abort()
    return
	}

	ctx.Set("username", claims.Username)
	ctx.Next()
	}
}