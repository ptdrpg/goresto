package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/entity"
)

func (c *Controller) FindAllUsers(ctx *gin.Context) {
	users, err := c.R.FindAllUsers()
	username, exist := ctx.Get("username")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't get username",
		})
		return
	}

	stringUsername := fmt.Sprintf("%v", username)
	findStaff := c.R.VerifToken(stringUsername)
	if findStaff != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": findStaff.Error(),
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data":   users,
	})
}

func (c *Controller) FindUserById(ctx *gin.Context) {
	customerId := ctx.Param("id")
	id, err := strconv.Atoi(customerId)
	var staff entity.Staff
	username, exist := ctx.Get("username")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't get username",
		})
		return
	}

	if findStaffErr := c.DB.Where("username = ?", username).First(&staff).Error; findStaffErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findStaffErr.Error(),
		})
		return
	}

	if staff.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "you'r not autorized to do this action",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "error parsing id",
		})
		return
	}

	userData, findErr := c.R.FindUserById(id)

	if findErr != nil || userData.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data":   userData,
	})
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var user entity.Customer
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.R.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created",
		"data":    user,
	})
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	var user entity.Customer
	findId := ctx.Param("id")
	userId, errConv := strconv.Atoi(findId)
	var staff entity.Staff
	username, exist := ctx.Get("username")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't get username",
		})
		return
	}

	if findStaffErr := c.DB.Where("username = ?", username).First(&staff).Error; findStaffErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findStaffErr.Error(),
		})
		return
	}

	if staff.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "you'r not autorized to do this action",
		})
		return
	}

	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errConv.Error(),
		})
	}

	usertemp, err := c.R.FindUserById(userId)

	if err != nil || usertemp.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	update := ctx.ShouldBindJSON(&user)
	if update != nil {
		ctx.JSON(http.StatusBadRequest, update.Error())
		return
	}

	user.ID = usertemp.ID
	c.R.UpdateUser(&user)

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user succefuly updated",
		"data":    user,
	})
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	findId := ctx.Param("id")
	userId, errConv := strconv.Atoi(findId)
	var staff entity.Staff
	username, exist := ctx.Get("username")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't get username",
		})
		return
	}

	if findStaffErr := c.DB.Where("username = ?", username).First(&staff).Error; findStaffErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findStaffErr.Error(),
		})
		return
	}

	if staff.Role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "you'r not autorized to do this action",
		})
		return
	}

	if errConv != nil {
		log.Fatal(errConv.Error())
	}

	deleting := c.R.DeleteUser(userId)
	if deleting != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": deleting.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user succefuly deleted",
	})
}
