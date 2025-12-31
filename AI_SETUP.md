# AI Task Generator Setup 🤖

## Overview

The AI Task Generator feature uses OpenAI's GPT-3.5 to automatically break down goals into actionable todo tasks with metadata (priority, estimated time, category).

## Setup Instructions

### 1. Get OpenAI API Key

1. Go to [OpenAI Platform](https://platform.openai.com/)
2. Sign up or log in
3. Navigate to [API Keys](https://platform.openai.com/api-keys)
4. Click "Create new secret key"
5. Copy the API key (you won't be able to see it again!)

### 2. Set Environment Variable

#### For Local Development (API Server)

Add to your `.env` file in the `api/` directory:

```bash
OPENAI_API_KEY=sk-your-api-key-here
```

Or export it in your terminal:

```bash
export OPENAI_API_KEY=sk-your-api-key-here
```

#### For Production (Railway)

1. Go to your Railway project dashboard
2. Navigate to "Variables" tab
3. Add new variable:
   - **Key**: `OPENAI_API_KEY`
   - **Value**: Your OpenAI API key
4. Save and redeploy

### 3. Restart API Server

After setting the environment variable, restart your API server:

```bash
cd api
go run main.go
```

## How It Works

### User Flow

1. **User enters goal**: "learning Go"
2. **AI generates breakdown**: Creates 5-8 actionable tasks with:
   - Task description
   - Priority (high/medium/low)
   - Estimated time
   - Category/tag
3. **User reviews & edits**: Can modify, delete, or add custom tasks
4. **Create all tasks**: All tasks are added to the todo list
5. **Track progress**: User can check off tasks as they complete them

### API Endpoints

#### Generate Task Breakdown
```
POST /api/todos/ai/breakdown
Body: { "goal": "learning Go" }
Response: {
  "success": true,
  "goal": "learning Go",
  "suggested_tasks": [
    {
      "text": "Install Go on your system",
      "priority": "high",
      "estimated_time": "15 minutes",
      "category": "setup"
    },
    ...
  ]
}
```

#### Create AI Tasks
```
POST /api/todos/ai/create
Body: {
  "tasks": [
    {
      "text": "Install Go on your system",
      "priority": "high",
      "estimated_time": "15 minutes",
      "category": "setup"
    },
    ...
  ]
}
Response: {
  "success": true,
  "message": "Created 5 task(s)",
  "data": [/* created todos */]
}
```

## Features

✅ **Smart Task Breakdown** - AI generates 5-8 actionable tasks  
✅ **Metadata Included** - Priority, estimated time, category  
✅ **Editable Before Creation** - Users can modify AI suggestions  
✅ **Add Custom Tasks** - Users can add their own tasks  
✅ **Delete Unwanted Tasks** - Remove tasks before creating  
✅ **Batch Creation** - Create all tasks at once  

## Cost Considerations

- Uses OpenAI GPT-3.5 Turbo (cost-effective)
- ~$0.001-0.002 per request (varies by response length)
- Typical breakdown generates ~5-8 tasks per request

## Troubleshooting

### Error: "OPENAI_API_KEY environment variable not set"
- Make sure you've set the environment variable
- Restart the API server after setting it
- Check that the variable name is exactly `OPENAI_API_KEY`

### Error: "OpenAI API error"
- Check your API key is valid
- Verify you have credits in your OpenAI account
- Check your internet connection

### Tasks not generating
- Check API server logs for errors
- Verify OpenAI API key has proper permissions
- Try a simpler goal first (e.g., "learn Python")

## Testing

Test the AI feature:

```bash
# Test breakdown endpoint
curl -X POST http://localhost:8080/api/todos/ai/breakdown \
  -H "Content-Type: application/json" \
  -d '{"goal": "learning Go"}'
```

## Next Steps

- ✅ AI integration complete
- 🎯 Consider adding task templates
- 🎯 Add regeneration feature
- 🎯 Smart suggestions based on existing todos


