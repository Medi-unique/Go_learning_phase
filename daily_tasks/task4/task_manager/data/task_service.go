package data

import (
	"net/http"

	"example.com/task_manager/models"

	"github.com/gin-gonic/gin"
)

// var tasks = map[string]models.Task{
var tasks = []models.Task{
	{
		ID:          "1",
		Title:       "Implement User Login",
		Description: "Create user login functionality using JWT.",
		Status:      "In Progress",
	},
	{
		ID:          "2",
		Title:       "Design Home Page",
		Description: "Design the UI for the home page with responsive elements.",
		Status:      "To Do",
	},
	{
		ID:          "3",
		Title:       "Database Migration",
		Description: "Migrate database schema to the latest version.",
		Status:      "Completed",
	},
}

func GetAllTasks(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})

}

func GetById(c *gin.Context) {

	id := c.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return

		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "page not found"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	for _, task := range tasks {
		if task.ID == id {
			if task.Title != "" {
				task.Title = updatedTask.Title
			}
			if task.Description != "" {
				task.Description = updatedTask.Description
			}
			c.JSON(http.StatusAccepted, gin.H{"message": "Task updated successfully"})
			return

		}

	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	for i, task := range tasks {

		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1])
			c.JSON(http.StatusOK, gin.H{"message": "Task removed successfully"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}

}

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "can not bind JSON"})
		return
	}
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "Task added successfully"})

}
