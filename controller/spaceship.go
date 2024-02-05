package controller

import (
	"net/http"
	"starwars_api/helper"
	"starwars_api/model"

	"github.com/gin-gonic/gin"
)

func Create(context *gin.Context) {
	var input model.Spaceship
	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	err = input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"success": "true"})
}

func Update(context *gin.Context) {
	var input model.Spaceship

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	err = input.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
	}

	context.JSON(http.StatusAccepted, gin.H{"success": "true"})
}

func FindByName(context *gin.Context) {
	name := context.Param("name")

	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	list, err := model.FindByShipName(name)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
	}

	context.JSON(http.StatusAccepted, gin.H{"data": list})
}

func Delete(context *gin.Context) {
	id := context.Param("id")

	_, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
		return
	}

	err = model.Delete(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"success": "false"})
	}

	context.JSON(http.StatusAccepted, gin.H{"success": "true"})

}
