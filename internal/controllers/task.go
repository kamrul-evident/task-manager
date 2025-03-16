package controllers

import (
	"net/http"
	"task-manager/internal/database"
	"task-manager/internal/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct{}

func (tc *TaskController) GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	// Check if task exists
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Bind updated fields
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields (only if provided)
	if input.Name != "" {
		task.Name = input.Name
	}
	if input.Description != nil { // Check if provided, since it's a pointer
		task.Description = input.Description
	}
	if input.DueDate != nil {
		task.DueDate = input.DueDate
	}
	if input.Priority != "" && input.Priority != task.Priority {
		task.Priority = input.Priority
	}
	if input.Category != nil {
		task.Category = input.Category
	}
	if input.Status != "" && input.Status != task.Status {
		task.Status = input.Status
	}
	if input.AssigneeID != nil {
		task.AssigneeID = input.AssigneeID
	}

	// Save changes
	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	// Check if task exists
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Delete task (soft delete)
	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}