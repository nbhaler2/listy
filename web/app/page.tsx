'use client';

import { useState, useEffect } from 'react';
import { Todo, getTodos, getPendingTodos, getCompletedTodos, createTodo, updateTodo, deleteTodo, toggleTodo } from '@/lib/api';
import AddTodoForm from '@/components/AddTodoForm';
import TodoList from '@/components/TodoList';
import AITaskGenerator from '@/components/AITaskGenerator';

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
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50 py-8 px-4 sm:py-12">
      <div className="max-w-3xl mx-auto">
        {/* Header */}
        <div className="text-center mb-10">
          <div className="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-2xl shadow-lg mb-4 transform hover:scale-105 transition-transform">
            <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
            </svg>
          </div>
          <h1 className="text-5xl font-extrabold bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent mb-2">
            Listy
          </h1>
          <p className="text-gray-600 text-lg">Your smart todo list manager</p>
        </div>

        {/* AI Task Generator Section */}
        <AITaskGenerator onTasksCreated={fetchTodos} />

        {/* Main Card */}
        <div className="bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden">
          <div className="p-6 sm:p-8">
            {/* Filter Tabs */}
            <div className="flex gap-2 mb-8 bg-gray-50 p-1 rounded-xl">
              <button
                onClick={() => setFilter('all')}
                className={`flex-1 px-4 py-3 font-semibold rounded-lg transition-all duration-200 ${
                  filter === 'all'
                    ? 'bg-white text-blue-600 shadow-md transform scale-105'
                    : 'text-gray-600 hover:text-gray-800 hover:bg-gray-100'
                }`}
              >
                All
              </button>
              <button
                onClick={() => setFilter('pending')}
                className={`flex-1 px-4 py-3 font-semibold rounded-lg transition-all duration-200 ${
                  filter === 'pending'
                    ? 'bg-white text-blue-600 shadow-md transform scale-105'
                    : 'text-gray-600 hover:text-gray-800 hover:bg-gray-100'
                }`}
              >
                Pending
              </button>
              <button
                onClick={() => setFilter('completed')}
                className={`flex-1 px-4 py-3 font-semibold rounded-lg transition-all duration-200 ${
                  filter === 'completed'
                    ? 'bg-white text-blue-600 shadow-md transform scale-105'
                    : 'text-gray-600 hover:text-gray-800 hover:bg-gray-100'
                }`}
              >
                Completed
              </button>
            </div>

          {/* Add Todo Form */}
          <AddTodoForm onAdd={handleAdd} />

            {/* Error Message */}
            {error && (
              <div className="mb-6 p-4 bg-red-50 border-l-4 border-red-500 rounded-lg flex items-center justify-between animate-in slide-in-from-top">
                <div className="flex items-center gap-3">
                  <svg className="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <span className="text-red-700 font-medium">{error}</span>
                </div>
                <button
                  onClick={() => setError(null)}
                  className="text-red-500 hover:text-red-700 hover:bg-red-100 rounded-full p-1 transition-colors"
                >
                  <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
            )}

            {/* Loading State */}
            {loading ? (
              <div className="text-center py-16">
                <div className="inline-block animate-spin rounded-full h-12 w-12 border-4 border-blue-200 border-t-blue-600"></div>
                <p className="mt-6 text-gray-600 font-medium">Loading your todos...</p>
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
              <div className="mt-8 pt-6 border-t border-gray-200">
                <div className="flex items-center justify-center gap-6 text-sm">
                  {filter === 'all' && (
                    <>
                      <div className="flex items-center gap-2">
                        <div className="w-3 h-3 rounded-full bg-amber-400"></div>
                        <span className="text-gray-600 font-medium">
                          <span className="font-bold text-gray-800">{todos.filter((t) => !t.done).length}</span> pending
                        </span>
                      </div>
                      <div className="flex items-center gap-2">
                        <div className="w-3 h-3 rounded-full bg-green-500"></div>
                        <span className="text-gray-600 font-medium">
                          <span className="font-bold text-gray-800">{todos.filter((t) => t.done).length}</span> completed
                        </span>
                      </div>
                    </>
                  )}
                  {filter === 'pending' && (
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded-full bg-amber-400"></div>
                      <span className="text-gray-600 font-medium">
                        <span className="font-bold text-gray-800">{todos.length}</span> pending {todos.length === 1 ? 'todo' : 'todos'}
                      </span>
                    </div>
                  )}
                  {filter === 'completed' && (
                    <div className="flex items-center gap-2">
                      <div className="w-3 h-3 rounded-full bg-green-500"></div>
                      <span className="text-gray-600 font-medium">
                        <span className="font-bold text-gray-800">{todos.length}</span> completed {todos.length === 1 ? 'todo' : 'todos'}
                      </span>
                    </div>
                  )}
                </div>
              </div>
            )}
          </div>
        </div>

        {/* Footer */}
        <div className="text-center mt-8 text-sm text-gray-500">
          <p className="flex items-center justify-center gap-2">
            <span>Built with</span>
            <span className="font-semibold text-blue-600">Go</span>
            <span>+</span>
            <span className="font-semibold text-indigo-600">Next.js</span>
          </p>
        </div>
        </div>
    </div>
  );
}
