package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/entity"
	"github.com/ptdrpg/resto/lib"
)

type EmployeeRes struct {
	ID        uint            `json:"id"`
	Customer  entity.Customer `json:"customer"`
	Job       string          `json:"job"`
	Hire_date string          `json:"hire_date"`
}

func (c *Controller) FindAllEmployee(ctx *gin.Context) {
	employes, err := c.R.FindAllEmployee()
	var allEmploye []EmployeeRes

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	for i := 0; i < len(employes); i++ {
		var empleTemp EmployeeRes
		findCustomer, findCustErr := c.R.FindUserById(employes[i].CustomerID)
		if findCustErr != nil || findCustomer.ID < 1 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "customer not found",
			})
			break
		}
		empleTemp.ID = employes[i].ID
		empleTemp.Customer = findCustomer
		empleTemp.Hire_date = employes[i].Hire_date
		empleTemp.Job = employes[i].Job

		allEmploye = append(allEmploye, empleTemp)
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": allEmploye,
	})
}

func (c *Controller) FindEmployeeById(ctx *gin.Context) {
	employeId := ctx.Param("id")
	id, err := strconv.Atoi(employeId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	employe, findEmploye := c.R.FindEmployeeById(id)

	if findEmploye != nil || employe.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": employe,
	})
}

func (c *Controller) CreateEmployee(ctx *gin.Context) {
	var employe entity.Employee
	err := ctx.ShouldBindJSON(&employe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.R.UpdateEmployee(&employe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	findCustomer, findCustErr := c.R.FindUserById(employe.CustomerID)
	if findCustErr != nil || findCustomer.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findCustErr.Error(),
		})
		return
	}

	var response EmployeeRes
	response.ID = employe.ID
	response.Customer = findCustomer
	response.Hire_date = employe.Hire_date
	response.Job = employe.Job

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "employee created",
		"data":    response,
	})
}

func (c *Controller) UpdateEmployee(ctx *gin.Context) {
	var employe entity.Employee
	err := ctx.ShouldBindJSON(&employe)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	CustomerId := ctx.Param("id")
	id, errorId := strconv.Atoi(CustomerId)
	if errorId != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": errorId.Error(),
		})
		return
	}

	employetemp, employeerr := c.R.FindEmployeeById(id)
	if employeerr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}
	employe.ID = employetemp.ID
	if employe.Hire_date == "" {
		employe.Hire_date = employetemp.Hire_date
	}
	if employe.Job == "" {
		employe.Job = employetemp.Job
	}
	updating := c.R.UpdateEmployee(&employe)
	if updating != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": updating.Error(),
		})
		return
	}

	findCustomer, errFindCustom := c.R.FindUserById(employetemp.CustomerID)
	if errFindCustom != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": errFindCustom.Error(),
		})
		return
	}

	var response EmployeeRes
	response.ID = employetemp.ID
	response.Hire_date = employe.Hire_date
	response.Job = employe.Job
	response.Customer = findCustomer

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "employee succefuly updated",
		"data":    response,
	})
}

func (c *Controller) DeleteEmployee(ctx *gin.Context) {
	getid := ctx.Param("id")
	id, err := strconv.Atoi(getid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	deleting := c.R.DeleteEmployee(id)
	if deleting != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": deleting.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "employee succefuly deleted",
	})

}

func (c *Controller) UploadAvatar(ctx *gin.Context) {
	avatar, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	path, imgErr := lib.CreateImage(avatar, ctx)
	if imgErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": imgErr.Error(),
		})
		return
	}

	employeeId := ctx.Param("id")
	id, convErr := strconv.Atoi(employeeId)
	if convErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	employee, findErr := c.R.FindEmployeeById(id)
	if findErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "employee not found",
		})
		return
	}

	employee.Avatar = path
	updateErr := c.R.UpdateEmployee(&employee)
	if updateErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": updateErr.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": employee,
	})
}
