package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health Check route
// @Summary Health Check
// @Description Check if API Gateway is healthy
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "API Gateway is healthy"})
}
