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

// GenerateSubtaskBreakdown uses OpenAI to generate subtasks for a specific task
// It intelligently determines if the task can be broken down into subtasks
func GenerateSubtaskBreakdown(task string) ([]models.AITask, error) {
	// Get API key from environment
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	// Initialize OpenAI client
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	// Create a smarter prompt for subtask breakdown
	prompt := fmt.Sprintf(`You are a JSON-only response assistant. Analyze the task and return ONLY a valid JSON array with no additional text.

Task: "%s"

RULES:
1. If the task is simple/atomic (e.g., "Buy milk", "Call John"), return: []
2. If the task can be broken down, return 3-6 subtasks as: [{"text": "Subtask 1"}, {"text": "Subtask 2"}]
3. Return ONLY the JSON array, no explanations, no markdown, no other text

Examples:
"Buy groceries" → []
"Learn Go programming" → [{"text": "Install Go compiler"}, {"text": "Read Go basics"}, {"text": "Write first program"}]
"Call mom" → []

Return ONLY the JSON array:`, task)

	// Make API call with system message to enforce JSON-only response
	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a JSON-only response assistant. Always return valid JSON arrays with no additional text, explanations, or markdown formatting.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.3, // Lower temperature for more consistent JSON output
			MaxTokens:   500,
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

	// Extract JSON array from response (AI might add explanatory text)
	// Find the first '[' and last ']' to extract the JSON array
	firstBracket := strings.Index(content, "[")
	lastBracket := strings.LastIndex(content, "]")

	if firstBracket == -1 || lastBracket == -1 || firstBracket >= lastBracket {
		// No array found, try to parse the whole content
		// If it's just "[]", that's valid
		if strings.TrimSpace(content) == "[]" {
			return []models.AITask{}, nil
		}
		return nil, fmt.Errorf("no valid JSON array found in response. Response was: %s", content)
	}

	// Extract just the JSON array part
	jsonContent := content[firstBracket : lastBracket+1]

	// Parse JSON response
	var tasks []models.AITask
	if err := json.Unmarshal([]byte(jsonContent), &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse AI response: %v. Extracted JSON: %s, Full response: %s", err, jsonContent, content)
	}

	// Empty array is valid - means task cannot be broken down
	return tasks, nil
}
