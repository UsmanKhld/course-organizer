package models

type Course struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Title      string `json:"title"`
	Instructor string `json:"instructor"`
	Semester   string `json:"semester"`
}
