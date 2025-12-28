package main

import (
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	Id   int
	Item string
	Done bool
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
	for _, todo := range todos {
		if todo.Done == false {
			fmt.Println(todo)
		}
	}
}

func ListCompleteTodos(todos []Todo) {
	for _, todo := range todos {
		if todo.Done == true {
			fmt.Println(todo)
		}
	}
}

func main() {
	var todolist []Todo
	nextId := 1

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
		todolist = AddTodos(todolist, itemName, &nextId)
		fmt.Printf("Added %s (Id: %d)\n", itemName, nextId-1)

	case "list":
		ListTodos(todolist)

	case "pending":
		ListPendingTodos(todolist)

	case "completed":
		ListCompleteTodos(todolist)

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
		if err := MarkCompleteByID(todolist, id); err != nil {
			fmt.Println("Error:", err)
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
		if err := MarkIncompleteByID(todolist, id); err != nil {
			fmt.Println("Error:", err)
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
		if err := ToggleDoneByID(todolist, id); err != nil {
			fmt.Println("Error:", err)
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
		if err := UpdateItemByID(todolist, id, newText); err != nil {
			fmt.Println("Error:", err)
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
		var removeErr error
		todolist, removeErr = RemoveTodos(todolist, id)
		if removeErr != nil {
			fmt.Println("Error:", removeErr)
		} else {
			fmt.Printf("Todo %d removed successfully\n", id)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run without arguments to see available commands")
	}
}
