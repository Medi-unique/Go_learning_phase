package router

import (
	"example.com/task_manager/data"
	"github.com/gin-gonic/gin"
)

func SetRoute(router *gin.Engine) *gin.Engine {

	router.GET("/tasks", data.GetAllTasks)
	router.GET("/tasks/:id", data.GetById)
	router.PUT("/tasks/:id", data.UpdateTask)
	router.DELETE("/tasks/:id", data.DeleteTask)
	router.POST("/tasks", data.CreateTask)

	return router

}
