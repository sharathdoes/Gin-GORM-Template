package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Post struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey"`
	Title string
	Body  string
	Tags  pq.StringArray  `gorm:"type:text[]"`
	Comments []Comment  `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UpdatedAt time.Time
}

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	PostID    uint     
	Message   string
	Author    string
	CreatedAt time.Time
}

