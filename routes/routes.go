package routes

import (
	"github.com/Rajesh-0907/go-gin-boilerplate/controllers"
	"github.com/Rajesh-0907/go-gin-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/loginuser", controllers.LoginUserHandler)
	auth.POST("/createuser", controllers.RegisterUserHandler)

	api := r.Group("/api")
	r.GET("/", controllers.HomeHandler)
	api.Use(middlewares.AuthMiddleware())
	r.Use(middlewares.RateLimitMiddleware())
	{
		api.GET("/aptitudequestion", controllers.GetAptitudeQuestions)
		api.GET("/ppequestion", controllers.GetPpeQuestions)
		api.POST("/ppescore", controllers.GetPpeScore)
		api.GET("/getuserinfo", controllers.GetUserInfo)
		api.POST("/logoutuser", controllers.LogoutHandler)
		api.GET("/topscore", controllers.GetTopScore)

	}

}
