package routes

import (
	"micro-mini/admin/controllers"
	"micro-mini/admin/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {

	router := r.Group("/admin")
	{
		router.POST("/login", controllers.Login)
		router.GET("/dashboard", middleware.AdminAuth, controllers.Dashboard)
		router.POST("/logout", middleware.AdminAuth, controllers.Logout)
	}
}
