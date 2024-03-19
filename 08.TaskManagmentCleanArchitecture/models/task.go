package models


type TaskCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	DueDate     string `json:"dueDate"`
}


type TaskUpdate struct {
	Title	   string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	DueDate     string `json:"dueDate"`
}