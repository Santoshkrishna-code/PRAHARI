import React from 'react';

interface KPICardProps {
  title: string;
  value: string | number;
  change?: string;
  trend?: 'up' | 'down' | 'neutral';
  icon?: React.ReactNode;
}

export const KPICard: React.FC<KPICardProps> = ({ title, value, change, trend, icon }) => {
  const trendColors = {
    up: 'text-success',
    down: 'text-danger',
    neutral: 'text-muted'
  };

  return (
    <div className="p-5 border border-border rounded-lg bg-surface flex items-center justify-between shadow-sm">
      <div className="flex flex-col gap-1">
        <span className="text-xs font-semibold uppercase tracking-wider text-text/60">{title}</span>
        <span className="text-2xl font-bold text-text">{value}</span>
        {change && (
          <span className={`text-xs font-medium ${trend ? trendColors[trend] : 'text-text/70'}`}>
            {change}
          </span>
        )}
      </div>
      {icon && <div className="text-text/50 p-2.5 bg-background rounded-md">{icon}</div>}
    </div>
  );
};
