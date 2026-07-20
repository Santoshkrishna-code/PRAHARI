import React from 'react';
import { KPICard } from '@prahari/components/KPICard.tsx';
import { Card } from '@prahari/components/Card.tsx';

interface DashboardProps {
  onNavigate: (screen: string) => void;
}

export const Dashboard: React.FC<DashboardProps> = ({ onNavigate }) => {
  return (
    <div className="flex flex-col gap-6 w-full max-w-lg mx-auto p-4">
      <div className="flex justify-between items-center">
        <h2 className="text-lg font-bold text-text">Field Operations Console</h2>
      </div>

      <div className="grid grid-cols-2 gap-4">
        <KPICard title="My Permits" value="3 Active" change="PRM-001 approved" trend="up" />
        <KPICard title="My Inspections" value="5 Queued" change="Due today" trend="neutral" />
      </div>

      <Card title="Technician Workspace Actions">
        <div className="flex flex-col gap-3">
          <button
            onClick={() => onNavigate('incidents')}
            className="p-4 border border-border rounded-lg text-left bg-surface hover:bg-primary hover:text-background transition-colors flex flex-col gap-0.5"
          >
            <span className="text-sm font-bold">Report New Incident</span>
            <span className="text-xs text-muted">Log offline incident with automatic GPS coordinates tag.</span>
          </button>

          <button
            onClick={() => onNavigate('inspections')}
            className="p-4 border border-border rounded-lg text-left bg-surface hover:bg-primary hover:text-background transition-colors flex flex-col gap-0.5"
          >
            <span className="text-sm font-bold">Start Equipment Inspection</span>
            <span className="text-xs text-muted">Go through check lists pass/fail audit form.</span>
          </button>
        </div>
      </Card>
    </div>
  );
};
