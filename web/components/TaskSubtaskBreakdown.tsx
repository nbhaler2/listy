'use client';

import { useState } from 'react';
import { Todo, generateSubtaskBreakdown, createAITasks } from '@/lib/api';

interface TaskSubtaskBreakdownProps {
  task: Todo;
  onClose: () => void;
  onTasksCreated: () => void;
}

export default function TaskSubtaskBreakdown({ task, onClose, onTasksCreated }: TaskSubtaskBreakdownProps) {
  const [isGenerating, setIsGenerating] = useState(false);
  const [subtasks, setSubtasks] = useState<string[]>([]);
  const [isCreating, setIsCreating] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [expanded, setExpanded] = useState(true);

  const handleGenerateSubtasks = async () => {
    setIsGenerating(true);
    setError(null);
    setSubtasks([]);

    try {
      const response = await generateSubtaskBreakdown(task.item);
      // Extract just the text from AI tasks
      const taskTexts = response.suggested_tasks.map(t => t.text);
      
      if (taskTexts.length === 0) {
        setError('This task cannot be meaningfully broken down into subtasks. It appears to be simple enough to complete as-is.');
      } else {
        setSubtasks(taskTexts);
        setExpanded(true);
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to generate subtasks');
      console.error('Error generating subtasks:', err);
    } finally {
      setIsGenerating(false);
    }
  };

  const handleCreateSubtasks = async () => {
    if (subtasks.length === 0) {
      setError('No subtasks to create');
      return;
    }

    setIsCreating(true);
    setError(null);

    try {
      // Create subtasks as AI tasks (with empty metadata)
      const aiTasks = subtasks.map(text => ({
        text,
        priority: '',
        estimated_time: '',
        category: '',
      }));

      // Create subtasks in the same list as the parent task
      await createAITasks(aiTasks, task.list_id || null);
      onTasksCreated(); // This will trigger a refresh in the parent
      onClose();
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create subtasks');
      console.error('Error creating subtasks:', err);
    } finally {
      setIsCreating(false);
    }
  };

  const handleEditSubtask = (index: number, newText: string) => {
    const updated = [...subtasks];
    updated[index] = newText;
    setSubtasks(updated);
  };

  const handleDeleteSubtask = (index: number) => {
    setSubtasks(subtasks.filter((_, i) => i !== index));
  };

  const handleAddCustomSubtask = () => {
    setSubtasks([...subtasks, '']);
  };

  return (
    <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50 p-4" onClick={onClose}>
      <div
        className="bg-white rounded-2xl shadow-2xl max-w-2xl w-full max-h-[90vh] overflow-hidden flex flex-col"
        onClick={(e) => e.stopPropagation()}
      >
        {/* Header */}
        <div className="p-6 border-b border-gray-200 bg-gradient-to-r from-purple-50 to-indigo-50">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-3">
              <div className="w-10 h-10 bg-gradient-to-br from-purple-500 to-indigo-600 rounded-lg flex items-center justify-center shadow-md">
                <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                </svg>
              </div>
              <div>
                <h2 className="text-xl font-bold text-gray-800">Generate Subtasks</h2>
                <p className="text-sm text-gray-600">Break down: "{task.item}"</p>
              </div>
            </div>
            <button
              onClick={onClose}
              className="p-2 text-gray-400 hover:text-gray-600 hover:bg-gray-100 rounded-lg transition-colors"
            >
              <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>

        {/* Content */}
        <div className="flex-1 overflow-y-auto p-6">
          {!subtasks.length && !isGenerating && (
            <div className="text-center py-12">
              <div className="inline-flex items-center justify-center w-16 h-16 bg-purple-100 rounded-full mb-4">
                <svg className="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                </svg>
              </div>
              <p className="text-gray-700 font-medium mb-2">Generate subtasks for this task</p>
              <p className="text-gray-500 text-sm mb-6">AI will break down "{task.item}" into actionable subtasks</p>
              <button
                onClick={handleGenerateSubtasks}
                className="px-6 py-3 bg-gradient-to-r from-purple-600 to-indigo-600 text-white rounded-xl hover:from-purple-700 hover:to-indigo-700 transition-all font-semibold shadow-lg hover:shadow-xl transform hover:scale-105"
              >
                Generate Subtasks
              </button>
            </div>
          )}

          {isGenerating && (
            <div className="text-center py-12">
              <div className="inline-block animate-spin rounded-full h-12 w-12 border-4 border-purple-200 border-t-purple-600 mb-4"></div>
              <p className="text-gray-600 font-medium">Generating subtasks...</p>
            </div>
          )}

          {error && (
            <div className={`mb-4 p-4 rounded-lg border-l-4 ${
              error.includes('cannot be meaningfully broken down')
                ? 'bg-blue-50 border-blue-500'
                : 'bg-red-50 border-red-500'
            }`}>
              <div className="flex items-start gap-3">
                {error.includes('cannot be meaningfully broken down') ? (
                  <svg className="w-5 h-5 text-blue-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                ) : (
                  <svg className="w-5 h-5 text-red-500 mt-0.5 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                )}
                <div className="flex-1">
                  <span className={`font-medium ${
                    error.includes('cannot be meaningfully broken down')
                      ? 'text-blue-700'
                      : 'text-red-700'
                  }`}>
                    {error}
                  </span>
                  {error.includes('cannot be meaningfully broken down') && (
                    <p className="text-sm text-blue-600 mt-1">
                      This task is simple enough to complete as-is. You can still add custom subtasks manually if needed.
                    </p>
                  )}
                </div>
              </div>
            </div>
          )}

          {(subtasks.length > 0 || (error && error.includes('cannot be meaningfully broken down'))) && (
            <div className="space-y-3">
              <div className="flex items-center justify-between mb-4">
                <h3 className="text-lg font-semibold text-gray-800">
                  {subtasks.length > 0 ? `Subtasks (${subtasks.length})` : 'Add Custom Subtasks'}
                </h3>
                <div className="flex gap-2">
                  <button
                    onClick={handleAddCustomSubtask}
                    className="px-3 py-1.5 text-sm bg-white border-2 border-purple-300 text-purple-600 rounded-lg hover:bg-purple-50 transition-colors font-medium"
                  >
                    + Add Custom
                  </button>
                  <button
                    onClick={() => setExpanded(!expanded)}
                    className="px-3 py-1.5 text-sm bg-gray-100 text-gray-600 rounded-lg hover:bg-gray-200 transition-colors"
                  >
                    {expanded ? 'Collapse' : 'Expand'}
                  </button>
                </div>
              </div>

              {expanded && (
                <div className="space-y-2 max-h-96 overflow-y-auto pr-2">
                  {subtasks.map((subtask, index) => (
                    <div
                      key={index}
                      className="flex items-start gap-3 p-4 bg-gray-50 rounded-xl border-2 border-gray-200 hover:border-purple-300 transition-colors"
                    >
                      <div className="flex-shrink-0 mt-1.5">
                        <div className="w-2 h-2 rounded-full bg-purple-500"></div>
                      </div>
                      <div className="flex-1">
                        <input
                          type="text"
                          value={subtask}
                          onChange={(e) => handleEditSubtask(index, e.target.value)}
                          className="w-full px-3 py-2 border-2 border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent bg-white text-gray-800"
                          placeholder="Subtask description..."
                        />
                      </div>
                      <button
                        onClick={() => handleDeleteSubtask(index)}
                        className="flex-shrink-0 p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                        title="Delete subtask"
                      >
                        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                      </button>
                    </div>
                  ))}
                </div>
              )}
            </div>
          )}
        </div>

        {/* Footer Actions */}
        {subtasks.length > 0 && (
          <div className="p-6 border-t border-gray-200 bg-gray-50">
            <div className="flex gap-3">
              <button
                onClick={onClose}
                className="flex-1 px-4 py-3 bg-white border-2 border-gray-300 text-gray-700 rounded-xl hover:bg-gray-50 transition-colors font-semibold"
              >
                Cancel
              </button>
              <button
                onClick={handleCreateSubtasks}
                disabled={isCreating || subtasks.filter((t) => t.trim()).length === 0}
                className="flex-1 px-4 py-3 bg-gradient-to-r from-green-600 to-emerald-600 text-white rounded-xl hover:from-green-700 hover:to-emerald-700 disabled:from-gray-300 disabled:to-gray-400 disabled:cursor-not-allowed transition-all font-semibold shadow-lg hover:shadow-xl"
              >
                {isCreating ? (
                  <span className="flex items-center justify-center gap-2">
                    <svg className="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
                      <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                      <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    Creating...
                  </span>
                ) : (
                  `Create ${subtasks.filter((t) => t.trim()).length} Subtask(s)`
                )}
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

