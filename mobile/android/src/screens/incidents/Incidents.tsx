import React, { useState } from 'react';
import { Button } from '@prahari/components/Button.tsx';
import { Input } from '@prahari/components/Input.tsx';
import { syncEngine } from '../../sync/engine.ts';

export const Incidents: React.FC = () => {
  const [title, setTitle] = useState('');
  const [loading, setLoading] = useState(false);
  const [gpsTag, setGpsTag] = useState('');
  const [photoAdded, setPhotoAdded] = useState(false);

  const fetchGPSCoordinates = () => {
    setGpsTag('Lat: 28.6139° N, Lon: 77.2090° E');
  };

  const handleReport = async () => {
    if (!title) return;
    setLoading(true);

    const record = {
      id: `INC-OFFLINE-${Date.now()}`,
      title,
      gps: gpsTag,
      photoAttached: photoAdded,
      timestamp: new Date().toISOString()
    };

    await syncEngine.queueOfflineIncident(record);
    setLoading(false);
    setTitle('');
    setGpsTag('');
    setPhotoAdded(false);
    alert('Incident saved offline! Will sync automatically when network is available.');
  };

  return (
    <div className="flex flex-col gap-6 w-full max-w-lg mx-auto p-4 animate-in slide-in-from-bottom-2 duration-300">
      <h2 className="text-lg font-bold text-text">Rapid Incident Logger</h2>

      <div className="flex flex-col gap-4 border border-border bg-surface rounded-lg p-5">
        <Input
          label="Incident Description"
          placeholder="E.g., Spill in warehouse block B"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />

        <div className="flex flex-col gap-2">
          <span className="text-xs font-semibold text-text/80">Location Tag</span>
          <div className="flex gap-2">
            <Button variant="secondary" size="sm" onClick={fetchGPSCoordinates}>Fetch GPS coordinates</Button>
            {gpsTag && <span className="text-xs text-success self-center font-semibold">{gpsTag}</span>}
          </div>
        </div>

        <div className="flex flex-col gap-2">
          <span className="text-xs font-semibold text-text/80">Evidence Attachment</span>
          <div className="flex gap-2">
            <Button variant="secondary" size="sm" onClick={() => setPhotoAdded(true)}>Launch Camera Capture</Button>
            {photoAdded && <span className="text-xs text-success self-center font-semibold">Photo Attached ✅</span>}
          </div>
        </div>

        <div className="border-t border-border pt-4 flex justify-end gap-2 mt-2">
          <Button variant="danger" onClick={handleReport} isLoading={loading} disabled={!title}>
            Queue Offline Incident
          </Button>
        </div>
      </div>
    </div>
  );
};
