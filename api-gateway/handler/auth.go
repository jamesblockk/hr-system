package handler

import (
	"hr-system/api-gateway/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register route
// @Summary User Registration
// @Description Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/register [post]
func Register(s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := s.Auth.Register(c, req.Email, req.Password)
		if err != nil {
			HandleError(c, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
}

type LoginReq struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"Password123"`
	Name     string `json:"name" example:"John Doe"`
}

// Login route
// @Summary User Login
// @Description Authenticate a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param			name	body		string			true	"{name:John}"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login [post]
func Login(s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		token, err := s.Auth.Login(c, req.Email, req.Password)
		if err != nil {
			HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
