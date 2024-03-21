package models

import "time"

type TaskCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	DueDate     time.Time `json:"dueDate"`
}

type TaskUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	DueDate     time.Time `json:"dueDate"`
}