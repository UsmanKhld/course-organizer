package models

import "time"

type Assignment struct {
	ID          int       `json:"id"`
	CourseID    int       `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Type        string    `json:"type"`
	Completed   bool      `json:"completed"`
}
