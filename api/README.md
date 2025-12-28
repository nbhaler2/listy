# Listy API Server

Go REST API server for the Listy todo list application.

## Setup

1. Make sure you have a `.env` file in the parent directory with:
   ```
   SUPABASE_URL=your-supabase-url
   SUPABASE_KEY=your-supabase-key
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

   Or build and run:
   ```bash
   go build -o api-server main.go
   ./api-server
   ```

The server will start on port 8080 (or PORT environment variable).

## API Endpoints

### Health Check
- `GET /api/health` - Check if API is running

### Todos
- `GET /api/todos` - Get all todos
- `GET /api/todos/pending` - Get pending todos
- `GET /api/todos/completed` - Get completed todos
- `GET /api/todos/:id` - Get todo by ID
- `POST /api/todos` - Create a new todo
- `PUT /api/todos/:id` - Update a todo
- `PATCH /api/todos/:id/toggle` - Toggle todo status
- `DELETE /api/todos/:id` - Delete a todo

## Request/Response Examples

### Create Todo
```bash
POST /api/todos
Content-Type: application/json

{
  "item": "Buy groceries"
}
```

Response:
```json
{
  "success": true,
  "data": {
    "id": 1,
    "item": "Buy groceries",
    "done": false
  }
}
```

### Update Todo
```bash
PUT /api/todos/1
Content-Type: application/json

{
  "item": "Buy groceries and milk",
  "done": true
}
```

### Toggle Todo
```bash
PATCH /api/todos/1/toggle
```

## Testing

Test with curl:
```bash
# Health check
curl http://localhost:8080/api/health

# Get all todos
curl http://localhost:8080/api/todos

# Create todo
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"item": "Test todo"}'

# Update todo
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"done": true}'

# Delete todo
curl -X DELETE http://localhost:8080/api/todos/1
```

## Project Structure

```
api/
├── main.go              # Server entry point
├── handlers/            # HTTP handlers
│   ├── todo_handler.go
│   └── health_handler.go
├── services/            # Business logic
│   └── todo_service.go
├── models/              # Data models
│   └── todo.go
└── database/            # Database layer
    └── supabase.go
```

