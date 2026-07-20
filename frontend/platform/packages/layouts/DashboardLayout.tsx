import React from 'react';

interface DashboardLayoutProps {
  sidebar: React.ReactNode;
  header: React.ReactNode;
  children: React.ReactNode;
}

export const DashboardLayout: React.FC<DashboardLayoutProps> = ({ sidebar, header, children }) => {
  return (
    <div className="flex h-screen bg-background overflow-hidden w-full">
      {/* Sidebar Panel Left */}
      <aside className="w-64 border-r border-border bg-surface hidden md:flex flex-col flex-shrink-0">
        {sidebar}
      </aside>

      {/* Main Content Area Right */}
      <div className="flex flex-col flex-1 overflow-hidden min-w-0">
        {/* Header Panel */}
        <header className="h-16 border-b border-border bg-surface flex items-center justify-between px-6 flex-shrink-0">
          {header}
        </header>

        {/* Content canvas */}
        <main className="flex-1 overflow-y-auto p-6 bg-background">
          {children}
        </main>
      </div>
    </div>
  );
};
