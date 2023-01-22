package dao

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey";type:longtext`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(20);not null"`
	Telephone string         `gorm:"type:varchar(20);not null;unipue"`
	Password  string         `gorm:"size:255";not null`
}
