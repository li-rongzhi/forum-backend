package models

import (
    "time"
)

type User struct {
    UserID      uint      `json:"user_id" gorm:"primaryKey"`
    UserName    string    `json:"user_name"`
    Password    string    `json:"-"`
    CreatedAt   time.Time `json:"created_at"`
    Threads     []Thread  `json:"threads"`
    Comments    []Comment `json:"comments"`
}
