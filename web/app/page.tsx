'use client';

import { useState, useEffect } from 'react';
import { Todo, getTodos, getPendingTodos, getCompletedTodos, createTodo, updateTodo, deleteTodo, toggleTodo } from '@/lib/api';
import AddTodoForm from '@/components/AddTodoForm';
import TodoList from '@/components/TodoList';

type Filter = 'all' | 'pending' | 'completed';

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [filter, setFilter] = useState<Filter>('all');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Fetch todos
  const fetchTodos = async () => {
    try {
      setLoading(true);
      setError(null);
      let fetchedTodos: Todo[];
      
      switch (filter) {
        case 'pending':
          fetchedTodos = await getPendingTodos();
          break;
        case 'completed':
          fetchedTodos = await getCompletedTodos();
          break;
        default:
          fetchedTodos = await getTodos();
      }
      
      setTodos(fetchedTodos);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to load todos');
      console.error('Error fetching todos:', err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchTodos();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [filter]);

  const handleAdd = async (item: string) => {
    try {
      const newTodo = await createTodo(item);
      await fetchTodos(); // Refresh list
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to add todo');
      throw err;
    }
  };

  const handleToggle = async (id: number) => {
    try {
      await toggleTodo(id);
      await fetchTodos(); // Refresh list
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to toggle todo');
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await deleteTodo(id);
      await fetchTodos(); // Refresh list
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to delete todo');
    }
  };

  const handleUpdate = async (id: number, item: string) => {
    try {
      await updateTodo(id, { item });
      await fetchTodos(); // Refresh list
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to update todo');
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 py-12 px-4">
      <div className="max-w-2xl mx-auto">
        {/* Header */}
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold text-gray-800 mb-2">Listy</h1>
          <p className="text-gray-600">Your Todo List Manager</p>
        </div>

        {/* Main Card */}
        <div className="bg-white rounded-xl shadow-lg p-6">
          {/* Filter Tabs */}
          <div className="flex gap-2 mb-6 border-b border-gray-200">
            <button
              onClick={() => setFilter('all')}
              className={`px-4 py-2 font-medium transition-colors ${
                filter === 'all'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800'
              }`}
            >
              All
            </button>
            <button
              onClick={() => setFilter('pending')}
              className={`px-4 py-2 font-medium transition-colors ${
                filter === 'pending'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800'
              }`}
            >
              Pending
            </button>
            <button
              onClick={() => setFilter('completed')}
              className={`px-4 py-2 font-medium transition-colors ${
                filter === 'completed'
                  ? 'text-blue-600 border-b-2 border-blue-600'
                  : 'text-gray-600 hover:text-gray-800'
              }`}
            >
              Completed
            </button>
          </div>

          {/* Add Todo Form */}
          <AddTodoForm onAdd={handleAdd} />

          {/* Error Message */}
          {error && (
            <div className="mb-4 p-3 bg-red-50 border border-red-200 text-red-700 rounded-lg">
              {error}
              <button
                onClick={() => setError(null)}
                className="ml-2 text-red-500 hover:text-red-700"
              >
                ×
              </button>
            </div>
          )}

          {/* Loading State */}
          {loading ? (
            <div className="text-center py-12">
              <div className="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-blue-600"></div>
              <p className="mt-4 text-gray-600">Loading todos...</p>
            </div>
          ) : (
            <TodoList
              todos={todos}
              onToggle={handleToggle}
              onDelete={handleDelete}
              onUpdate={handleUpdate}
            />
          )}

          {/* Stats */}
          {!loading && todos.length > 0 && (
            <div className="mt-6 pt-4 border-t border-gray-200 text-sm text-gray-600">
              {filter === 'all' && (
                <p>
                  {todos.filter((t) => !t.done).length} pending, {todos.filter((t) => t.done).length} completed
                </p>
              )}
              {filter === 'pending' && <p>{todos.length} pending todos</p>}
              {filter === 'completed' && <p>{todos.length} completed todos</p>}
            </div>
          )}
        </div>

        {/* Footer */}
        <div className="text-center mt-6 text-sm text-gray-500">
          <p>Powered by Go API + Next.js</p>
        </div>
      </div>
    </div>
  );
}
