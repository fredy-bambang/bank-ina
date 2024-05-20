package services

import (
	"github.com/fredy-bambang/bank-ina/models"
	"github.com/fredy-bambang/bank-ina/repositories"
)

// TaskService is a contract of what task service can do
type TaskService interface {
	CreateTask(task *models.Task) error
	FindAllTasks() ([]models.Task, error)
	FindTaskByID(id uint) (models.Task, error)
	UpdateTaskByID(id uint, task *models.Task) error
	DeleteTaskByID(id uint) error
}

// taskService is a struct to hold task repository
type taskService struct {
	taskRepository repositories.TaskRepository
}

// NewTaskService creates a new task service
func NewTaskService(tr repositories.TaskRepository) TaskService {
	return &taskService{
		taskRepository: tr,
	}
}

// CreateTask is a function to create a new task
func (s *taskService) CreateTask(task *models.Task) error {
	return s.taskRepository.CreateTask(task)
}

// FindAllTasks is a function to find all tasks
func (s *taskService) FindAllTasks() ([]models.Task, error) {
	return s.taskRepository.FindAllTasks()
}

// FindTaskByID is a function to find a task by id
func (s *taskService) FindTaskByID(id uint) (models.Task, error) {
	return s.taskRepository.FindTaskByID(id)
}

// UpdateTaskByID is a function to update a task by id
func (s *taskService) UpdateTaskByID(id uint, task *models.Task) error {
	return s.taskRepository.UpdateTaskByID(id, task)
}

// DeleteTaskByID is a function to delete a task by id
func (s *taskService) DeleteTaskByID(id uint) error {
	return s.taskRepository.DeleteTaskByID(id)
}
