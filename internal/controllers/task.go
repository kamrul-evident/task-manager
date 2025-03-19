package controllers

import (
	"time"
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

	// Fetch the task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Define input struct with only the fields we want to update
	var input struct {
		Name        string          `json:"name"`
		Description *string         `json:"description"`
		DueDate     *time.Time      `json:"due_date"`
		Priority    models.Priority `json:"priority"`
		Category    *string         `json:"category"`
		Status      models.TaskStatus `json:"status"`
		AssigneeID  *uint           `json:"assignee_id"` // Expect user ID or null
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if input.Name != "" {
		task.Name = input.Name
	}
	if input.Description != nil {
		task.Description = input.Description
	}
	if input.DueDate != nil {
		task.DueDate = input.DueDate
	}
	if input.Priority != "" {
		task.Priority = input.Priority
	}
	if input.Category != nil {
		task.Category = input.Category
	}
	if input.Status != "" {
		task.Status = input.Status
	}
	// Update AssigneeID if provided (can be a new ID or null)
	if input.AssigneeID != nil || task.AssigneeID != nil { // Check if assignee is changing
		if input.AssigneeID != nil && (task.AssigneeID == nil || *input.AssigneeID != *task.AssigneeID) ||
			(input.AssigneeID == nil && task.AssigneeID != nil) {
			task.AssigneeID = input.AssigneeID
			// Optional validation: Ensure user exists if ID is provided
			if input.AssigneeID != nil {
				var user models.User
				if err := database.DB.First(&user, *input.AssigneeID).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Assignee user not found"})
					return
				}
			}
		}
	}

	// Save the updated task
	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Optional: Preload Assignee for response
	if err := database.DB.Preload("Assignee").First(&task, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load assignee"})
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