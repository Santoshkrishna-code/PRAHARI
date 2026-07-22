import { eventBus } from './eventBus';

export type WorkflowStepState = 'Pending' | 'In Progress' | 'Completed' | 'Failed';

export interface WorkflowStep {
  id: string;
  name: string;
  actor: string;
  state: WorkflowStepState;
  timestamp?: string;
  details?: string;
}

export interface EnterpriseWorkflow {
  workflowId: string;
  name: string;
  asset: string;
  correlationId: string;
  status: 'Active' | 'Completed' | 'Escalated';
  steps: WorkflowStep[];
}

class WorkflowEngineService {
  private activeWorkflows: EnterpriseWorkflow[] = [
    {
      workflowId: 'wf-p102-anomaly',
      name: 'Pump P-102 High Vibration Safety Recovery Process',
      asset: 'PUMP-P102',
      correlationId: 'corr-p102-vib',
      status: 'Active',
      steps: [
        { id: 'step-1', name: 'Sensor Anomaly Ingestion', actor: 'Vibration Probe VP-102', state: 'Completed', timestamp: '09:44:12', details: 'Vibration velocity crossed 11.8 mm/s threshold' },
        { id: 'step-2', name: 'AI Supervisor Verification', actor: 'Supervisor Agent', state: 'Completed', timestamp: '09:44:18', details: '5-Whys RCA verified bearing race wear (72% probability)' },
        { id: 'step-3', name: 'Risk Score Calculation', actor: 'Risk Agent', state: 'Completed', timestamp: '09:44:32', details: 'Severity 4/5 × Likelihood 3/5 = Risk Index 18/25 (HIGH)' },
        { id: 'step-4', name: 'Maintenance Work Order Generation', actor: 'CMMS Engine', state: 'Completed', timestamp: '09:45:01', details: 'Issued Work Order WO-7821 (Bearing replacement)' },
        { id: 'step-5', name: 'Safety Permit Isolation Lock', actor: 'Safety Officer / Permit Agent', state: 'Completed', timestamp: '09:46:11', details: 'Verified PTW-8902 LOTO lock on Valve V-88' },
        { id: 'step-6', name: 'Maintenance Execution', actor: 'Mechanical Maintenance Crew', state: 'In Progress', timestamp: '09:47:00', details: 'Bearing replacement underway in Zone DC-101' },
        { id: 'step-7', name: 'Return Asset to Service', actor: 'Plant Operations Manager', state: 'Pending' },
        { id: 'step-8', name: 'Executive Compliance Report Update', actor: 'Executive Agent', state: 'Pending' },
      ],
    },
  ];

  public getWorkflows(): EnterpriseWorkflow[] {
    return this.activeWorkflows;
  }

  public advanceStep(workflowId: string, stepId: string, details?: string) {
    const wf = this.activeWorkflows.find(w => w.workflowId === workflowId);
    if (!wf) return;

    const stepIdx = wf.steps.findIndex(s => s.id === stepId);
    if (stepIdx !== -1) {
      wf.steps[stepIdx].state = 'Completed';
      wf.steps[stepIdx].timestamp = new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' });
      if (details) wf.steps[stepIdx].details = details;

      if (stepIdx + 1 < wf.steps.length) {
        wf.steps[stepIdx + 1].state = 'In Progress';
        wf.steps[stepIdx + 1].timestamp = new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' });
      } else {
        wf.status = 'Completed';
      }

      eventBus.emit({
        category: 'Safety',
        source: 'Workflow Engine',
        asset: wf.asset,
        severity: 'Info',
        correlationId: wf.correlationId,
        message: `Advanced workflow "${wf.name}": Step "${wf.steps[stepIdx].name}" completed`,
        priority: 2,
      });
    }
  }
}

export const workflowEngine = new WorkflowEngineService();
