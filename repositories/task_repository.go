package repositories

import (
	"github.com/fredy-bambang/bank-ina/models"
	"gorm.io/gorm"
)

// TaskRepository is a contract of what task repository can do
type TaskRepository interface {
	CreateTask(task *models.Task) error
	FindAllTasks() ([]models.Task, error)
	FindTaskByID(id uint) (models.Task, error)
	UpdateTaskByID(id uint, task *models.Task) error
	DeleteTaskByID(id uint) error
}

// taskRepository is a struct to hold database
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new task repository
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// CreateTask is a function to create a new task
func (r *taskRepository) CreateTask(task *models.Task) error {
	return r.db.Create(task).Error
}

// FindAllTasks is a function to find all tasks
func (r *taskRepository) FindAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

// FindTaskByID is a function to find a task by id
func (r *taskRepository) FindTaskByID(id uint) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return task, err
}

// UpdateTaskByID is a function to update a task by id
func (r *taskRepository) UpdateTaskByID(id uint, task *models.Task) error {
	return r.db.Model(&models.Task{}).Where("id = ?", id).Updates(task).Error
}

// DeleteTaskByID is a function to delete a task by id
func (r *taskRepository) DeleteTaskByID(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
