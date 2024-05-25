package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (c *Controller) ListOrder(ctx *gin.Context) {
	orders, err := c.R.ListOrder()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

type OrderCreateInput struct {
	Name string `json:"name"`
}

func (c *Controller) CreateOrder(ctx *gin.Context) {
	var input OrderCreateInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"input": input})
}
