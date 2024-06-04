package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/entity"
	"github.com/ptdrpg/resto/lib"
)

func (c *Controller) FindAllItems(ctx *gin.Context) {
	items, err := c.R.FindAllItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": items,
	})
}

func (c *Controller) FindItemById(ctx *gin.Context) {
	itemId := ctx.Param("id")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	item, findItemErr := c.R.FindItemById(id)
	if findItemErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findItemErr.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func (c *Controller) CreateItems(ctx *gin.Context) {
	var item entity.Item
	err := ctx.ShouldBindJSON(&item)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": err.Error(),
		})
		return
	}

	createErr := c.R.CreateItems(&item)
	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": createErr.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"data": item,
	})
}

func (c *Controller) UpdateItems(ctx *gin.Context) {
	var item entity.Item
	err := ctx.ShouldBindJSON(&item)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	itemId := ctx.Param("id")
	id, convId := strconv.Atoi(itemId)
	if convId != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": convId.Error(),
		})
		return
	}
	itemTemp, errTemp := c.R.FindItemById(id)
	if errTemp != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": errTemp.Error(),
		})
		return
	}

	item.ID = itemTemp.ID
	if item.Label == "" {
		item.Label = itemTemp.Label
	}
	if item.Short_desc == "" {
		item.Short_desc = itemTemp.Short_desc
	}
	if item.Price == 0 {
		item.Price = itemTemp.Price
	}
	updateErr := c.R.UpdateItems(&item)
	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": updateErr.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}

func (c *Controller) DeleteItems(ctx *gin.Context) {
	itemId := ctx.Param("id")
	id, convid := strconv.Atoi(itemId)
	if convid != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": convid.Error(),
		})
		return
	}

	deleting := c.R.DeleteItems(id)
	if deleting != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": deleting.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "item succefuly deleted",
	})
}

func (c *Controller) UploadImage(ctx *gin.Context) {
	picture, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	path, imgErr := lib.CreateImage(picture, ctx)
	if imgErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": imgErr.Error(),
		})
		return
	}

	itemID := ctx.Param("id")
	id, convertID := strconv.Atoi(itemID)
	if convertID != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": convertID.Error(),
		})
		return
	}

	item, findItemErr := c.R.FindItemById(id)
	if findItemErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findItemErr.Error(),
		})
		return
	}

	item.Picture = path
	updateItem := c.R.UpdateItems(&item)
	if updateItem != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": updateItem.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": item,
	})
}
