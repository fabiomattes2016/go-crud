package interfaces

import (
	"net/http"
	"strconv"

	"github.com/fabiomattes2016/go-crud/internal/application"
	"github.com/fabiomattes2016/go-crud/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *application.UserService
	AuthService *application.AuthService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.UserService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserService.GetUserByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AuthService.Register(&user); err != nil {
		if err.Error() == "email already registered" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, domain.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(id)

	if err := h.UserService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.UserService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
