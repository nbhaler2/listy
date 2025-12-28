package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

type Todo struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

func (i *Todo) UpdateItem(Uname string) {
	if Uname != "" {
		i.Item = Uname
	}
}

func (i *Todo) MarkComplete() {
	i.Done = true
}

func (i *Todo) MarkIncomplete() {
	i.Done = false
}

func (i *Todo) ToggleDone() {
	i.Done = !(i.Done)
}

func AddTodos(todos []Todo, NextItem string, nextID *int) []Todo {
	newTodo := Todo{
		Id:   *nextID,
		Item: NextItem,
		Done: false,
	}
	todos = append(todos, newTodo)
	*nextID++
	return todos
}

func ListTodos(todos []Todo) {
	if len(todos) == 0 {
		fmt.Println("No Todos found")
		return
	}
	// Sort todos by ID before displaying
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Id < todos[j].Id
	})
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func FindTodos(todos []Todo, ItemName string) (int, *Todo) {

	for i, todo := range todos {
		if todo.Item == ItemName {
			return i, &todos[i]
		}

	}
	return -1, nil

}

func FindTodosById(todos []Todo, Id int) (int, *Todo) {

	for i, todo := range todos {
		if todo.Id == Id {
			return i, &todos[i]
		}
	}
	return -1, nil
}

func RemoveTodos(todos []Todo, Id int) ([]Todo, error) {
	i, _ := FindTodosById(todos, Id)
	if i == -1 {
		return todos, fmt.Errorf("todo with Id %d not found", Id)
	}
	todos = append(todos[:i], todos[i+1:]...)
	return todos, nil
}

func MarkCompleteByID(todos []Todo, Id int) error {
	_, todo := FindTodosById(todos, Id)
	if todo == nil {
		return fmt.Errorf("todo with Id %d not found", Id)
	}
	todo.MarkComplete()
	return nil
}

func MarkIncompleteByID(todos []Todo, Id int) error {
	_, todo := FindTodosById(todos, Id)
	if todo == nil {
		return fmt.Errorf("todo with Id %d not found", Id)
	}
	todo.MarkIncomplete()
	return nil
}

func ToggleDoneByID(todos []Todo, Id int) error {
	_, todo := FindTodosById(todos, Id)
	if todo == nil {
		return fmt.Errorf("todo with Id %d not found", Id)
	}
	todo.ToggleDone()
	return nil
}

func UpdateItemByID(todos []Todo, Id int, Newname string) error {
	_, todo := FindTodosById(todos, Id)
	if todo == nil {
		return fmt.Errorf("todo with Id %d not found", Id)
	}
	todo.UpdateItem(Newname)
	return nil
}

func ListPendingTodos(todos []Todo) {
	// Sort todos by ID before displaying
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Id < todos[j].Id
	})
	found := false
	for _, todo := range todos {
		if todo.Done == false {
			fmt.Println(todo)
			found = true
		}
	}
	if !found {
		fmt.Println("No pending todos found")
	}
}

func ListCompleteTodos(todos []Todo) {
	// Sort todos by ID before displaying
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Id < todos[j].Id
	})
	found := false
	for _, todo := range todos {
		if todo.Done == true {
			fmt.Println(todo)
			found = true
		}
	}
	if !found {
		fmt.Println("No completed todos found")
	}
}

// Supabase client (global variable)
var supabaseClient *supabase.Client

// InitSupabase initializes the Supabase client
func InitSupabase() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		// .env file is optional, continue without it
		log.Println("Warning: .env file not found, using environment variables")
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

	supabaseClient = client
	return nil
}

// InsertTodo inserts a single todo into Supabase
func InsertTodo(todo Todo) error {
	if supabaseClient == nil {
		return fmt.Errorf("Supabase client not initialized")
	}

	_, _, err := supabaseClient.From("todos").Insert(todo, false, "", "", "").Execute()
	if err != nil {
		return fmt.Errorf("error inserting todo to Supabase: %v", err)
	}

	return nil
}

// UpdateTodo updates a single todo in Supabase by ID
func UpdateTodo(id int, todo Todo) error {
	if supabaseClient == nil {
		return fmt.Errorf("Supabase client not initialized")
	}

	// Update the todo with the specified ID
	_, _, err := supabaseClient.From("todos").Update(todo, "", "").Eq("id", strconv.Itoa(id)).Execute()
	if err != nil {
		return fmt.Errorf("error updating todo in Supabase: %v", err)
	}

	return nil
}

// DeleteTodo deletes a single todo from Supabase by ID
func DeleteTodo(id int) error {
	if supabaseClient == nil {
		return fmt.Errorf("Supabase client not initialized")
	}

	// Delete with WHERE clause using filter builder
	// Use Eq filter like Update does
	_, _, err := supabaseClient.From("todos").Delete("", "").Eq("id", strconv.Itoa(id)).Execute()
	if err != nil {
		return fmt.Errorf("error deleting todo from Supabase: %v", err)
	}

	return nil
}

// LoadTodos loads todos from Supabase
func LoadTodos() ([]Todo, error) {
	if supabaseClient == nil {
		return nil, fmt.Errorf("Supabase client not initialized")
	}

	var todos []Todo
	data, _, err := supabaseClient.From("todos").Select("*", "", false).Execute()
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
		return []Todo{}, nil
	}

	return todos, nil
}

