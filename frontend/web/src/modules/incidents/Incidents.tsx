import React from 'react';
import { DataTable } from '@prahari/components/DataTable.tsx';
import { Badge } from '@prahari/components/Badge.tsx';

interface IncidentLog {
  id: string;
  title: string;
  severity: 'CRITICAL' | 'HIGH' | 'MEDIUM';
  status: string;
  timestamp: string;
}

const cols = [
  { header: 'ID', accessor: (item: IncidentLog) => item.id },
  { header: 'Incident Title', accessor: (item: IncidentLog) => item.title },
  {
    header: 'Severity',
    accessor: (item: IncidentLog) => (
      <Badge variant={item.severity === 'CRITICAL' ? 'danger' : item.severity === 'HIGH' ? 'warning' : 'primary'}>
        {item.severity}
      </Badge>
    )
  },
  { header: 'Status', accessor: (item: IncidentLog) => item.status },
  { header: 'Logged Time', accessor: (item: IncidentLog) => item.timestamp }
];

const mockIncidents: IncidentLog[] = [
  { id: 'INC-2026-001', title: 'Gas leakage near storage valve v10', severity: 'CRITICAL', status: 'INVESTIGATING', timestamp: '2026-07-20 18:04' },
  { id: 'INC-2026-002', title: 'Chemical spill during transfer in warehouse', severity: 'HIGH', status: 'MITIGATED', timestamp: '2026-07-19 12:30' }
];

export const Incidents: React.FC = () => {
  return (
    <div className="flex flex-col gap-4 w-full">
      <div className="flex justify-between items-center">
        <h2 className="text-lg font-bold text-text">Incidents Registry</h2>
      </div>
      <DataTable columns={cols} data={mockIncidents} />
    </div>
  );
};
