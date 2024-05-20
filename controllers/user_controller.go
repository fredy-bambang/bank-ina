package controllers

import (
	"net/http"
	"strconv"

	"github.com/fredy-bambang/bank-ina/models"
	"github.com/fredy-bambang/bank-ina/services"
	"github.com/gin-gonic/gin"
)

// UserController is a contract of what user controller can do
type UserController struct {
	userService services.UserService
}

// NewUserController creates a new user controller
func NewUserController(us services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

// CreateUser is a function to create a new user
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uc.userService.CreateUser(&user)
	c.JSON(http.StatusCreated, user)
}

// FindAllUsers is a function to find all users
func (uc *UserController) FindAllUsers(c *gin.Context) {
	users, _ := uc.userService.FindAllUsers()
	c.JSON(http.StatusOK, users)
}

// FindUserByID is a function to find a user by id
func (uc *UserController) FindUserByID(c *gin.Context) {
	id := c.Param("id")

	// convert id to uint
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := uc.userService.FindUserByID(uint(uid))
	c.JSON(http.StatusOK, user)
}

// UpdateUserByID is a function to update a user by id
func (uc *UserController) UpdateUserByID(c *gin.Context) {
	id := c.Param("id")

	// convert id to uint
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.userService.UpdateUserByID(uint(uid), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUserByID is a function to delete a user by id
func (uc *UserController) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	// convert id to uint
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc.userService.DeleteUserByID(uint(uid))
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
