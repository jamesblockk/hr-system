package models

import (
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model            // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string     `gorm:"size:255;unique;not null" json:"name"`
	Employees  []Employee `gorm:"foreignKey:DepartmentID" json:"employees"`
}
