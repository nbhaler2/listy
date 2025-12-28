package models

// Todo represents a todo item
type Todo struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

// CreateTodoRequest represents the request body for creating a todo
type CreateTodoRequest struct {
	Item string `json:"item" binding:"required"`
}

// UpdateTodoRequest represents the request body for updating a todo
type UpdateTodoRequest struct {
	Item *string `json:"item,omitempty"`
	Done *bool   `json:"done,omitempty"`
}
