import React, { useState } from 'react';
import { Card } from '@prahari/components/Card.tsx';
import { Badge } from '@prahari/components/Badge.tsx';
import { Button } from '@prahari/components/Button.tsx';

export const VisionConsole: React.FC = () => {
  const [alerts, setAlerts] = useState<string[]>([
    'Cam-001: Banned chemical valve v10 bypass check observed',
    'Cam-002: Safety zone restricted area intrusion detected'
  ]);

  const triggerSim = () => {
    setAlerts((prev) => [
      `Cam-00${Math.floor(Math.random() * 5) + 1}: Live model detected safety vest/helmet missing`,
      ...prev
    ]);
  };

  return (
    <div className="flex flex-col gap-6 w-full">
      <div className="flex justify-between items-center">
        <h2 className="text-lg font-bold text-text">Perception Camera Stream Console</h2>
        <Button onClick={triggerSim}>Simulate Bounding Box Warning</Button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {/* Stream Placeholder 1 */}
        <Card title="Live Stream: Cam-001 - Reactor Gate">
          <div className="aspect-video bg-black flex items-center justify-center text-white/50 text-xs font-semibold rounded relative">
            <div className="absolute top-2 left-2 flex gap-1">
              <Badge variant="success">RTSP Pulling</Badge>
              <Badge variant="primary">PPE Model v1.2</Badge>
            </div>
            [Reactor Gate Feed Placeholder]
          </div>
        </Card>

        {/* Stream Placeholder 2 */}
        <Card title="Live Stream: Cam-002 - Warehouse Storage">
          <div className="aspect-video bg-black flex items-center justify-center text-white/50 text-xs font-semibold rounded relative">
            <div className="absolute top-2 left-2 flex gap-1">
              <Badge variant="success">WebRTC streaming</Badge>
              <Badge variant="primary">Spill Model v2.0</Badge>
            </div>
            [Warehouse Storage Feed Placeholder]
          </div>
        </Card>
      </div>

      <Card title="Model Inference Live Violation Alerts Feed">
        <div className="flex flex-col gap-2 max-h-[250px] overflow-y-auto">
          {alerts.map((al, idx) => (
            <div key={idx} className="p-3 bg-danger/10 border border-danger/20 rounded-md text-sm text-danger font-semibold flex items-center justify-between">
              <span>{al}</span>
              <Badge variant="danger">Critical</Badge>
            </div>
          ))}
        </div>
      </Card>
    </div>
  );
};
