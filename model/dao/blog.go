package dao

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	UUID      string
	BlogID    string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(20);not null"`
	CoverPath string         `gorm:"type:varchar(100);not null;unipue"`
	CoverName string         `gorm:"type:varchar(100);not null;unipue"`
	Html      string         `gorm:"type:longtext;not null"`
}
