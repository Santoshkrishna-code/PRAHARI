import React, { useState } from 'react';
import { Button } from '@prahari/components/Button.tsx';
import { Card } from '@prahari/components/Card.tsx';
import { Badge } from '@prahari/components/Badge.tsx';

export const Permits: React.FC = () => {
  const [signature, setSignature] = useState(false);
  const [approved, setApproved] = useState(false);

  return (
    <div className="flex flex-col gap-6 w-full max-w-lg mx-auto p-4 animate-in slide-in-from-bottom-2 duration-300">
      <h2 className="text-xl font-bold tracking-tight text-text">Verify dispatched Permit</h2>

      <Card title="Permit Details: PRM-002">
        <div className="flex flex-col gap-3">
          <div className="flex justify-between items-center text-xs border-b border-border pb-2">
            <span className="font-bold">Permit Type</span>
            <span>Confined Space Entry</span>
          </div>
          <div className="flex justify-between items-center text-xs border-b border-border pb-2">
            <span className="font-bold">Location</span>
            <span>Reactor Building Unit B</span>
          </div>
          <div className="flex justify-between items-center text-xs pb-2">
            <span className="font-bold">Status</span>
            <Badge variant={approved ? 'success' : 'warning'}>{approved ? 'APPROVED' : 'PENDING SIGNATURE'}</Badge>
          </div>
        </div>
      </Card>

      {!approved && (
        <Card title="Biometric Authentication Signatures">
          <div className="flex flex-col gap-4">
            <div className="h-32 border border-dashed border-border rounded flex items-center justify-center bg-background text-xs text-muted">
              {signature ? '[Touch ID/Face ID Signature Verified]' : 'Confirm Credentials to Apply Signature'}
            </div>
            <div className="flex gap-2">
              <Button variant="secondary" onClick={() => setSignature(true)}>FaceID Signature</Button>
              <Button variant="success" disabled={!signature} onClick={() => setApproved(true)}>Approve & Dispatch</Button>
            </div>
          </div>
        </Card>
      )}
    </div>
  );
};
