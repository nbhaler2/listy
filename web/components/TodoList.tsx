'use client';

import { Todo } from '@/lib/api';
import TodoItem from './TodoItem';

interface TodoListProps {
  todos: Todo[];
  onToggle: (id: number) => void;
  onDelete: (id: number) => void;
  onUpdate: (id: number, item: string) => void;
}

export default function TodoList({ todos, onToggle, onDelete, onUpdate }: TodoListProps) {
  if (todos.length === 0) {
    return (
      <div className="text-center py-16">
        <div className="inline-flex items-center justify-center w-20 h-20 bg-gray-100 rounded-full mb-4">
          <svg className="w-10 h-10 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
        </div>
        <p className="text-xl font-semibold text-gray-700 mb-2">No todos found</p>
        <p className="text-gray-500">Add a todo above to get started!</p>
      </div>
    );
  }

  return (
    <div className="space-y-3 animate-in fade-in slide-in-from-bottom-4">
      {todos.map((todo, index) => (
        <div
          key={todo.id}
          style={{ animationDelay: `${index * 50}ms` }}
          className="animate-in fade-in slide-in-from-left-4"
        >
          <TodoItem
            todo={todo}
            onToggle={onToggle}
            onDelete={onDelete}
            onUpdate={onUpdate}
          />
        </div>
      ))}
    </div>
  );
}

