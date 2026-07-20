import React, { useState } from 'react';
import { Card } from '@prahari/components/Card.tsx';
import { KPICard } from '@prahari/components/KPICard.tsx';
import { Button } from '@prahari/components/Button.tsx';

export const TwinCanvas: React.FC = () => {
  const [pressure, setPressure] = useState(12.4);
  const [simStatus, setSimStatus] = useState('READY');

  const triggerSimulation = () => {
    setSimStatus('RUNNING');
    setTimeout(() => {
      setPressure(22.8);
      setSimStatus('FINISHED');
    }, 1000);
  };

  return (
    <div className="flex flex-col gap-6 w-full">
      <div className="flex justify-between items-center">
        <h2 className="text-lg font-bold text-text">Digital Twin Plant Topology canvas</h2>
        <div className="flex gap-2">
          <Button variant="secondary" onClick={() => setPressure(12.4)}>Reset Sensors State</Button>
          <Button variant="success" onClick={triggerSimulation} isLoading={simStatus === 'RUNNING'}>
            Trigger Cascade Pressure Simulation
          </Button>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <KPICard title="DC-101 Pressure binding" value={`${pressure.toFixed(1)} bar`} change={pressure > 20 ? 'OVER LIMIT WARNING' : 'NORMAL RANGE'} trend={pressure > 20 ? 'up' : 'neutral'} />
        <KPICard title="Active Simulation scenario" value={simStatus} change="Run ID: sim-991" trend="neutral" />
        <KPICard title="Graph query path latency" value="6.5 ms" change="Neo4j graph client good" trend="up" />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2">
          <Card title="Interactive Facility Layout explorer">
            <div className="aspect-video bg-background border border-border rounded flex flex-col items-center justify-center p-6 text-text/60 gap-4">
              <span className="text-sm font-semibold">[Unified Plant Topology Map Graph canvas]</span>
              <div className="flex gap-4">
                <span className="text-xs border border-border rounded px-2.5 py-1 bg-surface">[Building A - Valve v102]</span>
                <span className="text-xs border border-border rounded px-2.5 py-1 bg-surface">[Area B - Tank 01]</span>
              </div>
            </div>
          </Card>
        </div>

        <div>
          <Card title="Live Telemetry status logs">
            <div className="flex flex-col gap-3">
              <div className="flex justify-between items-center text-xs border-b border-border pb-2">
                <span className="font-bold">Sensor: TI-101</span>
                <span>45.2 °C (GOOD)</span>
              </div>
              <div className="flex justify-between items-center text-xs border-b border-border pb-2">
                <span className="font-bold">Sensor: PI-202</span>
                <span>3.4 bar (GOOD)</span>
              </div>
            </div>
          </Card>
        </div>
      </div>
    </div>
  );
};
