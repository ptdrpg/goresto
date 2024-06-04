package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/resto/entity"
	"github.com/ptdrpg/resto/lib"
)

func (c *Controller) FindEntrepriesById(ctx *gin.Context) {
	entrepriseID := ctx.Param("id")
	id, convErr := strconv.Atoi(entrepriseID)
	if convErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": convErr.Error(),
		})
		return
	}
	entreprise, err := c.R.FindEntrepriseById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"data": entreprise,
	})	
}

func (c *Controller) CreateEntreprise(ctx *gin.Context) {
	var entreprise entity.Entreprise
	err := ctx.ShouldBindJSON(entreprise)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	staff, findStaffErr := c.R.FindStaffById(int(entreprise.AdminID))
	if findStaffErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": findStaffErr,
		})
		return
	}

	createErr := c.R.CreateEntreprise(&entreprise)
	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": createErr.Error(),
		})
		return
	}

	staff.EntrepriseID = int(entreprise.ID)
	updateStaff := c.R.UpdateStaff(&staff)
	if updateStaff != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": updateStaff.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"data": entreprise,
	})
}

func (c *Controller) DeleteEntreprise(ctx *gin.Context) {
	entrepriseID := ctx.Param("id")
	id, convErr := strconv.Atoi(entrepriseID)
	if convErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": convErr.Error(),
		})
		return
	}
	
	entreprise, err := c.R.FindEntrepriseById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return	
	}
	
	staff, staffErr := c.R.FindStaffById(int(entreprise.AdminID))
	if staffErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": staffErr.Error(),
		})
		return
	}

	staff.EntrepriseID = 0
	c.R.UpdateStaff(&staff)

	deleting := c.R.DeleteEntreprise(id)
	if deleting != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": deleting.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "entreprise succefuly deleted",
	})
}

func (c *Controller) UploadEntrepriseAvatar(ctx *gin.Context) {
	entrepriseID := ctx.Param("id")
	id, convErr := strconv.Atoi(entrepriseID)
	if convErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": convErr.Error(),
		})
		return
	}

	avatar, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	path, pathErr := lib.CreateImage(avatar, ctx)
	if pathErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": pathErr.Error(),
		})
		return
	}

	entreprise, epErr := c.R.FindEntrepriseById(id)
	if epErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": epErr.Error(),
		})
		return
	}

	entreprise.Avatar = path
	c.R.UpdateEntreprise(&entreprise)

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": entreprise,
	})
}
