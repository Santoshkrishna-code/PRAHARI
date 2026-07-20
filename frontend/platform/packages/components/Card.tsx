import React from 'react';

interface CardProps {
  title?: string;
  children: React.ReactNode;
  className?: string;
}

export const Card: React.FC<CardProps> = ({ title, children, className = '' }) => {
  return (
    <div className={`p-5 border border-border rounded-lg bg-surface shadow-sm ${className}`}>
      {title && (
        <div className="border-b border-border pb-3 mb-4">
          <h3 className="text-base font-semibold text-text">{title}</h3>
        </div>
      )}
      <div className="text-sm text-text">{children}</div>
    </div>
  );
};
