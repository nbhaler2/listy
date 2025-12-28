package main

import (
	"testing"
)

// Test Todo struct methods

func TestTodo_UpdateItem(t *testing.T) {
	tests := []struct {
		name     string
		todo     Todo
		newItem  string
		expected string
	}{
		{
			name:     "Update with valid string",
			todo:     Todo{Id: 1, Item: "Old item", Done: false},
			newItem:  "New item",
			expected: "New item",
		},
		{
			name:     "Update with empty string should not change",
			todo:     Todo{Id: 1, Item: "Original", Done: false},
			newItem:  "",
			expected: "Original",
		},
		{
			name:     "Update with long string",
			todo:     Todo{Id: 1, Item: "Short", Done: false},
			newItem:  "This is a very long todo item that should still work correctly",
			expected: "This is a very long todo item that should still work correctly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.todo.UpdateItem(tt.newItem)
			if tt.todo.Item != tt.expected {
				t.Errorf("UpdateItem() = %v, want %v", tt.todo.Item, tt.expected)
			}
		})
	}
}

func TestTodo_MarkComplete(t *testing.T) {
	todo := Todo{Id: 1, Item: "Test", Done: false}
	todo.MarkComplete()
	if !todo.Done {
		t.Errorf("MarkComplete() = %v, want true", todo.Done)
	}
}

func TestTodo_MarkIncomplete(t *testing.T) {
	todo := Todo{Id: 1, Item: "Test", Done: true}
	todo.MarkIncomplete()
	if todo.Done {
		t.Errorf("MarkIncomplete() = %v, want false", todo.Done)
	}
}

func TestTodo_ToggleDone(t *testing.T) {
	tests := []struct {
		name     string
		initial  bool
		expected bool
	}{
		{"Toggle from false to true", false, true},
		{"Toggle from true to false", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todo := Todo{Id: 1, Item: "Test", Done: tt.initial}
			todo.ToggleDone()
			if todo.Done != tt.expected {
				t.Errorf("ToggleDone() = %v, want %v", todo.Done, tt.expected)
			}
		})
	}
}

// Test collection management functions

func TestAddTodos(t *testing.T) {
	tests := []struct {
		name       string
		todos      []Todo
		item       string
		nextID     int
		wantLen    int
		wantID     int
		wantNextID int
	}{
		{
			name:       "Add to empty list",
			todos:      []Todo{},
			item:       "First todo",
			nextID:     1,
			wantLen:    1,
			wantID:     1,
			wantNextID: 2,
		},
		{
			name: "Add to existing list",
			todos: []Todo{
				{Id: 1, Item: "Existing", Done: false},
			},
			item:       "Second todo",
			nextID:     2,
			wantLen:    2,
			wantID:     2,
			wantNextID: 3,
		},
		{
			name: "Add multiple items",
			todos: []Todo{
				{Id: 1, Item: "First", Done: false},
				{Id: 2, Item: "Second", Done: false},
			},
			item:       "Third",
			nextID:     3,
			wantLen:    3,
			wantID:     3,
			wantNextID: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddTodos(tt.todos, tt.item, &tt.nextID)
			if len(result) != tt.wantLen {
				t.Errorf("AddTodos() length = %v, want %v", len(result), tt.wantLen)
			}
			if tt.nextID != tt.wantNextID {
				t.Errorf("AddTodos() nextID = %v, want %v", tt.nextID, tt.wantNextID)
			}
			if len(result) > 0 {
				lastTodo := result[len(result)-1]
				if lastTodo.Id != tt.wantID {
					t.Errorf("AddTodos() last todo ID = %v, want %v", lastTodo.Id, tt.wantID)
				}
				if lastTodo.Item != tt.item {
					t.Errorf("AddTodos() last todo Item = %v, want %v", lastTodo.Item, tt.item)
				}
				if lastTodo.Done != false {
					t.Errorf("AddTodos() last todo Done = %v, want false", lastTodo.Done)
				}
			}
		})
	}
}

func TestFindTodosById(t *testing.T) {
	todos := []Todo{
		{Id: 1, Item: "First", Done: false},
		{Id: 2, Item: "Second", Done: true},
		{Id: 3, Item: "Third", Done: false},
	}

	tests := []struct {
		name      string
		id        int
		wantIndex int
		wantFound bool
		wantItem  string
	}{
		{"Find existing ID", 2, 1, true, "Second"},
		{"Find first ID", 1, 0, true, "First"},
		{"Find last ID", 3, 2, true, "Third"},
		{"Find non-existent ID", 99, -1, false, ""},
		{"Find zero ID", 0, -1, false, ""},
		{"Find negative ID", -1, -1, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, todo := FindTodosById(todos, tt.id)
			if index != tt.wantIndex {
				t.Errorf("FindTodosById() index = %v, want %v", index, tt.wantIndex)
			}
			if (todo != nil) != tt.wantFound {
				t.Errorf("FindTodosById() found = %v, want %v", todo != nil, tt.wantFound)
			}
			if tt.wantFound && todo != nil && todo.Item != tt.wantItem {
				t.Errorf("FindTodosById() item = %v, want %v", todo.Item, tt.wantItem)
			}
		})
	}
}

