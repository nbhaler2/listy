# Listy Web UI

Next.js frontend for the Listy todo list application.

## Setup

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Configure API URL (optional):**
   Create `.env.local` file:
   ```
   NEXT_PUBLIC_API_URL=http://localhost:8080
   ```
   Default is `http://localhost:8080` if not set.

3. **Make sure Go API is running:**
   ```bash
   cd ../api
   go run main.go
   ```

## Development

Run the development server:

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) in your browser.

## Features

- ✅ Add todos
- ✅ List all todos
- ✅ Filter by status (All/Pending/Completed)
- ✅ Mark todos as complete/incomplete
- ✅ Toggle todo status
- ✅ Edit todos (click on todo text)
- ✅ Delete todos
- ✅ Real-time updates
- ✅ Beautiful, modern UI
- ✅ Responsive design

## Project Structure

```
web/
├── app/
│   ├── page.tsx          # Main page
│   ├── layout.tsx        # Root layout
│   └── globals.css       # Global styles
├── components/
│   ├── AddTodoForm.tsx   # Add todo form
│   ├── TodoItem.tsx      # Individual todo item
│   └── TodoList.tsx      # Todo list container
├── lib/
│   └── api.ts            # API client
└── package.json
```

## API Connection

The frontend connects to the Go API server at `http://localhost:8080` by default.

All API calls go through `/lib/api.ts` which handles:
- GET /api/todos
- POST /api/todos
- PUT /api/todos/:id
- PATCH /api/todos/:id/toggle
- DELETE /api/todos/:id

## Build

Build for production:

```bash
npm run build
```

Start production server:

```bash
npm start
```

## Troubleshooting

**Error: "Failed to fetch todos"**
- Make sure Go API server is running on port 8080
- Check API health: `curl http://localhost:8080/api/health`

**CORS errors**
- API server has CORS enabled for `http://localhost:3000`
- If using different port, update API CORS config

**API URL different**
- Set `NEXT_PUBLIC_API_URL` in `.env.local`
- Restart Next.js dev server after changing env vars
