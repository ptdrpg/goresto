package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/entity"
	"github.com/ptdrpg/resto/lib"
	"golang.org/x/crypto/bcrypt"
)

type StaffResponse struct {
	ID           uint            `gorm:"primary_key" json:"id"`
	Customer     entity.Customer `json:"customer"`
	Username     string          `json:"username"`
	Password     string          `json:"password"`
}

type CreateStaffResponse struct {
	ID           uint            `gorm:"primary_key" json:"id"`
	Customer     entity.Customer `json:"customer"`
	Username     string          `json:"username"`
	Password     string          `json:"password"`
	Token        string          `json:"token"`
	RefreshToken string          `json:"refresh_token"`
}

func (c *Controller) FindAllStaff(ctx *gin.Context) {
	staffs, err := c.R.FindAllStaff()
	var allStaffs []StaffResponse

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	for i := 0; i < len(staffs); i++ {
		var staffTemp StaffResponse
		customerId := staffs[i].CustomerID
		findCustomer, findCustomerErr := c.R.FindUserById(int(customerId))
		if findCustomerErr != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": findCustomerErr.Error(),
			})
			return
		}
		staffTemp.ID = staffs[i].ID
		staffTemp.Customer = findCustomer
		staffTemp.Password = staffs[i].Password
		staffTemp.Username = staffs[i].Username
		allStaffs = append(allStaffs, staffTemp)
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": allStaffs,
	})
}

func (c *Controller) FindStaffById(ctx *gin.Context) {
	staffId := ctx.Param("id")
	id, err := strconv.Atoi(staffId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "errror parsing id",
		})
		return
	}

	satffData, findErr := c.R.FindStaffById(id)

	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "satff not found",
		})
		return
	}

	findCustomer, errFindCustomer := c.R.FindUserById(int(satffData.CustomerID))
	if errFindCustomer != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messgae": "staff not found",
		})
		return
	}

	var staffRes StaffResponse
	staffRes.ID = satffData.ID
	staffRes.Username = satffData.Username
	staffRes.Password = satffData.Password
	staffRes.Customer = findCustomer

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": staffRes,
	})
}

type StaffCreateInput struct {
	CustomerID uint   `json:"customer_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type StaffUpdateInput struct {
	CustomerID int    `json:"customer_id"`
	Username   string `json:"username"`
}

func (c *Controller) CreateStaff(ctx *gin.Context) {
	var staff entity.Staff
	var input StaffCreateInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	customer, err := c.R.FindUserById(int(input.CustomerID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "customer not found",
		})
		return
	}

	password := input.Password
	hashedPass, pasErr := bcrypt.GenerateFromPassword([]byte(password), 10)

	if pasErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	staff = entity.Staff{
		CustomerID: int(customer.ID),
		Username:   input.Username,
		Password:   string(hashedPass),
	}

	generateToken, genTokErr := lib.GenerateToken(input.Username)
	if genTokErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messgae": genTokErr.Error(),
		})
		return
	}

	generateRefreshToken, genRefreshError := lib.GenerateRefreshToken(input.Username)
	if genRefreshError != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": genRefreshError.Error(),
		})
		return
	}

	var staffRes CreateStaffResponse
	staffRes.ID = staff.ID
	staffRes.Username = staff.Username
	staffRes.Password = staff.Password
	staffRes.Customer = customer
	staffRes.Token = generateToken
	staffRes.RefreshToken = generateRefreshToken

	err = c.R.CreateStaff(&staff)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "staff succefuly",
		"data": staffRes,
	})
}

func (c *Controller) UpdateStaff(ctx *gin.Context) {
	var staff entity.Staff
	var input StaffUpdateInput
	findId := ctx.Param("id")
	staffId, errConv := strconv.Atoi(findId)
	if errConv != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	staffTemp, err := c.R.FindStaffById(staffId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "satff not found",
		})
		return
	}

	update := ctx.ShouldBindJSON(&input)
	if update != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": update.Error(),
		})
		return
	}

	staff = entity.Staff{
		ID:         staffTemp.ID,
		CustomerID: input.CustomerID,
		Username:   input.Username,
		Password:   staffTemp.Password,
	}
	c.R.UpdateStaff(&staff)

	var staffRes StaffResponse
	findCustomer, errFind := c.R.FindUserById(staff.CustomerID)
	if errFind != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "customer not found",
		})
		return
	}
	staffRes.ID = staff.ID
	staffRes.Customer = findCustomer
	staffRes.Username = staff.Username
	staffRes.Password = staff.Password

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "staff succefuly updated",
		"data":    staffRes,
	})
}

func (c *Controller) DeleteStaff(ctx *gin.Context) {
	findId := ctx.Param("id")
	staffId, errConv := strconv.Atoi(findId)

	if errConv != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "staff not found",
		})
		return
	}

	c.R.DeleteStaff(staffId)

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "staff succefuly deleted",
	})
}

type UpdatePass struct {
	OldPass string `json:"old_password"`
	NewPass string `json:"new_password"`
}

func (c *Controller) UpdatePassword(ctx *gin.Context) {
	var staff entity.Staff
	findId := ctx.Param("id")
	staffId, errConv := strconv.Atoi(findId)
	if errConv != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	staffTemp, err := c.R.FindStaffById(staffId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "staff not found",
		})
		return
	}

	var password UpdatePass

	checkErr := ctx.ShouldBindJSON(&password)
	if checkErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	oldPassword := password.OldPass
	newpassword := password.NewPass

	compare := bcrypt.CompareHashAndPassword([]byte(staffTemp.Password), []byte(oldPassword))

	if compare != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": compare.Error(),
		})
		return
	}

	hashedNewPass, hashing := bcrypt.GenerateFromPassword([]byte(newpassword), 10)

	if hashing != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": hashing.Error(),
		})
		return
	}

	staff.CustomerID = staffTemp.CustomerID
	staff.ID = uint(staffId)
	staff.Password = string(hashedNewPass)
	staff.Username = staffTemp.Username

	c.R.UpdatePassword(&staff)

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "password succefuly updated",
		"data":    staff,
	})

}
