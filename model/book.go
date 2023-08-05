package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id        uint           `gorm:"primaryKey;autoIncrement;column:id;->;<-:create"`
	Title     string         `gorm:"column:title;unique;size:256;default:"";<-"`
	Author    string         `gorm:"column:author;size:256;default:"";<-"`
	CreatedAt time.Time      `gorm:"column:created_at;->;<-:create"`
	UpdatedAt time.Time      `gorm:"column:updated_at;<-"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at;<-"`
}
