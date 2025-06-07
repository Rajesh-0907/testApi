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
	r.GET("/topscore", controllers.GetTopScore)
	api.Use(middlewares.AuthMiddleware())
	r.Use(middlewares.RateLimitMiddleware())
	{
		api.GET("/aptitudequestion", controllers.GetAptitudeQuestions)
		api.GET("/ppequestion", controllers.GetPpeQuestions)
		api.POST("/ppescore", controllers.GetPpeScore)
		api.POST("/createuser", controllers.RegisterUserHandler)
		api.GET("/getuserinfo", controllers.GetUserInfo)
		api.POST("/logoutuser", controllers.LogoutHandler)
	}

}
