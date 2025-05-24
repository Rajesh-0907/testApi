package routes

import (
	"github.com/Rajesh-0907/go-gin-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/question", controllers.GetQuestions)
		api.POST("/score", controllers.GetScore)
	}
}
