package routers

import (
	"task-manager/Delivery/controllers"
	"task-manager/Infrastructure/jwt_services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(jwt_services.JWTAuthMiddleware())
	router.GET("/", controllers.HomeHandler)
	return router
}