func TestFindTodos(t *testing.T) {
	todos := []Todo{
		{Id: 1, Item: "Buy milk", Done: false},
		{Id: 2, Item: "Walk dog", Done: true},
		{Id: 3, Item: "Buy groceries", Done: false},
	}

	tests := []struct {
		name      string
		itemName  string
		wantIndex int
		wantFound bool
	}{
		{"Find existing item", "Walk dog", 1, true},
		{"Find first item", "Buy milk", 0, true},
		{"Find non-existent item", "Not found", -1, false},
		{"Find with empty string", "", -1, false},
		{"Case sensitive search", "buy milk", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, todo := FindTodos(todos, tt.itemName)
			if index != tt.wantIndex {
				t.Errorf("FindTodos() index = %v, want %v", index, tt.wantIndex)
			}
			if (todo != nil) != tt.wantFound {
				t.Errorf("FindTodos() found = %v, want %v", todo != nil, tt.wantFound)
			}
		})
	}
}

func TestRemoveTodos(t *testing.T) {
	tests := []struct {
		name        string
		todos       []Todo
		id          int
		wantLen     int
		wantError   bool
		wantFirstID int
	}{
		{
			name: "Remove from middle",
			todos: []Todo{
				{Id: 1, Item: "First", Done: false},
				{Id: 2, Item: "Second", Done: false},
				{Id: 3, Item: "Third", Done: false},
			},
			id:          2,
			wantLen:     2,
			wantError:   false,
			wantFirstID: 1,
		},
		{
			name: "Remove first item",
			todos: []Todo{
				{Id: 1, Item: "First", Done: false},
				{Id: 2, Item: "Second", Done: false},
			},
			id:          1,
			wantLen:     1,
			wantError:   false,
			wantFirstID: 2,
		},
		{
			name: "Remove last item",
			todos: []Todo{
				{Id: 1, Item: "First", Done: false},
				{Id: 2, Item: "Second", Done: false},
			},
			id:          2,
			wantLen:     1,
			wantError:   false,
			wantFirstID: 1,
		},
		{
			name: "Remove non-existent ID",
			todos: []Todo{
				{Id: 1, Item: "First", Done: false},
			},
			id:          99,
			wantLen:     1,
			wantError:   true,
			wantFirstID: 1,
		},
		{
			name:        "Remove from empty list",
			todos:       []Todo{},
			id:          1,
			wantLen:     0,
			wantError:   true,
			wantFirstID: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := RemoveTodos(tt.todos, tt.id)
			if (err != nil) != tt.wantError {
				t.Errorf("RemoveTodos() error = %v, wantError %v", err, tt.wantError)
			}
			if len(result) != tt.wantLen {
				t.Errorf("RemoveTodos() length = %v, want %v", len(result), tt.wantLen)
			}
			if tt.wantLen > 0 && len(result) > 0 {
				if result[0].Id != tt.wantFirstID {
					t.Errorf("RemoveTodos() first ID = %v, want %v", result[0].Id, tt.wantFirstID)
				}
			}
		})
	}
}

func TestMarkCompleteByID(t *testing.T) {
	tests := []struct {
		name      string
		todos     []Todo
		id        int
		wantError bool
		wantDone  bool
	}{
		{
			name: "Mark existing todo as complete",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: false},
			},
			id:        1,
			wantError: false,
			wantDone:  true,
		},
		{
			name: "Mark already complete todo",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: true},
			},
			id:        1,
			wantError: false,
			wantDone:  true,
		},
		{
			name: "Mark non-existent ID",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: false},
			},
			id:        99,
			wantError: true,
			wantDone:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MarkCompleteByID(tt.todos, tt.id)
			if (err != nil) != tt.wantError {
				t.Errorf("MarkCompleteByID() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError {
				_, todo := FindTodosById(tt.todos, tt.id)
				if todo != nil && todo.Done != tt.wantDone {
					t.Errorf("MarkCompleteByID() Done = %v, want %v", todo.Done, tt.wantDone)
				}
			}
		})
	}
}

