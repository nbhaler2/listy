package handlers

import (
	"net/http"
	"strconv"

	"listy-api/models"
	"listy-api/services"

	"github.com/gin-gonic/gin"
)

// GetTodos handles GET /api/todos
func GetTodos(c *gin.Context) {
	todos, err := services.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": todos})
}

// GetPendingTodos handles GET /api/todos/pending
func GetPendingTodos(c *gin.Context) {
	todos, err := services.GetPendingTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": todos})
}

// GetCompletedTodos handles GET /api/todos/completed
func GetCompletedTodos(c *gin.Context) {
	todos, err := services.GetCompletedTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": todos})
}

// GetTodoByID handles GET /api/todos/:id
func GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := services.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": todo})
}

// CreateTodo handles POST /api/todos
func CreateTodo(c *gin.Context) {
	var req models.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := services.CreateTodo(req.Item, req.ListId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": true, "data": todo})
}

// GetTodosByList handles GET /api/todos/list/:listId
// If listId is "main" or empty, returns main list todos (list_id is NULL)
func GetTodosByList(c *gin.Context) {
	listIdParam := c.Param("listId")
	var listId *string
	
	if listIdParam != "" && listIdParam != "main" {
		listId = &listIdParam
	}

	todos, err := services.GetTodosByListId(listId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": todos})
}

// GetAllLists handles GET /api/lists
func GetAllLists(c *gin.Context) {
	listIds, err := services.GetAllListIds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": listIds})
}

// UpdateTodo handles PUT /api/todos/:id
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req models.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := services.UpdateTodo(id, req)
	if err != nil {
		if err.Error() == "todo with ID "+strconv.Itoa(id)+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": todo})
}

// DeleteTodo handles DELETE /api/todos/:id
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = services.DeleteTodo(id)
	if err != nil {
		if err.Error() == "todo with ID "+strconv.Itoa(id)+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Todo deleted successfully"})
}

// ToggleTodo handles PATCH /api/todos/:id/toggle
func ToggleTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	todo, err := services.ToggleTodo(id)
	if err != nil {
		if err.Error() == "todo with ID "+strconv.Itoa(id)+" not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": todo})
}
