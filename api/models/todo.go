package models

// Todo represents a todo item
type Todo struct {
	Id            int     `json:"id"`
	Item          string  `json:"item"`
	Done          bool    `json:"done"`
	ListId        *string `json:"list_id,omitempty"`        // NULL means main list, otherwise it's a list identifier
	Priority      *string `json:"priority,omitempty"`        // "high", "medium", "low"
	EstimatedTime *string `json:"estimated_time,omitempty"` // e.g., "30 minutes", "1 hour"
	Category      *string `json:"category,omitempty"`        // Optional category/tag
}

// CreateTodoRequest represents the request body for creating a todo
type CreateTodoRequest struct {
	Item   string  `json:"item" binding:"required"`
	ListId *string `json:"list_id,omitempty"` // Optional: if provided, adds to specific list
}

// UpdateTodoRequest represents the request body for updating a todo
type UpdateTodoRequest struct {
	Item *string `json:"item,omitempty"`
	Done *bool   `json:"done,omitempty"`
}
