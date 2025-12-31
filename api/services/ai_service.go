package services

import (
	"context"
	"encoding/json"
	"fmt"
	"listy-api/models"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// GenerateTaskBreakdown uses OpenAI to generate a breakdown of tasks for a given goal
func GenerateTaskBreakdown(goal string) ([]models.AITask, error) {
	// Get API key from environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	// Initialize OpenAI client
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	// Create a prompt for task breakdown
	prompt := fmt.Sprintf(`You are a helpful task breakdown assistant. Given a goal or task, break it down into 5-8 actionable, specific subtasks.

Goal: "%s"

Generate a JSON array of tasks. Each task should have:
- text: A clear, actionable task description

Return ONLY a valid JSON array, no other text. Example format:
[
  {"text": "Install Go on your system"},
  {"text": "Read Go documentation basics"},
  {"text": "Write your first Hello World program"}
]`, goal)

	// Make API call
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.7,
			MaxTokens:   1000,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("OpenAI API error: %v", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	// Extract the response content
	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	// Clean up the response - remove markdown code blocks if present
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	// Parse JSON response
	var tasks []models.AITask
	if err := json.Unmarshal([]byte(content), &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v. Response was: %s", err, content)
	}

	// Validate tasks
	if len(tasks) == 0 {
		return nil, fmt.Errorf("no tasks generated")
	}

	return tasks, nil
}
