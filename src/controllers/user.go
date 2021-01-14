package controllers

import (
	"golang-pagination/src/models"
	"golang-pagination/src/repo"
	"golang-pagination/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)

	var user models.User

	userLists, err := repo.GetAllUsers(&user, &pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userLists,
	})
}
