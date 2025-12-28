package services

import (
	"fmt"
	"listy-api/database"
	"listy-api/models"
	"sort"
)

// GetNextID calculates the next ID based on existing todos
func GetNextID(todos []models.Todo) int {
	if len(todos) == 0 {
		return 1
	}
	maxID := 0
	for _, todo := range todos {
		if todo.Id > maxID {
			maxID = todo.Id
		}
	}
	return maxID + 1
}

// GetAllTodos returns all todos sorted by ID
func GetAllTodos() ([]models.Todo, error) {
	todos, err := database.LoadTodos()
	if err != nil {
		return nil, err
	}

	// Sort by ID
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Id < todos[j].Id
	})

	return todos, nil
}

// GetPendingTodos returns only pending todos
func GetPendingTodos() ([]models.Todo, error) {
	todos, err := GetAllTodos()
	if err != nil {
		return nil, err
	}

	var pending []models.Todo
	for _, todo := range todos {
		if !todo.Done {
			pending = append(pending, todo)
		}
	}

	return pending, nil
}

// GetCompletedTodos returns only completed todos
func GetCompletedTodos() ([]models.Todo, error) {
	todos, err := GetAllTodos()
	if err != nil {
		return nil, err
	}

	var completed []models.Todo
	for _, todo := range todos {
		if todo.Done {
			completed = append(completed, todo)
		}
	}

	return completed, nil
}

// GetTodoByID finds a todo by ID
func GetTodoByID(id int) (*models.Todo, error) {
	todos, err := GetAllTodos()
	if err != nil {
		return nil, err
	}

	for i := range todos {
		if todos[i].Id == id {
			return &todos[i], nil
		}
	}

	return nil, fmt.Errorf("todo with ID %d not found", id)
}

// CreateTodo creates a new todo
func CreateTodo(item string) (*models.Todo, error) {
	// Get all todos to calculate next ID
	todos, err := GetAllTodos()
	if err != nil {
		return nil, err
	}

	nextID := GetNextID(todos)
	newTodo := models.Todo{
		Id:   nextID,
		Item: item,
		Done: false,
	}

	err = database.InsertTodo(newTodo)
	if err != nil {
		return nil, err
	}

	return &newTodo, nil
}

// UpdateTodo updates an existing todo
func UpdateTodo(id int, req models.UpdateTodoRequest) (*models.Todo, error) {
	// Get existing todo
	todo, err := GetTodoByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Item != nil {
		todo.Item = *req.Item
	}
	if req.Done != nil {
		todo.Done = *req.Done
	}

	// Save to database
	err = database.UpdateTodo(id, *todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(id int) error {
	// Check if todo exists
	_, err := GetTodoByID(id)
	if err != nil {
		return err
	}

	return database.DeleteTodo(id)
}

// ToggleTodo toggles the done status of a todo
func ToggleTodo(id int) (*models.Todo, error) {
	todo, err := GetTodoByID(id)
	if err != nil {
		return nil, err
	}

	todo.Done = !todo.Done

	err = database.UpdateTodo(id, *todo)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
