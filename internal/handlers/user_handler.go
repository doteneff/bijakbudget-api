package handlers

import (
	"net/http"

	"github.com/doteneff/bijakbudget-api/internal/models"
	"github.com/doteneff/bijakbudget-api/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) Register(c *gin.Context) {
        var body models.User
        if err := c.ShouldBindJSON(&body); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
        }

        authData, err := h.service.RegisterUser(&body)
        if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
        }

        c.JSON(http.StatusCreated, authData)
}

func (h *UserHandler) Login(c *gin.Context) {
        var body struct {
                Email    string `json:"email" binding:"required"`
                Password string `json:"password" binding:"required"`
        }

        if err := c.ShouldBindJSON(&body); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
        }

        authData, err := h.service.LoginUser(body.Email, body.Password)
        if err != nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                return
        }

        c.JSON(http.StatusOK, authData)
}

func (h *UserHandler) SSOGoogle(c *gin.Context) {
        var body struct {
                IdToken string `json:"idToken" binding:"required"`
        }

        if err := c.ShouldBindJSON(&body); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
        }

        authData, err := h.service.LoginSSOGoogle(body.IdToken)
        if err != nil {
                c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                return
        }

        c.JSON(http.StatusOK, authData)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
