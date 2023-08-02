package model

import (
	"database/sql"
	"time"
)

type Book struct {
	Id        uint         `gorm:"primaryKey;autoIncrement;column:id;->;<-:create"`
	Title     string       `gorm:"column:title;->;<-:create"`
	Author    string       `gorm:"column:author;->;<-:create"`
	CreatedAt time.Time    `gorm:"column:created_at;->;<-:create"`
	UpdatedAt time.Time    `gorm:"column:updated_at;->;<-:create"`
	DeletedAt sql.NullTime `gorm:"index"`
}
