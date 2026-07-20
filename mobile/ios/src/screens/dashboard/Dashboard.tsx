import React from 'react';
import { KPICard } from '@prahari/components/KPICard.tsx';
import { Card } from '@prahari/components/Card.tsx';

interface DashboardProps {
  onNavigate: (screen: string) => void;
}

export const Dashboard: React.FC<DashboardProps> = ({ onNavigate }) => {
  return (
    <div className="flex flex-col gap-6 w-full max-w-lg mx-auto p-4 bg-background">
      <div className="flex justify-between items-center">
        <h2 className="text-xl font-bold tracking-tight text-text">Operations Panel</h2>
      </div>

      <div className="grid grid-cols-2 gap-4">
        <KPICard title="My Permits" value="2 Active" change="PRM-002 approved" trend="up" />
        <KPICard title="My Inspections" value="4 Queued" change="Due tomorrow" trend="neutral" />
      </div>

      <Card title="Technician Actions Workspace">
        <div className="flex flex-col gap-3">
          <button
            onClick={() => onNavigate('incidents')}
            className="p-4 border border-border rounded-xl text-left bg-surface hover:bg-primary hover:text-background transition-all flex flex-col gap-0.5"
          >
            <span className="text-sm font-semibold">Report Incident Log</span>
            <span className="text-xs text-muted">Log details with Apple Core Location GPS coords.</span>
          </button>

          <button
            onClick={() => onNavigate('inspections')}
            className="p-4 border border-border rounded-xl text-left bg-surface hover:bg-primary hover:text-background transition-all flex flex-col gap-0.5"
          >
            <span className="text-sm font-semibold">Equipment Inspection</span>
            <span className="text-xs text-muted">Run through asset check list audit form.</span>
          </button>
        </div>
      </Card>
    </div>
  );
};