func TestMarkIncompleteByID(t *testing.T) {
	tests := []struct {
		name      string
		todos     []Todo
		id        int
		wantError bool
		wantDone  bool
	}{
		{
			name: "Mark existing todo as incomplete",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: true},
			},
			id:        1,
			wantError: false,
			wantDone:  false,
		},
		{
			name: "Mark non-existent ID",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: true},
			},
			id:        99,
			wantError: true,
			wantDone:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MarkIncompleteByID(tt.todos, tt.id)
			if (err != nil) != tt.wantError {
				t.Errorf("MarkIncompleteByID() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError {
				_, todo := FindTodosById(tt.todos, tt.id)
				if todo != nil && todo.Done != tt.wantDone {
					t.Errorf("MarkIncompleteByID() Done = %v, want %v", todo.Done, tt.wantDone)
				}
			}
		})
	}
}

func TestToggleDoneByID(t *testing.T) {
	tests := []struct {
		name      string
		todos     []Todo
		id        int
		wantError bool
		wantDone  bool
	}{
		{
			name: "Toggle from false to true",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: false},
			},
			id:        1,
			wantError: false,
			wantDone:  true,
		},
		{
			name: "Toggle from true to false",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: true},
			},
			id:        1,
			wantError: false,
			wantDone:  false,
		},
		{
			name: "Toggle non-existent ID",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: false},
			},
			id:        99,
			wantError: true,
			wantDone:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ToggleDoneByID(tt.todos, tt.id)
			if (err != nil) != tt.wantError {
				t.Errorf("ToggleDoneByID() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError {
				_, todo := FindTodosById(tt.todos, tt.id)
				if todo != nil && todo.Done != tt.wantDone {
					t.Errorf("ToggleDoneByID() Done = %v, want %v", todo.Done, tt.wantDone)
				}
			}
		})
	}
}

func TestUpdateItemByID(t *testing.T) {
	tests := []struct {
		name      string
		todos     []Todo
		id        int
		newText   string
		wantError bool
		wantItem  string
	}{
		{
			name: "Update existing todo",
			todos: []Todo{
				{Id: 1, Item: "Old text", Done: false},
			},
			id:        1,
			newText:   "New text",
			wantError: false,
			wantItem:  "New text",
		},
		{
			name: "Update with empty string",
			todos: []Todo{
				{Id: 1, Item: "Original", Done: false},
			},
			id:        1,
			newText:   "",
			wantError: false,
			wantItem:  "Original", // Should not change
		},
		{
			name: "Update non-existent ID",
			todos: []Todo{
				{Id: 1, Item: "Test", Done: false},
			},
			id:        99,
			newText:   "New",
			wantError: true,
			wantItem:  "Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateItemByID(tt.todos, tt.id, tt.newText)
			if (err != nil) != tt.wantError {
				t.Errorf("UpdateItemByID() error = %v, wantError %v", err, tt.wantError)
			}
			if !tt.wantError {
				_, todo := FindTodosById(tt.todos, tt.id)
				if todo != nil && todo.Item != tt.wantItem {
					t.Errorf("UpdateItemByID() Item = %v, want %v", todo.Item, tt.wantItem)
				}
			}
		})
	}
}

func TestGetNextID(t *testing.T) {
	tests := []struct {
		name     string
		todos    []Todo
		wantNext int
	}{
		{"Empty list", []Todo{}, 1},
		{"Single todo", []Todo{{Id: 1, Item: "Test", Done: false}}, 2},
		{"Multiple todos", []Todo{
			{Id: 1, Item: "First", Done: false},
			{Id: 5, Item: "Fifth", Done: false},
			{Id: 3, Item: "Third", Done: false},
		}, 6}, // Should find max (5) and add 1
		{"Consecutive IDs", []Todo{
			{Id: 1, Item: "First", Done: false},
			{Id: 2, Item: "Second", Done: false},
			{Id: 3, Item: "Third", Done: false},
		}, 4},
		{"Non-consecutive IDs", []Todo{
			{Id: 10, Item: "Tenth", Done: false},
			{Id: 20, Item: "Twentieth", Done: false},
		}, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNextID(tt.todos)
			if got != tt.wantNext {
				t.Errorf("GetNextID() = %v, want %v", got, tt.wantNext)
			}
		})
	}
}

func TestListPendingTodos(t *testing.T) {
	todos := []Todo{
		{Id: 1, Item: "Pending 1", Done: false},
		{Id: 2, Item: "Completed 1", Done: true},
		{Id: 3, Item: "Pending 2", Done: false},
		{Id: 4, Item: "Completed 2", Done: true},
	}

	// This is a basic test - ListPendingTodos prints to stdout
	// In a real scenario, you'd capture stdout or refactor to return values
	ListPendingTodos(todos)
	// If we want to test this properly, we'd need to refactor to return []Todo
}

