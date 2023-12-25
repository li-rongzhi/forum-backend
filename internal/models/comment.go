package models

import (
    "time"
)

type Comment struct {
    CommentID uint      `json:"comment_id" gorm:"primaryKey"`
    ThreadID  uint      `json:"thread_id" gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ThreadID"` // Foreign key to Thread
    UserID    uint      `json:"user_id" gorm:"index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"` // Correctly set as a foreign key.
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}

