package routes

import (
	"github.com/Rajesh-0907/go-gin-boilerplate/controllers"
	"github.com/Rajesh-0907/go-gin-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/loginuser", controllers.LoginUserHandler)
	api := r.Group("/api")
	r.GET("/", controllers.HomeHandler)
	api.Use(middlewares.AuthMiddleware())
	{
		api.GET("/aptitudequestion", controllers.GetAptitudeQuestions)
		api.GET("/ppequestion", controllers.GetPpeQuestions)
		api.POST("/ppescore", controllers.GetPpeScore)
		api.POST("/createuser", controllers.RegisterUserHandler)
		api.GET("/getname", controllers.GetUsername)
	}

}
