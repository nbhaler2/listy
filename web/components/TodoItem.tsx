'use client';

import { Todo } from '@/lib/api';
import { useState } from 'react';

interface TodoItemProps {
  todo: Todo;
  onToggle: (id: number) => void;
  onDelete: (id: number) => void;
  onUpdate: (id: number, item: string) => void;
}

export default function TodoItem({ todo, onToggle, onDelete, onUpdate }: TodoItemProps) {
  const [isEditing, setIsEditing] = useState(false);
  const [editText, setEditText] = useState(todo.item);

  const handleUpdate = () => {
    if (editText.trim() && editText !== todo.item) {
      onUpdate(todo.id, editText.trim());
    }
    setIsEditing(false);
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      handleUpdate();
    } else if (e.key === 'Escape') {
      setEditText(todo.item);
      setIsEditing(false);
    }
  };

  return (
    <div className={`group flex items-center gap-4 p-5 bg-white rounded-xl shadow-sm border-2 transition-all duration-200 hover:shadow-lg hover:border-blue-200 ${
      todo.done ? 'border-gray-100 bg-gray-50' : 'border-gray-200'
    }`}>
      <button
        onClick={() => onToggle(todo.id)}
        className={`flex-shrink-0 w-7 h-7 rounded-full border-2 flex items-center justify-center transition-all duration-200 transform hover:scale-110 ${
          todo.done
            ? 'bg-gradient-to-br from-green-400 to-green-600 border-green-500 text-white shadow-md'
            : 'border-gray-300 hover:border-green-400 hover:bg-green-50'
        }`}
      >
        {todo.done && (
          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={3} d="M5 13l4 4L19 7" />
          </svg>
        )}
      </button>

      <div className="flex-1 min-w-0">
        {isEditing ? (
          <input
            type="text"
            value={editText}
            onChange={(e) => setEditText(e.target.value)}
            onBlur={handleUpdate}
            onKeyDown={handleKeyPress}
            className="w-full px-4 py-2 border-2 border-blue-400 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 bg-white text-gray-800 font-medium"
            autoFocus
          />
        ) : (
          <div>
            <span
              onClick={() => setIsEditing(true)}
              className={`cursor-text font-medium transition-colors block ${
                todo.done 
                  ? 'line-through text-gray-400' 
                  : 'text-gray-800 hover:text-blue-600'
              }`}
            >
              {todo.item}
            </span>
            {/* Metadata badges */}
            {(todo.priority || todo.estimated_time || todo.category) && (
              <div className="flex flex-wrap gap-2 mt-2">
                {todo.priority && (
                  <span
                    className={`px-2 py-1 text-xs font-semibold rounded-full ${
                      todo.priority === 'high'
                        ? 'bg-red-100 text-red-700 border border-red-300'
                        : todo.priority === 'medium'
                        ? 'bg-yellow-100 text-yellow-700 border border-yellow-300'
                        : 'bg-green-100 text-green-700 border border-green-300'
                    }`}
                  >
                    {todo.priority}
                  </span>
                )}
                {todo.estimated_time && (
                  <span className="px-2 py-1 text-xs font-medium rounded-full bg-blue-100 text-blue-700 border border-blue-300 flex items-center gap-1">
                    <svg className="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    {todo.estimated_time}
                  </span>
                )}
                {todo.category && (
                  <span className="px-2 py-1 text-xs font-medium rounded-full bg-purple-100 text-purple-700 border border-purple-300 flex items-center gap-1">
                    <svg className="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                    </svg>
                    {todo.category}
                  </span>
                )}
              </div>
            )}
          </div>
        )}
      </div>

      <div className="flex gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
        <button
          onClick={() => setIsEditing(true)}
          className="p-2 text-blue-600 hover:bg-blue-50 rounded-lg transition-colors"
          title="Edit"
        >
          <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
          </svg>
        </button>
        <button
          onClick={() => onDelete(todo.id)}
          className="p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors"
          title="Delete"
        >
          <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
        </button>
      </div>
    </div>
  );
}

