package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"listy-api/models"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

// InitSupabase initializes the Supabase client
func InitSupabase() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		// Try loading from parent directory
		err = godotenv.Load("../.env")
		if err != nil {
			log.Println("Warning: .env file not found, using environment variables")
		}
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		return fmt.Errorf("SUPABASE_URL and SUPABASE_KEY must be set in environment variables or .env file")
	}

	client, err := supabase.NewClient(supabaseURL, supabaseKey, nil)
	if err != nil {
		return fmt.Errorf("failed to create Supabase client: %v", err)
	}

	Client = client
	return nil
}

// InsertTodo inserts a single todo into Supabase
func InsertTodo(todo models.Todo) error {
	if Client == nil {
		return fmt.Errorf("Supabase client not initialized")
	}

	_, _, err := Client.From("todos").Insert(todo, false, "", "", "").Execute()
	if err != nil {
		return fmt.Errorf("error inserting todo to Supabase: %v", err)
	}

	return nil
}

// UpdateTodo updates a single todo in Supabase by ID
func UpdateTodo(id int, todo models.Todo) error {
	if Client == nil {
		return fmt.Errorf("Supabase client not initialized")
	}

	_, _, err := Client.From("todos").Update(todo, "", "").Eq("id", strconv.Itoa(id)).Execute()
	if err != nil {
		return fmt.Errorf("error updating todo in Supabase: %v", err)
	}

	return nil
}

// DeleteTodo deletes a single todo from Supabase by ID
func DeleteTodo(id int) error {
	if Client == nil {
		return fmt.Errorf("Supabase client not initialized")
	}

	_, _, err := Client.From("todos").Delete("", "").Eq("id", strconv.Itoa(id)).Execute()
	if err != nil {
		return fmt.Errorf("error deleting todo from Supabase: %v", err)
	}

	return nil
}

// LoadTodos loads all todos from Supabase
func LoadTodos() ([]models.Todo, error) {
	if Client == nil {
		return nil, fmt.Errorf("Supabase client not initialized")
	}

	var todos []models.Todo
	data, _, err := Client.From("todos").Select("*", "", false).Execute()
	if err != nil {
		return nil, fmt.Errorf("error loading todos from Supabase: %v", err)
	}

	// Parse the JSON response
	if len(data) > 0 {
		err = json.Unmarshal(data, &todos)
		if err != nil {
			return nil, fmt.Errorf("error parsing todos: %v", err)
		}
	}

	// If no todos found, return empty slice
	if todos == nil {
		return []models.Todo{}, nil
	}

	return todos, nil
}
