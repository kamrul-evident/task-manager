package routes

import (
	"task-manager/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()

	userCtrl := &controllers.UserController{}
	taskCtrl := &controllers.TaskController{}

	// Route for login
	r.POST("/login", userCtrl.Login)

	// User routes
	r.GET("/users", userCtrl.GetUsers)
	r.POST("/users", userCtrl.CreateUser)
	r.GET("/users/:id", userCtrl.GetUser)
	r.PUT("/users/:id", userCtrl.UpdateUser)
	r.DELETE("/users/:id", userCtrl.DeleteUser)

	// Task routes
	r.GET("/tasks", taskCtrl.GetTasks)
	r.POST("/tasks", taskCtrl.CreateTask)
	r.GET("/tasks/:id", taskCtrl.GetTask)
	r.PUT("/tasks/:id", taskCtrl.UpdateTask)   // New: Update task
	r.DELETE("/tasks/:id", taskCtrl.DeleteTask) // New: Delete task

	// Protected routes (require JWT)
	// api := r.Group("/").Use(middleware.AuthMiddleware())
	// {
	// 	// User routes
	// 	api.GET("/users", userCtrl.GetUsers)
	// 	api.POST("/users", userCtrl.CreateUser)
	// 	api.GET("/users/:id", userCtrl.GetUser)
	// 	api.PUT("/users/:id", userCtrl.UpdateUser)
	// 	api.DELETE("/users/:id", userCtrl.DeleteUser)

	// 	// Task routes
	// 	api.GET("/tasks", taskCtrl.GetTasks)
	// 	api.POST("/tasks", taskCtrl.CreateTask)
	// 	api.GET("/tasks/:id", taskCtrl.GetTask)
	// 	api.PUT("/tasks/:id", taskCtrl.UpdateTask)
	// 	api.DELETE("/tasks/:id", taskCtrl.DeleteTask)
	// }

	return r
}