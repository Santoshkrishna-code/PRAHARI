import React, { useState } from 'react';
import { Button } from '@prahari/components/Button.tsx';
import { Input } from '@prahari/components/Input.tsx';
import { DataTable } from '@prahari/components/DataTable.tsx';
import { Badge } from '@prahari/components/Badge.tsx';

interface Permit {
  id: string;
  type: string;
  applicant: string;
  status: string;
}

const cols = [
  { header: 'ID', accessor: (item: Permit) => item.id },
  { header: 'Permit Work Type', accessor: (item: Permit) => item.type },
  { header: 'Applicant', accessor: (item: Permit) => item.applicant },
  { header: 'Status', accessor: (item: Permit) => <Badge variant={item.status === 'APPROVED' ? 'success' : 'warning'}>{item.status}</Badge> }
];

export const Permits: React.FC = () => {
  const [step, setStep] = useState(1);
  const [permits, setPermits] = useState<Permit[]>([
    { id: 'PRM-001', type: 'Hot Work - Welding', applicant: 'Operator Santosh', status: 'APPROVED' },
    { id: 'PRM-002', type: 'Confined Space Entry', applicant: 'Auditor Rahul', status: 'PENDING' }
  ]);
  const [workType, setWorkType] = useState('');

  const handleCreate = () => {
    if (!workType) return;
    const newPermit: Permit = {
      id: `PRM-00${permits.length + 1}`,
      type: workType,
      applicant: 'Operator Santosh',
      status: 'PENDING'
    };
    setPermits([...permits, newPermit]);
    setWorkType('');
    setStep(1);
  };

  return (
    <div className="flex flex-col gap-6 w-full">
      <div className="flex justify-between items-center border-b border-border pb-4">
        <h2 className="text-lg font-bold text-text">Safe Work Permits</h2>
        <Button onClick={() => setStep(2)}>Request Permit Wizard</Button>
      </div>

      {step === 1 ? (
        <DataTable columns={cols} data={permits} />
      ) : (
        <div className="p-6 border border-border bg-surface rounded-lg max-w-md animate-in slide-in-from-bottom-2 duration-300">
          <h3 className="text-sm font-bold text-text mb-4">New Permit Request Stepper (Step 2/2)</h3>
          <div className="flex flex-col gap-4">
            <Input
              label="Work Description / Area Target"
              placeholder="E.g., Tank-01 hot welding"
              value={workType}
              onChange={(e) => setWorkType(e.target.value)}
            />
            <div className="flex justify-end gap-2 mt-2">
              <Button variant="secondary" onClick={() => setStep(1)}>Cancel</Button>
              <Button variant="success" onClick={handleCreate}>Submit Permit</Button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
