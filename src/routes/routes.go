package routes

import (
	"golang-pagination/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	u1 := r.Group("/admin")
	{
		u1.GET("all", controllers.GetAllUsers)
	}

	return r
}
