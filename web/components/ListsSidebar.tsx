'use client';

import { useState, useEffect } from 'react';
import { getAllLists, getTodosByList } from '@/lib/api';

interface ListsSidebarProps {
  selectedListId: string | null;
  onSelectList: (listId: string | null) => void;
  onListsChanged: number; // Counter to trigger refresh
}

export default function ListsSidebar({ selectedListId, onSelectList, onListsChanged }: ListsSidebarProps) {
  const [lists, setLists] = useState<string[]>([]);
  const [loading, setLoading] = useState(true);
  const [listCounts, setListCounts] = useState<Record<string, number>>({});

  const fetchLists = async () => {
    try {
      setLoading(true);
      const fetchedLists = await getAllLists();
      setLists(fetchedLists);

      // Fetch counts for each list
      const counts: Record<string, number> = {};
      for (const listId of fetchedLists) {
        try {
          const todos = await getTodosByList(listId);
          counts[listId] = todos.length;
        } catch (err) {
          counts[listId] = 0;
        }
      }
      setListCounts(counts);
    } catch (err) {
      console.error('Error fetching lists:', err);
      setLists([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchLists();
  }, []);

  // Refresh when lists change
  useEffect(() => {
    const interval = setInterval(() => {
      fetchLists();
    }, 3000); // Refresh every 3 seconds

    return () => clearInterval(interval);
  }, []);

  // Expose refresh function to parent
  useEffect(() => {
    // This will be called when parent wants to refresh
    if (listsChanged > 0) {
      fetchLists();
    }
  }, [listsChanged]);

  const getListDisplayName = (listId: string) => {
    // Capitalize first letter and replace underscores with spaces
    return listId
      .split('_')
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');
  };

  return (
    <div className="w-64 bg-white border-r border-gray-200 h-full overflow-y-auto">
      <div className="p-4 border-b border-gray-200">
        <h2 className="text-lg font-bold text-gray-800 flex items-center gap-2">
          <svg className="w-5 h-5 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
          Lists
        </h2>
      </div>

      <div className="p-2">
        {/* Main List */}
        <button
          onClick={() => onSelectList(null)}
          className={`w-full text-left px-4 py-3 rounded-lg transition-all mb-2 ${
            selectedListId === null
              ? 'bg-gradient-to-r from-blue-500 to-indigo-600 text-white shadow-md'
              : 'hover:bg-gray-100 text-gray-700'
          }`}
        >
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
              <span className="font-semibold">Main List</span>
            </div>
          </div>
        </button>

        {/* AI Generated Lists */}
        {loading ? (
          <div className="text-center py-8">
            <div className="inline-block animate-spin rounded-full h-6 w-6 border-2 border-blue-200 border-t-blue-600"></div>
          </div>
        ) : lists.length === 0 ? (
          <div className="text-center py-8 text-gray-500 text-sm">
            <p>No separate lists yet</p>
            <p className="text-xs mt-1">Create one using AI Task Generator</p>
          </div>
        ) : (
          <div className="space-y-1">
            {lists.map((listId) => (
              <button
                key={listId}
                onClick={() => onSelectList(listId)}
                className={`w-full text-left px-4 py-3 rounded-lg transition-all ${
                  selectedListId === listId
                    ? 'bg-gradient-to-r from-purple-500 to-indigo-600 text-white shadow-md'
                    : 'hover:bg-gray-100 text-gray-700'
                }`}
              >
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-2 flex-1 min-w-0">
                    <svg className="w-4 h-4 flex-shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M7 7h.01M7 3h5c.512 0 1.024.195 1.414.586l7 7a2 2 0 010 2.828l-7 7a2 2 0 01-2.828 0l-7-7A1.994 1.994 0 013 12V7a4 4 0 014-4z" />
                    </svg>
                    <span className="font-medium truncate">{getListDisplayName(listId)}</span>
                  </div>
                  {listCounts[listId] !== undefined && (
                    <span
                      className={`ml-2 px-2 py-0.5 rounded-full text-xs font-semibold flex-shrink-0 ${
                        selectedListId === listId
                          ? 'bg-white/20 text-white'
                          : 'bg-gray-200 text-gray-600'
                      }`}
                    >
                      {listCounts[listId]}
                    </span>
                  )}
                </div>
              </button>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

