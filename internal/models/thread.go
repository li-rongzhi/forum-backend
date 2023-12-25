package models

import (
	"time"
)

type Thread struct {
    ThreadID   uint       `json:"thread_id" gorm:"primaryKey"`
    Title      string     `json:"title"`
    Content    string     `json:"content"`
    UserID     uint       `json:"user_id" gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"` // Correctly set as a foreign key.
    Comments   []Comment  `json:"comments"`
    Categories []Category `json:"categories,omitempty" gorm:"many2many:thread_categories;"`
    CreatedAt  time.Time  `json:"created_at"`
}


