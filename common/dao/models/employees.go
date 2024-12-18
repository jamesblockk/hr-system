package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model              // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Name         string     `gorm:"size:255;not null" json:"name"`
	Email        string     `gorm:"size:255;unique;not null" json:"email"`
	Phone        string     `gorm:"size:20" json:"phone"`
	DepartmentID uint       `json:"department_id"`
	PositionID   uint       `json:"position_id"`
	HireDate     time.Time  `json:"hire_date"`
	Salary       float64    `json:"salary"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"department"`
	Position     Position   `gorm:"foreignKey:PositionID" json:"position"`
}
