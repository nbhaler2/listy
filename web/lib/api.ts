// API client for connecting to Go API Server

// Ensure API URL has protocol, remove trailing slash
const getApiUrl = () => {
  const url = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
  // Add https:// if missing protocol
  if (!url.startsWith('http://') && !url.startsWith('https://')) {
    return `https://${url}`;
  }
  return url.replace(/\/$/, ''); // Remove trailing slash
};

const API_BASE_URL = getApiUrl();

export interface Todo {
  id: number;
  item: string;
  done: boolean;
}

export interface ApiResponse<T> {
  success: boolean;
  data: T;
  error?: string;
  message?: string;
}

// Get all todos
export async function getTodos(): Promise<Todo[]> {
  const response = await fetch(`${API_BASE_URL}/api/todos`);
  if (!response.ok) {
    throw new Error('Failed to fetch todos');
  }
  const result: ApiResponse<Todo[]> = await response.json();
  if (!result.success) {
    throw new Error(result.error || 'Failed to fetch todos');
  }
  return result.data;
}

// Get pending todos
export async function getPendingTodos(): Promise<Todo[]> {
  const response = await fetch(`${API_BASE_URL}/api/todos/pending`);
  if (!response.ok) {
    throw new Error('Failed to fetch pending todos');
  }
  const result: ApiResponse<Todo[]> = await response.json();
  if (!result.success) {
    throw new Error(result.error || 'Failed to fetch pending todos');
  }
  return result.data;
}

// Get completed todos
export async function getCompletedTodos(): Promise<Todo[]> {
  const response = await fetch(`${API_BASE_URL}/api/todos/completed`);
  if (!response.ok) {
    throw new Error('Failed to fetch completed todos');
  }
  const result: ApiResponse<Todo[]> = await response.json();
  if (!result.success) {
    throw new Error(result.error || 'Failed to fetch completed todos');
  }
  return result.data;
}

// Create a new todo
export async function createTodo(item: string): Promise<Todo> {
  const response = await fetch(`${API_BASE_URL}/api/todos`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ item }),
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to create todo');
  }
  const result: ApiResponse<Todo> = await response.json();
  if (!result.success) {
    throw new Error(result.error || 'Failed to create todo');
  }
  return result.data;
}

// Update a todo
export async function updateTodo(id: number, updates: { item?: string; done?: boolean }): Promise<Todo> {
  const response = await fetch(`${API_BASE_URL}/api/todos/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(updates),
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to update todo');
  }
  const result: ApiResponse<Todo> = await response.json();
  if (!result.success) {
    throw new Error(result.error || 'Failed to update todo');
  }
  return result.data;
}

// Delete a todo
export async function deleteTodo(id: number): Promise<void> {
  const response = await fetch(`${API_BASE_URL}/api/todos/${id}`, {
    method: 'DELETE',
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to delete todo');
  }
}

// Toggle todo status
export async function toggleTodo(id: number): Promise<Todo> {
  const response = await fetch(`${API_BASE_URL}/api/todos/${id}/toggle`, {
    method: 'PATCH',
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to toggle todo');
  }
  const result: ApiResponse<Todo> = await response.json();
  if (!result.success) {
    throw new Error(result.error || 'Failed to toggle todo');
  }
  return result.data;
}

