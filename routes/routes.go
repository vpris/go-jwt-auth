package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vpris/test-jwt/controllers"
	"github.com/vpris/test-jwt/middleware"
)

func RoutesInit() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/healthz", healthz)
		api.POST("/auth/register", controllers.Signup)
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/logout", controllers.Logout)
		// api.POST("/auth/refresh", controllers.RefreshToken)
		secure := api.Group("/secure").Use(middleware.RequireAuth)
		{
			secure.GET("/validate", controllers.Validate)

		}
		admin := api.Group("/admin").Use(middleware.IsAdmin, middleware.RequireAuth)
		{
			admin.GET("/", controllers.ValidAdmin)
		}
	}
	return router
}

func healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "ok",
	})
}
