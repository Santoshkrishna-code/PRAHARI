import React, { useState } from 'react';
import { Button } from '@prahari/components/Button.tsx';
import { Card } from '@prahari/components/Card.tsx';
import { Badge } from '@prahari/components/Badge.tsx';

export const Permits: React.FC = () => {
  const [signature, setSignature] = useState(false);
  const [approved, setApproved] = useState(false);

  return (
    <div className="flex flex-col gap-6 w-full max-w-lg mx-auto p-4 animate-in slide-in-from-bottom-2 duration-300">
      <h2 className="text-lg font-bold text-text">Review Active Permit</h2>

      <Card title="Permit Details: PRM-001">
        <div className="flex flex-col gap-3">
          <div className="flex justify-between items-center text-xs border-b border-border pb-2">
            <span className="font-bold">Permit Type</span>
            <span>Hot Work - Pipe Welding</span>
          </div>
          <div className="flex justify-between items-center text-xs border-b border-border pb-2">
            <span className="font-bold">Area Location</span>
            <span>Reactor Building Unit A</span>
          </div>
          <div className="flex justify-between items-center text-xs pb-2">
            <span className="font-bold">Status</span>
            <Badge variant={approved ? 'success' : 'warning'}>{approved ? 'APPROVED' : 'PENDING SIGNATURE'}</Badge>
          </div>
        </div>
      </Card>

      {!approved && (
        <Card title="Digital Verification Signature Required">
          <div className="flex flex-col gap-4">
            <div className="h-32 border border-dashed border-border rounded flex items-center justify-center bg-background text-xs text-muted">
              {signature ? '[Operator Santosh Digital Signature Verified]' : 'Tap Sign to Apply Signature'}
            </div>
            <div className="flex gap-2">
              <Button variant="secondary" onClick={() => setSignature(true)}>Sign</Button>
              <Button variant="success" disabled={!signature} onClick={() => setApproved(true)}>Approve & Dispatch</Button>
            </div>
          </div>
        </Card>
      )}
    </div>
  );
};
