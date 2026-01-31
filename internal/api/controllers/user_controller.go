package controllers

import (
	"net/http"
	"strconv"

	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/models"
	"github.com/ThanhTNV/RestfulAPI-GoLang/internal/service"
	"github.com/gin-gonic/gin"
)

// UserController handles HTTP requests for user operations
type UserController struct {
	userService service.UserService
}

// NewUserController creates a new UserController
func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// GetUsers handles GET /users
func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser handles GET /users/:id
func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := ctrl.parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser handles POST /users
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.userService.CreateUser(&user); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles PUT /users/:id
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := ctrl.parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.userService.UpdateUser(id, &updateData); err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Error(err)
		return
	}

	// Fetch updated user to return
	user, err := ctrl.userService.GetUserByID(id)
	if err != nil {
		// This should not happen since we just updated successfully
		// but handle it just in case
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE /users/:id
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := ctrl.parseID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := ctrl.userService.DeleteUser(id); err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// parseID is a helper function to parse and validate ID from URL parameter
func (ctrl *UserController) parseID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