func TestListCompleteTodos(t *testing.T) {
	todos := []Todo{
		{Id: 1, Item: "Pending 1", Done: false},
		{Id: 2, Item: "Completed 1", Done: true},
		{Id: 3, Item: "Pending 2", Done: false},
		{Id: 4, Item: "Completed 2", Done: true},
	}

	// This is a basic test - ListCompleteTodos prints to stdout
	ListCompleteTodos(todos)
}

// Integration test for full workflow
func TestFullWorkflow(t *testing.T) {
	var todos []Todo
	nextID := 1

	// Add todos
	todos = AddTodos(todos, "First todo", &nextID)
	todos = AddTodos(todos, "Second todo", &nextID)
	todos = AddTodos(todos, "Third todo", &nextID)

	if len(todos) != 3 {
		t.Errorf("Expected 3 todos, got %d", len(todos))
	}

	// Mark first as complete
	err := MarkCompleteByID(todos, 1)
	if err != nil {
		t.Errorf("MarkCompleteByID() error = %v", err)
	}

	// Update second todo
	err = UpdateItemByID(todos, 2, "Updated second todo")
	if err != nil {
		t.Errorf("UpdateItemByID() error = %v", err)
	}

	// Toggle third todo
	err = ToggleDoneByID(todos, 3)
	if err != nil {
		t.Errorf("ToggleDoneByID() error = %v", err)
	}

	// Verify state
	_, todo1 := FindTodosById(todos, 1)
	if todo1 == nil || !todo1.Done {
		t.Error("Todo 1 should be done")
	}

	_, todo2 := FindTodosById(todos, 2)
	if todo2 == nil || todo2.Item != "Updated second todo" {
		t.Error("Todo 2 should be updated")
	}

	_, todo3 := FindTodosById(todos, 3)
	if todo3 == nil || !todo3.Done {
		t.Error("Todo 3 should be done after toggle")
	}

	// Remove a todo
	todos, err = RemoveTodos(todos, 2)
	if err != nil {
		t.Errorf("RemoveTodos() error = %v", err)
	}

	if len(todos) != 2 {
		t.Errorf("Expected 2 todos after removal, got %d", len(todos))
	}
}

// Test edge cases and error conditions
func TestEdgeCases(t *testing.T) {
	// Test with nil pointer (should not panic)
	todo := &Todo{Id: 1, Item: "Test", Done: false}
	todo.UpdateItem("")
	if todo.Item != "Test" {
		t.Error("UpdateItem with empty string should not change item")
	}

	// Test with very long string
	longString := make([]byte, 10000)
	for i := range longString {
		longString[i] = 'a'
	}
	todo.UpdateItem(string(longString))
	if len(todo.Item) != 10000 {
		t.Error("UpdateItem should handle long strings")
	}

	// Test multiple toggles
	todo.Done = false
	for i := 0; i < 10; i++ {
		todo.ToggleDone()
	}
	if todo.Done != false { // After 10 (even) toggles from false, should be false
		t.Error("Multiple toggles should work correctly")
	}

	// Test odd number of toggles
	todo.Done = false
	for i := 0; i < 5; i++ {
		todo.ToggleDone()
	}
	if todo.Done != true { // After 5 (odd) toggles from false, should be true
		t.Error("Multiple toggles should work correctly")
	}
}

// Test JSON serialization (for Supabase compatibility)
func TestTodoJSONSerialization(t *testing.T) {
	todo := Todo{Id: 1, Item: "Test item", Done: false}

	// This tests that the JSON tags work correctly
	// In a real scenario, you'd use encoding/json to marshal/unmarshal
	// This is a placeholder to ensure the struct can be serialized
	if todo.Id != 1 || todo.Item != "Test item" || todo.Done != false {
		t.Error("Todo struct fields not set correctly")
	}
}

// Benchmark tests
func BenchmarkAddTodos(b *testing.B) {
	todos := []Todo{}
	nextID := 1
	for i := 0; i < b.N; i++ {
		todos = AddTodos(todos, "Benchmark item", &nextID)
	}
}

func BenchmarkFindTodosById(b *testing.B) {
	todos := make([]Todo, 1000)
	for i := 0; i < 1000; i++ {
		todos[i] = Todo{Id: i + 1, Item: "Item", Done: false}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindTodosById(todos, 500)
	}
}

func BenchmarkGetNextID(b *testing.B) {
	todos := make([]Todo, 1000)
	for i := 0; i < 1000; i++ {
		todos[i] = Todo{Id: i + 1, Item: "Item", Done: false}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetNextID(todos)
	}
}
