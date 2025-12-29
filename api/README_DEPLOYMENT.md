# API Server Deployment

This API server can be deployed to Railway, Fly.io, Render, or any platform that supports Go/Docker.

## Quick Deploy to Railway

1. **Push to GitHub:**
   ```bash
   git add .
   git commit -m "Ready for deployment"
   git push origin main
   ```

2. **In Railway:**
   - New Project → Deploy from GitHub
   - Select repository
   - Set root directory: `api`
   - Add environment variables (see below)
   - Deploy!

3. **Environment Variables:**
   ```
   SUPABASE_URL=your-supabase-url
   SUPABASE_KEY=your-supabase-key
   PORT=8080
   ALLOWED_ORIGIN=https://your-vercel-app.vercel.app
   ```

## Build Locally (Test)

```bash
# Build
go build -o api-server ./main.go

# Run
./api-server
```

## Docker Build (Test)

```bash
# Build image
docker build -t listy-api .

# Run container
docker run -p 8080:8080 \
  -e SUPABASE_URL=your-url \
  -e SUPABASE_KEY=your-key \
  listy-api
```

## Health Check

After deployment, test:
```bash
curl https://your-api.railway.app/api/health
```

Should return: `{"service":"listy-api","status":"healthy"}`

## API Endpoints

- `GET /api/health` - Health check
- `GET /api/todos` - List all todos
- `GET /api/todos/pending` - List pending todos
- `GET /api/todos/completed` - List completed todos
- `GET /api/todos/:id` - Get todo by ID
- `POST /api/todos` - Create todo
- `PUT /api/todos/:id` - Update todo
- `PATCH /api/todos/:id/toggle` - Toggle todo status
- `DELETE /api/todos/:id` - Delete todo

