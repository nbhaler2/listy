# How to Run Listy - Step by Step Guide

## Prerequisites

1. Make sure you have a `.env` file in the root directory with:
   ```
   SUPABASE_URL=your-supabase-url
   SUPABASE_KEY=your-supabase-key
   ```

2. Make sure Go is installed and working

## Step 1: Start the API Server

Open **Terminal Window 1** (keep this running):

```bash
cd "/Users/namitabhalerao/Go Tutorial/Todolist/api"
go run main.go
```

You should see:
```
🚀 Server starting on port 8080
📡 API endpoints available at http://localhost:8080/api
❤️  Health check: http://localhost:8080/api/health
```

**Keep this terminal window open!** The API server needs to keep running.

## Step 2: Test the API (Optional)

Open **Terminal Window 2** (new window):

```bash
# Test health check
curl http://localhost:8080/api/health

# Should return: {"service":"listy-api","status":"healthy"}
```

## Step 3: Use the CLI

In **Terminal Window 2** (or a new window), run CLI commands:

```bash
cd "/Users/namitabhalerao/Go Tutorial/Todolist/cli"

# List all todos
go run main.go api_client.go config.go list

# Add a todo
go run main.go api_client.go config.go add "Buy groceries"

# Add another
go run main.go api_client.go config.go add "Walk the dog"

# List again to see your todos
go run main.go api_client.go config.go list

# Mark one as complete
go run main.go api_client.go config.go complete 1

# List pending todos
go run main.go api_client.go config.go pending

# List completed todos
go run main.go api_client.go config.go completed

# Toggle a todo
go run main.go api_client.go config.go toggle 2

# Update a todo
go run main.go api_client.go config.go update 1 "Buy groceries and milk"

# Remove a todo
go run main.go api_client.go config.go remove 2

# See final list
go run main.go api_client.go config.go list
```

## Quick Test Sequence

Run these commands in order to see it working:

```bash
# 1. Start API (Terminal 1)
cd api && go run main.go

# 2. In Terminal 2, test CLI
cd cli
go run main.go api_client.go config.go list
go run main.go api_client.go config.go add "Test from CLI"
go run main.go api_client.go config.go list
go run main.go api_client.go config.go complete 1
go run main.go api_client.go config.go list
```

## Troubleshooting

### Error: "API server is not running"
- Make sure API server is running in Terminal 1
- Check if port 8080 is available
- Verify API is accessible: `curl http://localhost:8080/api/health`

### Error: "Failed to connect to API"
- Check API server is running
- Verify `.env` file exists with Supabase credentials
- Check API logs in Terminal 1

### Port already in use
- Change port: `PORT=8081 go run main.go` in api directory
- Update CLI: `LISTY_API_URL=http://localhost:8081 go run main.go ...`

## What You Should See

### API Server (Terminal 1):
```
🚀 Server starting on port 8080
📡 API endpoints available at http://localhost:8080/api
❤️  Health check: http://localhost:8080/api/health
```

### CLI Commands (Terminal 2):
```
$ go run main.go api_client.go config.go list
{1 Test Supabase integration true}
{2 Test incremental update false}
{3 Updated CLI Test Todo false}

$ go run main.go api_client.go config.go add "New todo"
Added New todo (Id: 4)

$ go run main.go api_client.go config.go list
{1 Test Supabase integration true}
{2 Test incremental update false}
{3 Updated CLI Test Todo false}
{4 New todo false}
```

## Pro Tips

1. **Keep API running**: Leave Terminal 1 open with API server
2. **Use multiple terminals**: One for API, one for CLI commands
3. **Check API health**: `curl http://localhost:8080/api/health` anytime
4. **View API logs**: Check Terminal 1 for request logs

