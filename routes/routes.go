package routes

import (
	"github.com/fredy-bambang/bank-ina/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controllers.UserController, taskController *controllers.TaskController) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.FindAllUsers)
		userRoutes.GET("/:id", userController.FindUserByID)
		userRoutes.PUT("/:id", userController.UpdateUserByID)
		userRoutes.DELETE("/:id", userController.DeleteUserByID)
	}

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.POST("", taskController.CreateTask)
		taskRoutes.GET("", taskController.FindAllTasks)
		taskRoutes.GET("/:id", taskController.FindTaskByID)
		taskRoutes.PUT("/:id", taskController.UpdateTaskByID)
		taskRoutes.DELETE("/:id", taskController.DeleteTaskByID)
	}
}
