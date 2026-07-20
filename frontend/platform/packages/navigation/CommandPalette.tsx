import React, { useState, useEffect } from 'react';

interface Command {
  name: string;
  action: () => void;
}

interface CommandPaletteProps {
  isOpen: boolean;
  onClose: () => void;
  commands: Command[];
}

export const CommandPalette: React.FC<CommandPaletteProps> = ({ isOpen, onClose, commands }) => {
  const [query, setQuery] = useState('');

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'k' && (e.metaKey || e.ctrlKey)) {
        e.preventDefault();
        if (isOpen) onClose();
        // Trigger external toggle logic if needed. Let's make it self-contained in playground
      }
    };
    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, [isOpen, onClose]);

  if (!isOpen) return null;

  const filtered = commands.filter((cmd) =>
    cmd.name.toLowerCase().includes(query.toLowerCase())
  );

  return (
    <div className="fixed inset-0 z-50 flex items-start justify-center pt-[15vh] p-4">
      {/* Backdrop */}
      <div className="fixed inset-0 bg-black/50" onClick={onClose} />

      {/* Palette Box */}
      <div className="relative bg-surface border border-border rounded-lg shadow-2xl w-full max-w-lg overflow-hidden animate-in fade-in zoom-in-95 duration-200">
        <div className="flex items-center border-b border-border px-3 py-2.5">
          <svg className="h-5 w-5 text-muted mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
          <input
            className="w-full text-sm bg-transparent border-none focus:outline-none text-text placeholder-muted"
            placeholder="Type a command or query..."
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            autoFocus
          />
        </div>

        <div className="max-h-[300px] overflow-y-auto p-2">
          {filtered.length === 0 ? (
            <div className="py-6 text-center text-xs text-muted">No commands found.</div>
          ) : (
            filtered.map((cmd, index) => (
              <button
                key={index}
                onClick={() => {
                  cmd.action();
                  onClose();
                }}
                className="w-full text-left px-3 py-2 text-sm text-text rounded-md hover:bg-primary hover:text-background focus:outline-none transition-colors"
              >
                {cmd.name}
              </button>
            ))
          )}
        </div>
      </div>
    </div>
  );
};
