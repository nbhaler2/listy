package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// APIClient handles all API communication
type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewAPIClient creates a new API client
func NewAPIClient(baseURL string) *APIClient {
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return &APIClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// APIResponse represents the standard API response format
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

// Todo represents a todo item (matches API model)
type Todo struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

// CreateTodoRequest represents the request for creating a todo
type CreateTodoRequest struct {
	Item string `json:"item"`
}

// UpdateTodoRequest represents the request for updating a todo
type UpdateTodoRequest struct {
	Item *string `json:"item,omitempty"`
	Done *bool   `json:"done,omitempty"`
}

// GetTodos fetches all todos from the API
func (c *APIClient) GetTodos() ([]Todo, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/api/todos")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	// Convert data to []Todo
	dataBytes, _ := json.Marshal(apiResp.Data)
	var todos []Todo
	if err := json.Unmarshal(dataBytes, &todos); err != nil {
		return nil, fmt.Errorf("failed to parse todos: %v", err)
	}

	return todos, nil
}

// GetPendingTodos fetches pending todos from the API
func (c *APIClient) GetPendingTodos() ([]Todo, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/api/todos/pending")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	dataBytes, _ := json.Marshal(apiResp.Data)
	var todos []Todo
	if err := json.Unmarshal(dataBytes, &todos); err != nil {
		return nil, fmt.Errorf("failed to parse todos: %v", err)
	}

	return todos, nil
}

// GetCompletedTodos fetches completed todos from the API
func (c *APIClient) GetCompletedTodos() ([]Todo, error) {
	resp, err := c.httpClient.Get(c.baseURL + "/api/todos/completed")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	dataBytes, _ := json.Marshal(apiResp.Data)
	var todos []Todo
	if err := json.Unmarshal(dataBytes, &todos); err != nil {
		return nil, fmt.Errorf("failed to parse todos: %v", err)
	}

	return todos, nil
}

// CreateTodo creates a new todo via the API
func (c *APIClient) CreateTodo(item string) (*Todo, error) {
	reqBody := CreateTodoRequest{Item: item}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := c.httpClient.Post(c.baseURL+"/api/todos", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	dataBytes, _ := json.Marshal(apiResp.Data)
	var todo Todo
	if err := json.Unmarshal(dataBytes, &todo); err != nil {
		return nil, fmt.Errorf("failed to parse todo: %v", err)
	}

	return &todo, nil
}

// UpdateTodo updates a todo via the API
func (c *APIClient) UpdateTodo(id int, req UpdateTodoRequest) (*Todo, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	reqHTTP, err := http.NewRequest("PUT", c.baseURL+"/api/todos/"+strconv.Itoa(id), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	reqHTTP.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(reqHTTP)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	dataBytes, _ := json.Marshal(apiResp.Data)
	var todo Todo
	if err := json.Unmarshal(dataBytes, &todo); err != nil {
		return nil, fmt.Errorf("failed to parse todo: %v", err)
	}

	return &todo, nil
}

// DeleteTodo deletes a todo via the API
func (c *APIClient) DeleteTodo(id int) error {
	reqHTTP, err := http.NewRequest("DELETE", c.baseURL+"/api/todos/"+strconv.Itoa(id), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(reqHTTP)
	if err != nil {
		return fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error: %s", string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	if !apiResp.Success {
		return fmt.Errorf("API error: %s", apiResp.Error)
	}

	return nil
}

// ToggleTodo toggles a todo's done status via the API
func (c *APIClient) ToggleTodo(id int) (*Todo, error) {
	reqHTTP, err := http.NewRequest("PATCH", c.baseURL+"/api/todos/"+strconv.Itoa(id)+"/toggle", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(reqHTTP)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if !apiResp.Success {
		return nil, fmt.Errorf("API error: %s", apiResp.Error)
	}

	dataBytes, _ := json.Marshal(apiResp.Data)
	var todo Todo
	if err := json.Unmarshal(dataBytes, &todo); err != nil {
		return nil, fmt.Errorf("failed to parse todo: %v", err)
	}

	return &todo, nil
}

// CheckHealth checks if the API is available
func (c *APIClient) CheckHealth() error {
	resp, err := c.httpClient.Get(c.baseURL + "/api/health")
	if err != nil {
		return fmt.Errorf("API server is not running: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API server returned status: %d", resp.StatusCode)
	}

	return nil
}
