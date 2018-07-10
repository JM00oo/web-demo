package model

import "time"

type Post struct {
	ID        string    `json:"id" db:"id"`
	Content   string    `json:"content" binding:"required" db:"content"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
