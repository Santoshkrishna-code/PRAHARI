import React from 'react';

interface NavItem {
  label: string;
  href: string;
  icon?: React.ReactNode;
}

interface SidebarProps {
  items: NavItem[];
  currentPath?: string;
  onNavigate?: (href: string) => void;
}

export const Sidebar: React.FC<SidebarProps> = ({ items, currentPath = '/', onNavigate }) => {
  return (
    <div className="flex flex-col h-full py-4 px-3 gap-6">
      {/* Title */}
      <div className="flex items-center gap-2 px-3">
        <span className="text-sm font-black tracking-wider text-text uppercase">PRAHARI Platform</span>
      </div>

      {/* Nav List */}
      <nav className="flex flex-col gap-1.5 flex-1 overflow-y-auto">
        {items.map((item, index) => {
          const isActive = currentPath === item.href;
          return (
            <button
              key={index}
              onClick={() => onNavigate?.(item.href)}
              className={`flex items-center gap-3 px-3 py-2 text-sm font-semibold rounded-md transition-colors w-full text-left ${
                isActive
                  ? 'bg-primary text-background'
                  : 'text-text/75 hover:bg-border/20 hover:text-text'
              }`}
            >
              {item.icon && <span className="h-4 w-4 flex-shrink-0">{item.icon}</span>}
              <span>{item.label}</span>
            </button>
          );
        })}
      </nav>
    </div>
  );
};
