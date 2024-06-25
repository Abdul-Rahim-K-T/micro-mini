package routes

import (
	"micro-mini/user/controllers"

	"micro-mini/user/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// Load HTML templates
	// r.LoadHTMLGlob("user/templates/*.html")

	router := r.Group("/user")
	{
		router.POST("/signup", controllers.Signup)
		router.POST("/signup/validate", controllers.ValidateOtp)
		router.POST("/login", controllers.Login)
		router.GET("/logout", middleware.UserAuth, controllers.Logout)
	}
}
