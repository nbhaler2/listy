package handlers

import (
	"fmt"
	"listy-api/models"
	"listy-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateTaskBreakdown handles POST /api/todos/ai/breakdown
func GenerateTaskBreakdown(c *gin.Context) {
	var req models.AITaskBreakdownRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate goal is not empty
	if req.Goal == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "goal is required"})
		return
	}

	// Generate task breakdown using AI
	tasks, err := services.GenerateTaskBreakdown(req.Goal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to generate task breakdown",
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, models.AITaskBreakdownResponse{
		Success:        true,
		Goal:           req.Goal,
		SuggestedTasks: tasks,
		Message:        "Task breakdown generated successfully",
	})
}

// CreateAITasks handles POST /api/todos/ai/create
// Creates multiple todos from AI-generated tasks
func CreateAITasks(c *gin.Context) {
	var req models.CreateAITasksRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.Tasks) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "at least one task is required"})
		return
	}

	// Create todos from AI tasks
	var createdTodos []models.Todo
	var errors []string

	for _, aiTask := range req.Tasks {
		// Use clean task text only
		taskText := aiTask.Text

		todo, err := services.CreateTodo(taskText, req.ListId)
		if err != nil {
			errors = append(errors, "Failed to create task: "+aiTask.Text+" - "+err.Error())
			continue
		}
		createdTodos = append(createdTodos, *todo)
	}

	if len(createdTodos) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create any tasks",
			"details": errors,
		})
		return
	}

	response := gin.H{
		"success": true,
		"message": fmt.Sprintf("Created %d task(s)", len(createdTodos)),
		"data":    createdTodos,
	}

	if len(errors) > 0 {
		response["warnings"] = errors
	}

	c.JSON(http.StatusCreated, response)
}
