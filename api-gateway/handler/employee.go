package handler

import (
	"fmt"
	"hr-system/api-gateway/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetEmployeeByID retrieves a single employee by ID
// @Summary Get an employee
// @Description Fetch a single employee by ID
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/employees/{id} [get]
func GetEmployeeByID(s *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		employee, err := s.Employee.GetEmployeeByID(c, uint64(id))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		c.JSON(http.StatusOK, employee)
	}
}
