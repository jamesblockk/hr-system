package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model             // 包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Username    string     `gorm:"size:100;unique;not null" json:"username"` // 帳號名稱
	Password    string     `gorm:"size:255;not null" json:"-"`               // 密碼 (加密儲存，不回傳)
	Email       string     `gorm:"size:100;unique;not null" json:"email"`    // 電子郵件
	Phone       string     `gorm:"size:20" json:"phone"`                     // 電話號碼
	Role        string     `gorm:"size:50;default:'user'" json:"role"`       // 角色，如 user/admin
	Status      string     `gorm:"size:50;default:'active'" json:"status"`   // 帳號狀態，如 active/inactive
	LastLoginAt *time.Time `json:"last_login_at"`                            // 最後登入時間
}
