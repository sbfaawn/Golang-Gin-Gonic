package model

import (
	"time"

	"gorm.io/gorm"
)

type Credential struct {
	Id         uint           `gorm:"primaryKey;autoIncrement;column:id;->;<-:create"`
	Email      string         `gotm:"column:email;unique;size:256;"`
	Username   string         `gorm:"column:username;unique;size:256;default:'';"`
	Password   string         `gorm:"column:password;size:2000;default:'';"`
	IsVerified bool           `gorm:"column:verified;default:false"`
	CreatedAt  time.Time      `gorm:"column:created_at;->;<-:create"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;<-"`
	DeletedAt  gorm.DeletedAt `gorm:"index;column:deleted_at;<-"`
}
