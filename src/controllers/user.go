package controllers

import (
	models "golang-pagination/src/models"
	"golang-pagination/src/repo"
	utils "golang-pagination/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
