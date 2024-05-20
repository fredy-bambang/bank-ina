package controllers

import (
	"net/http"
	"strconv"

	"github.com/fredy-bambang/bank-ina/models"
	"github.com/fredy-bambang/bank-ina/services"
	"github.com/gin-gonic/gin"
)

// TaskController is a contract of what task controller can do
type TaskController struct {
	taskService services.TaskService
}

// NewTaskController creates a new task controller
func NewTaskController(ts services.TaskService) *TaskController {
	return &TaskController{
		taskService: ts,
	}
}

// CreateTask is a function to create a new task
func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tc.taskService.CreateTask(&task)
	c.JSON(http.StatusCreated, task)
}

// FindAllTasks is a function to find all tasks
func (tc *TaskController) FindAllTasks(c *gin.Context) {
	tasks, err := tc.taskService.FindAllTasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// FindTaskByID is a function to find a task by id
func (tc *TaskController) FindTaskByID(c *gin.Context) {
	id := c.Param("id")

	// convert id to uint
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, _ := tc.taskService.FindTaskByID(uint(uid))
	c.JSON(http.StatusOK, task)
}

// UpdateTaskByID is a function to update a task by id
func (tc *TaskController) UpdateTaskByID(c *gin.Context) {
	id := c.Param("id")

	// convert id to uint
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task
	if err := c.ShouldBind(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc.taskService.UpdateTaskByID(uint(uid), &task)
	c.JSON(http.StatusOK, task)
}

// DeleteTaskByID is a function to delete a task by id
func (tc *TaskController) DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")

	// convert id to uint
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc.taskService.DeleteTaskByID(uint(uid))
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
