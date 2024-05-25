package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshToken struct {
	Refresh string `json:"refresh_token"`
}

func (c *Controller) Refresh(ctx *gin.Context) {
	var token RefreshToken
	if err := ctx.ShouldBindJSON(token); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	
}