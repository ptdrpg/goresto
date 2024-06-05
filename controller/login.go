package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/entity"
	"github.com/ptdrpg/resto/lib"
	"golang.org/x/crypto/bcrypt"
)

type Logininput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Loginresponse struct {
	ID           uint            `gorm:"primary_key" json:"id"`
	Customer     entity.Customer `json:"customer"`
	Username     string          `json:"username"`
	Token        string          `json:"token"`
	RefreshToken string          `json:"refresh_token"`
}

func (c *Controller) Login(ctx *gin.Context) {
	var customer []entity.Customer
	var input Logininput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	allCustomer, fetchCustom := c.R.FindAllUsers()
	if fetchCustom != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fetchCustom.Error(),
		})
		return
	}

	for i := 0; i < len(allCustomer); i++ {
		if allCustomer[i].Email == input.Email {
			customer = append(customer, allCustomer[i])
			break
		}
	}

	var staff []entity.Staff
	allStaff, findStaffErr := c.R.FindAllStaff()
	if findStaffErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "invalide email",
		})
		return
	}

	for i := 0; i < len(allStaff); i++ {
		if allStaff[i].CustomerID == int(customer[0].ID) {
			staff = append(staff, allStaff[i])
			break
		}
	}

	comparePasss := bcrypt.CompareHashAndPassword([]byte(staff[0].Password), []byte(input.Password))
	if comparePasss != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": comparePasss.Error(),
		})
		return
	}

	generatetoken, GenErr := lib.GenerateToken(staff[0].Username)
	if GenErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": GenErr.Error(),
		})
		return
	}

	genRefresh, GenRefreshErr := lib.GenerateRefreshToken(generatetoken)
	if GenRefreshErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": GenRefreshErr.Error(),
		})
		return
	}

	var response Loginresponse
	response.ID = staff[0].ID
	response.Customer = customer[0]
	response.Username = staff[0].Username
	response.Token = generatetoken
	response.RefreshToken = genRefresh

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "welcome",
		"data":    response,
	})
}
