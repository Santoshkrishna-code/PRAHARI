import React, { useEffect } from 'react';

interface ToastProps {
  message: string;
  type?: 'success' | 'warning' | 'danger' | 'info';
  onClose: () => void;
}

export const Toast: React.FC<ToastProps> = ({ message, type = 'info', onClose }) => {
  useEffect(() => {
    const timer = setTimeout(onClose, 4000);
    return () => clearTimeout(timer);
  }, [onClose]);

  const borderColors = {
    success: 'border-l-success',
    warning: 'border-l-warning',
    danger: 'border-l-danger',
    info: 'border-l-primary'
  };

  return (
    <div className={`fixed bottom-5 right-5 z-50 flex items-center gap-3 p-4 bg-surface text-text border border-border border-l-4 ${borderColors[type]} rounded-md shadow-lg min-w-[280px] max-w-sm animate-in slide-in-from-bottom-5 duration-300`}>
      <span className="text-sm flex-1 font-medium">{message}</span>
      <button onClick={onClose} className="text-text/60 hover:text-text focus:outline-none">
        <svg className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  );
};
