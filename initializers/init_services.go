package initializers

import (
	"github.com/fredy-bambang/bank-ina/repositories"
	"github.com/fredy-bambang/bank-ina/services"
)

var (
	userService services.UserService
	taskService services.TaskService
)

// InitServices initializes all services
func InitServices() {
	// init repositories
	userRepo := repositories.NewUserRepository(DB)
	taskRepo := repositories.NewTaskRepository(DB)

	// init services
	// UserService := services.NewUserService(userRepo)
	userService = services.NewUserService(userRepo)
	taskService = services.NewTaskService(taskRepo)
}
