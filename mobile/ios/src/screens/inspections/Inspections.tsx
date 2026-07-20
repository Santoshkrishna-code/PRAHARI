import React, { useState } from 'react';
import { Card } from '@prahari/components/Card.tsx';
import { Button } from '@prahari/components/Button.tsx';
import { Switch } from '@prahari/components/Switch.tsx';

export const Inspections: React.FC = () => {
  const [lotoVerified, setLotoVerified] = useState(false);
  const [gasVerified, setGasVerified] = useState(false);
  const [harnessVerified, setHarnessVerified] = useState(false);

  const handleFinish = () => {
    alert('Inspection checklist completed successfully!');
  };

  return (
    <div className="flex flex-col gap-6 w-full max-w-lg mx-auto p-4 bg-background">
      <h2 className="text-xl font-bold tracking-tight text-text">Asset Safety Inspection</h2>

      <Card title="Pre-Job Checklist: Storage Vessel Tank-01">
        <div className="flex flex-col gap-4">
          <Switch checked={lotoVerified} onChange={setLotoVerified} label="LOTO Tag applied and locked" />
          <Switch checked={gasVerified} onChange={setGasVerified} label="Gas level under target thresholds" />
          <Switch checked={harnessVerified} onChange={setHarnessVerified} label="Fall protection harness secure" />
          
          <Button variant="success" className="mt-4" onClick={handleFinish} disabled={!(lotoVerified && gasVerified && harnessVerified)}>
            Complete Inspection
          </Button>
        </div>
      </Card>
    </div>
  );
};
