package main

import (
	"fmt"
	"log"
	"os"

	"listy-api/database"
	"listy-api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Supabase
	err := database.InitSupabase()
	if err != nil {
		log.Fatalf("Failed to initialize Supabase: %v", err)
	}

	// Set up Gin router
	r := gin.Default()

	// CORS middleware - allow requests from Next.js frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Next.js default port
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/api/health", handlers.HealthCheck)

	// Todo routes
	api := r.Group("/api/todos")
	{
		api.GET("", handlers.GetTodos)                    // GET /api/todos
		api.GET("/pending", handlers.GetPendingTodos)     // GET /api/todos/pending
		api.GET("/completed", handlers.GetCompletedTodos) // GET /api/todos/completed
		api.GET("/:id", handlers.GetTodoByID)             // GET /api/todos/:id
		api.POST("", handlers.CreateTodo)                 // POST /api/todos
		api.PUT("/:id", handlers.UpdateTodo)              // PUT /api/todos/:id
		api.PATCH("/:id/toggle", handlers.ToggleTodo)     // PATCH /api/todos/:id/toggle
		api.DELETE("/:id", handlers.DeleteTodo)           // DELETE /api/todos/:id
	}

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 Server starting on port %s\n", port)
	fmt.Printf("📡 API endpoints available at http://localhost:%s/api\n", port)
	fmt.Printf("❤️  Health check: http://localhost:%s/api/health\n", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
