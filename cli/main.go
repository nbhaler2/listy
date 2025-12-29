package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get API URL from environment or use default
	apiURL := GetAPIURL()
	client := NewAPIClient(apiURL)

	// Check if API is available
	if err := client.CheckHealth(); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("\nMake sure the API server is running:")
		fmt.Println("  cd api && go run main.go")
		fmt.Println("\nOr set LISTY_API_URL environment variable to point to your API server.")
		os.Exit(1)
	}

	// Check if user provided a command
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd(client)

	case "list":
		handleList(client)

	case "pending":
		handlePending(client)

	case "completed":
		handleCompleted(client)

	case "complete":
		handleComplete(client)

	case "incomplete":
		handleIncomplete(client)

	case "toggle":
		handleToggle(client)

	case "update":
		handleUpdate(client)

	case "remove":
		handleRemove(client)

	case "help":
		printHelp()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage: go run main.go <command>")
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
	fmt.Println("  help                 - Show this help message")
	fmt.Println("\nEnvironment Variables:")
	fmt.Println("  LISTY_API_URL        - API server URL (default: http://localhost:8080)")
}

func handleAdd(client *APIClient) {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide an item to add")
		return
	}
	itemName := os.Args[2]
	todo, err := client.CreateTodo(itemName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Added %s (Id: %d)\n", itemName, todo.Id)
}

func handleList(client *APIClient) {
	todos, err := client.GetTodos()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("No Todos found")
		return
	}
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func handlePending(client *APIClient) {
	todos, err := client.GetPendingTodos()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("No pending todos found")
		return
	}
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func handleCompleted(client *APIClient) {
	todos, err := client.GetCompletedTodos()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if len(todos) == 0 {
		fmt.Println("No completed todos found")
		return
	}
	for _, todo := range todos {
		fmt.Println(todo)
	}
}

func handleComplete(client *APIClient) {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide a todo ID")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID. Please provide a number")
		return
	}
	done := true
	todo, err := client.UpdateTodo(id, UpdateTodoRequest{Done: &done})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Todo %d marked as complete\n", todo.Id)
}

func handleIncomplete(client *APIClient) {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide a todo ID")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID. Please provide a number")
		return
	}
	done := false
	todo, err := client.UpdateTodo(id, UpdateTodoRequest{Done: &done})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Todo %d marked as incomplete\n", todo.Id)
}

func handleToggle(client *APIClient) {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide a todo ID")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID. Please provide a number")
		return
	}
	todo, err := client.ToggleTodo(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Todo %d status toggled\n", todo.Id)
}

func handleUpdate(client *APIClient) {
	if len(os.Args) < 4 {
		fmt.Println("Error: Please provide a todo ID and new text")
		fmt.Println("Usage: go run main.go update <id> \"New text\"")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID. Please provide a number")
		return
	}
	newText := os.Args[3]
	item := newText
	todo, err := client.UpdateTodo(id, UpdateTodoRequest{Item: &item})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Todo %d updated successfully\n", todo.Id)
}

func handleRemove(client *APIClient) {
	if len(os.Args) < 3 {
		fmt.Println("Error: Please provide a todo ID")
		return
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error: Invalid ID. Please provide a number")
		return
	}
	err = client.DeleteTodo(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Todo %d removed successfully\n", id)
}
