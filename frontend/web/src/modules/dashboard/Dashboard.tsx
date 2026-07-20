import React from 'react';
import { KPICard } from '@prahari/components/KPICard.tsx';
import { Card } from '@prahari/components/Card.tsx';
import { Badge } from '@prahari/components/Badge.tsx';

export const Dashboard: React.FC = () => {
  return (
    <div className="flex flex-col gap-6 w-full">
      {/* Overview Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <KPICard title="Active Safe Work Permits" value="28 Permits" change="+6 since Monday" trend="up" />
        <KPICard title="Unresolved Incidents" value="2 Critical" change="-1 vs yesterday" trend="down" />
        <KPICard title="Compliance Target Score" value="98.5%" change="+0.4% this quarter" trend="up" />
        <KPICard title="System Sensors Status" value="ONLINE" change="Telemetry ping <2ms" trend="neutral" />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        {/* Active alerts log */}
        <Card title="Live AI Workspace & Vision Warnings">
          <div className="flex flex-col gap-3">
            <div className="flex items-center justify-between p-3 bg-danger/10 border border-danger/20 rounded-md">
              <div className="flex flex-col gap-0.5">
                <span className="text-sm font-bold text-danger">Safety Alert: PPE Intrusion Detected</span>
                <span className="text-xs text-text/80">Camera Cam-002 Reactor Area B - No Helmet observed.</span>
              </div>
              <Badge variant="danger">High</Badge>
            </div>

            <div className="flex items-center justify-between p-3 bg-warning/10 border border-warning/20 rounded-md">
              <div className="flex flex-col gap-0.5">
                <span className="text-sm font-bold text-warning">Digital Twin Simulation Finished</span>
                <span className="text-xs text-text/80">Cascading failure simulation DC-101 pipeline complete.</span>
              </div>
              <Badge variant="warning">Medium</Badge>
            </div>
          </div>
        </Card>

        {/* Quick actions panel */}
        <Card title="Operational Quick Actions">
          <div className="grid grid-cols-2 gap-3">
            <button className="p-4 border border-border rounded-md hover:bg-primary hover:text-background text-left transition-colors flex flex-col gap-1">
              <span className="text-sm font-bold">Request Permit</span>
              <span className="text-xs text-muted">Initialize hot work workflow stepper.</span>
            </button>
            <button className="p-4 border border-border rounded-md hover:bg-primary hover:text-background text-left transition-colors flex flex-col gap-1">
              <span className="text-sm font-bold">Log Safety Observation</span>
              <span className="text-xs text-muted">Register hazard or near-miss observation.</span>
            </button>
          </div>
        </Card>
      </div>
    </div>
  );
};
