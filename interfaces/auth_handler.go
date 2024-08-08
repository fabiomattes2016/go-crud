package interfaces

import (
	"net/http"

	"github.com/fabiomattes2016/go-crud/internal/application"
	"github.com/fabiomattes2016/go-crud/internal/domain"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.AuthService.Login(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
