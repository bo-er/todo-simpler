package models

import "time"

// Base 是基础模型
type Base struct {
	ID        uint64    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt" gorm:"DEFAULT CURRENT_TIMESTAMP"`
	UpdateAt  time.Time `json:"updateAt" gorm:"DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