// GetNextID calculates the next ID based on existing todos
func GetNextID(todos []Todo) int {
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

func main() {
	// Initialize Supabase client
	err := InitSupabase()
	if err != nil {
		fmt.Printf("Error initializing Supabase: %v\n", err)
		fmt.Println("Make sure SUPABASE_URL and SUPABASE_KEY are set in .env file or environment variables")
		os.Exit(1)
	}

	// Load todos from Supabase at startup
	todolist, err := LoadTodos()
	if err != nil {
		fmt.Printf("Warning: Could not load todos: %v\n", err)
		todolist = []Todo{}
	}

	// Calculate next ID from existing todos
	nextId := GetNextID(todolist)

	// Check if user provided a command
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run todolist.go <command>")
		fmt.Println("\nCommands:")
		fmt.Println("  add <item>           - Add a new todo item")
		fmt.Println("  list                 - List all todos")
		fmt.Println("  pending              - List only pending todos")
		fmt.Println("  completed            - List only completed todos")
		fmt.Println("  complete <id>        - Mark a todo as complete")
		fmt.Println("  incomplete <id>      - Mark a todo as incomplete")
		fmt.Println("  toggle <id>          - Toggle todo status")
		fmt.Println("  update <id> <text>   - Update todo item text")
		fmt.Println("  remove <id>          - Remove a todo")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide an item to add")
			return
		}
		itemName := os.Args[2]
		newTodo := Todo{
			Id:   nextId,
			Item: itemName,
			Done: false,
		}
		// Add to local list
		todolist = append(todolist, newTodo)
		nextId++
		// Insert to Supabase
		if err := InsertTodo(newTodo); err != nil {
			fmt.Printf("Error saving todo: %v\n", err)
		} else {
			fmt.Printf("Added %s (Id: %d)\n", itemName, newTodo.Id)
		}

	case "list":
		// Reload from database to get latest state
		loadedTodos, err := LoadTodos()
		if err != nil {
			fmt.Printf("Error loading todos: %v\n", err)
			return
		}
		ListTodos(loadedTodos)

	case "pending":
		// Reload from database to get latest state
		loadedTodos, err := LoadTodos()
		if err != nil {
			fmt.Printf("Error loading todos: %v\n", err)
			return
		}
		ListPendingTodos(loadedTodos)

	case "completed":
		// Reload from database to get latest state
		loadedTodos, err := LoadTodos()
		if err != nil {
			fmt.Printf("Error loading todos: %v\n", err)
			return
		}
		ListCompleteTodos(loadedTodos)

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a todo ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2]) // Convert string to int
		if err != nil {
			fmt.Println("Error: Invalid ID. Please provide a number")
			return
		}
		// Find and update locally
		_, todo := FindTodosById(todolist, id)
		if todo == nil {
			fmt.Printf("Error: todo with ID %d not found\n", id)
			return
		}
		todo.MarkComplete()
		// Update in Supabase
		if err := UpdateTodo(id, *todo); err != nil {
			fmt.Printf("Error updating todo: %v\n", err)
		} else {
			fmt.Printf("Todo %d marked as complete\n", id)
		}

	case "incomplete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a todo ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID. Please provide a number")
			return
		}
		// Find and update locally
		_, todo := FindTodosById(todolist, id)
		if todo == nil {
			fmt.Printf("Error: todo with ID %d not found\n", id)
			return
		}
		todo.MarkIncomplete()
		// Update in Supabase
		if err := UpdateTodo(id, *todo); err != nil {
			fmt.Printf("Error updating todo: %v\n", err)
		} else {
			fmt.Printf("Todo %d marked as incomplete\n", id)
		}

	case "toggle":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a todo ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID. Please provide a number")
			return
		}
		// Find and update locally
		_, todo := FindTodosById(todolist, id)
		if todo == nil {
			fmt.Printf("Error: todo with ID %d not found\n", id)
			return
		}
		todo.ToggleDone()
		// Update in Supabase
		if err := UpdateTodo(id, *todo); err != nil {
			fmt.Printf("Error updating todo: %v\n", err)
		} else {
			fmt.Printf("Todo %d status toggled\n", id)
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide a todo ID and new text")
			fmt.Println("Usage: go run todolist.go update <id> \"New text\"")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID. Please provide a number")
			return
		}
		newText := os.Args[3]
		// Find and update locally
		_, todo := FindTodosById(todolist, id)
		if todo == nil {
			fmt.Printf("Error: todo with ID %d not found\n", id)
			return
		}
		todo.UpdateItem(newText)
		// Update in Supabase
		if err := UpdateTodo(id, *todo); err != nil {
			fmt.Printf("Error updating todo: %v\n", err)
		} else {
			fmt.Printf("Todo %d updated successfully\n", id)
		}

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a todo ID")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID. Please provide a number")
			return
		}
		// Check if todo exists
		_, todo := FindTodosById(todolist, id)
		if todo == nil {
			fmt.Printf("Error: todo with ID %d not found\n", id)
			return
		}
		// Remove from local list
		todolist, err = RemoveTodos(todolist, id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Delete from Supabase
		if err := DeleteTodo(id); err != nil {
			fmt.Printf("Error deleting todo: %v\n", err)
		} else {
			fmt.Printf("Todo %d removed successfully\n", id)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run without arguments to see available commands")
	}
}
