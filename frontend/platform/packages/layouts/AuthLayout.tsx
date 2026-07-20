import React from 'react';

interface AuthLayoutProps {
  children: React.ReactNode;
  title?: string;
}

export const AuthLayout: React.FC<AuthLayoutProps> = ({ children, title = 'PRAHARI Platform Log In' }) => {
  return (
    <div className="flex min-h-screen bg-background items-center justify-center p-4 w-full">
      <div className="w-full max-w-md p-8 border border-border rounded-lg bg-surface shadow-lg animate-in fade-in duration-300">
        <div className="flex flex-col gap-2 text-center mb-6">
          <h2 className="text-xl font-bold text-text">{title}</h2>
          <span className="text-xs text-muted">Enter credentials to authenticate session</span>
        </div>
        {children}
      </div>
    </div>
  );
};
