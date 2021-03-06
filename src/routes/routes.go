package routes

import (
	controllers "golang-pagination/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	u1 := r.Group("/admin")
	{
		u1.GET("all", controllers.GetAllUsers)
	}

	return r
}
