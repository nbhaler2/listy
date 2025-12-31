'use client';

import { useState } from 'react';
import { AITask, generateTaskBreakdown, createAITasks } from '@/lib/api';

interface AITaskGeneratorProps {
  onTasksCreated: () => void;
}

export default function AITaskGenerator({ onTasksCreated }: AITaskGeneratorProps) {
  const [goal, setGoal] = useState('');
  const [isGenerating, setIsGenerating] = useState(false);
  const [generatedTasks, setGeneratedTasks] = useState<AITask[]>([]);
  const [isCreating, setIsCreating] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleGenerate = async () => {
    if (!goal.trim()) {
      setError('Please enter a goal or task');
      return;
    }

    setIsGenerating(true);
    setError(null);
    setGeneratedTasks([]);

    try {
      const response = await generateTaskBreakdown(goal.trim());
      setGeneratedTasks(response.suggested_tasks);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to generate tasks');
      console.error('Error generating tasks:', err);
    } finally {
      setIsGenerating(false);
    }
  };

  const handleEditTask = (index: number, newText: string) => {
    const updated = [...generatedTasks];
    updated[index] = { ...updated[index], text: newText };
    setGeneratedTasks(updated);
  };

  const handleDeleteTask = (index: number) => {
    setGeneratedTasks(generatedTasks.filter((_, i) => i !== index));
  };

  const handleAddCustomTask = () => {
    setGeneratedTasks([
      ...generatedTasks,
      {
        text: '',
        priority: 'medium',
        estimated_time: '',
        category: '',
      },
    ]);
  };

  const handleCreateAll = async () => {
    if (generatedTasks.length === 0) {
      setError('No tasks to create');
      return;
    }

    // Filter out empty tasks
    const validTasks = generatedTasks.filter((task) => task.text.trim() !== '');
    if (validTasks.length === 0) {
      setError('Please add at least one valid task');
      return;
    }

    setIsCreating(true);
    setError(null);

    try {
      await createAITasks(validTasks);
      // Reset form
      setGoal('');
      setGeneratedTasks([]);
      // Notify parent to refresh todos
      onTasksCreated();
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to create tasks');
      console.error('Error creating tasks:', err);
    } finally {
      setIsCreating(false);
    }
  };

  const getPriorityColor = (priority: string) => {
    switch (priority) {
      case 'high':
        return 'bg-red-100 text-red-700 border-red-300';
      case 'medium':
        return 'bg-yellow-100 text-yellow-700 border-yellow-300';
      case 'low':
        return 'bg-green-100 text-green-700 border-green-300';
      default:
        return 'bg-gray-100 text-gray-700 border-gray-300';
    }
  };

  return (
    <div className="mb-8 p-6 bg-gradient-to-br from-purple-50 to-indigo-50 rounded-2xl border-2 border-purple-200 shadow-lg">
      {/* Header */}
      <div className="flex items-center gap-3 mb-4">
        <div className="w-10 h-10 bg-gradient-to-br from-purple-500 to-indigo-600 rounded-lg flex items-center justify-center shadow-md">
          <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
          </svg>
        </div>
        <div>
          <h2 className="text-xl font-bold text-gray-800">AI Task Generator</h2>
          <p className="text-sm text-gray-600">Describe your goal and we'll break it down into actionable tasks</p>
        </div>
      </div>

      {/* Input Section */}
      <div className="mb-4">
        <div className="flex gap-3">
          <div className="flex-1 relative">
            <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
              <svg className="w-5 h-5 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
            </div>
            <input
              type="text"
              value={goal}
              onChange={(e) => setGoal(e.target.value)}
              onKeyDown={(e) => {
                if (e.key === 'Enter' && !isGenerating) {
                  handleGenerate();
                }
              }}
              placeholder="e.g., learning Go, planning a trip, building a website..."
              className="w-full pl-12 pr-4 py-3 border-2 border-purple-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent bg-white text-gray-800 placeholder-gray-400"
              disabled={isGenerating || isCreating}
            />
          </div>
          <button
            onClick={handleGenerate}
            disabled={!goal.trim() || isGenerating || isCreating}
            className="px-6 py-3 bg-gradient-to-r from-purple-600 to-indigo-600 text-white rounded-xl hover:from-purple-700 hover:to-indigo-700 disabled:from-gray-300 disabled:to-gray-400 disabled:cursor-not-allowed transition-all font-semibold shadow-lg hover:shadow-xl transform hover:scale-105 active:scale-95 disabled:transform-none"
          >
            {isGenerating ? (
              <span className="flex items-center gap-2">
                <svg className="animate-spin h-5 w-5" fill="none" viewBox="0 0 24 24">
                  <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                  <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Generating...
              </span>
            ) : (
              'Generate'
            )}
          </button>
        </div>
      </div>

      {/* Error Message */}
      {error && (
        <div className="mb-4 p-4 bg-red-50 border-l-4 border-red-500 rounded-lg flex items-center justify-between">
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

      {/* Generated Tasks Preview */}
      {generatedTasks.length > 0 && (
        <div className="mt-6 space-y-3">
          <div className="flex items-center justify-between mb-4">
            <h3 className="text-lg font-semibold text-gray-800">
              Generated Tasks ({generatedTasks.length})
            </h3>
            <button
              onClick={handleAddCustomTask}
              className="px-3 py-1.5 text-sm bg-white border-2 border-purple-300 text-purple-600 rounded-lg hover:bg-purple-50 transition-colors font-medium"
            >
              + Add Custom
            </button>
          </div>

          <div className="space-y-2 max-h-96 overflow-y-auto pr-2">
            {generatedTasks.map((task, index) => (
              <div
                key={index}
                className="flex items-start gap-3 p-4 bg-white rounded-xl border-2 border-purple-100 hover:border-purple-300 transition-colors"
              >
                <div className="flex-shrink-0 mt-1">
                  <div className={`px-2 py-1 text-xs font-semibold rounded border ${getPriorityColor(task.priority)}`}>
                    {task.priority}
                  </div>
                </div>
                <div className="flex-1">
                  <input
                    type="text"
                    value={task.text}
                    onChange={(e) => handleEditTask(index, e.target.value)}
                    className="w-full px-3 py-2 border-2 border-gray-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent bg-white text-gray-800"
                    placeholder="Task description..."
                  />
                  {(task.estimated_time || task.category) && (
                    <div className="mt-2 flex gap-3 text-xs text-gray-500">
                      {task.estimated_time && (
                        <span className="flex items-center gap-1">
                          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                          </svg>
                          {task.estimated_time}
                        </span>
                      )}
                      {task.category && (
                        <span className="flex items-center gap-1">
                          <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                          </svg>
                          {task.category}
                        </span>
                      )}
                    </div>
                  )}
                </div>
                <button
                  onClick={() => handleDeleteTask(index)}
                  className="flex-shrink-0 p-2 text-red-600 hover:bg-red-50 rounded-lg transition-colors"
                  title="Delete task"
                >
                  <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                </button>
              </div>
            ))}
          </div>

          {/* Create All Button */}
          <div className="flex gap-3 pt-4 border-t border-purple-200">
            <button
              onClick={() => {
                setGeneratedTasks([]);
                setGoal('');
              }}
              className="flex-1 px-4 py-3 bg-white border-2 border-gray-300 text-gray-700 rounded-xl hover:bg-gray-50 transition-colors font-semibold"
            >
              Cancel
            </button>
            <button
              onClick={handleCreateAll}
              disabled={isCreating || generatedTasks.filter((t) => t.text.trim()).length === 0}
              className="flex-1 px-4 py-3 bg-gradient-to-r from-green-600 to-emerald-600 text-white rounded-xl hover:from-green-700 hover:to-emerald-700 disabled:from-gray-300 disabled:to-gray-400 disabled:cursor-not-allowed transition-all font-semibold shadow-lg hover:shadow-xl transform hover:scale-105 active:scale-95 disabled:transform-none"
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
                `Create All Tasks (${generatedTasks.filter((t) => t.text.trim()).length})`
              )}
            </button>
          </div>
        </div>
      )}
    </div>
  );
}


