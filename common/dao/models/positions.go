package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Title     string     `gorm:"size:255;unique;not null" json:"title"`
	Level     string     `gorm:"size:100" json:"level"` // 職位等級，如 "Junior", "Senior", "Manager"
	Employees []Employee `gorm:"foreignKey:PositionID" json:"employees"`
}
