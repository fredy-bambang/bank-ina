package initializers

import "github.com/fredy-bambang/bank-ina/controllers"

var (
	UserController controllers.UserController
	TaskController controllers.TaskController
)

// InitControllers initializes all controllers
func InitControllers() {
	UserController = *controllers.NewUserController(userService)
	TaskController = *controllers.NewTaskController(taskService)
}
