import { UserRole } from '../types/auth';

export type EventCategory =
  | 'Telemetry'
  | 'Asset'
  | 'Maintenance'
  | 'Safety'
  | 'Permit'
  | 'Incident'
  | 'Vision'
  | 'AI'
  | 'Notification'
  | 'Authentication'
  | 'Audit'
  | 'System';

export interface EnterpriseEvent {
  eventId: string;
  timestamp: string;
  category: EventCategory;
  source: string;
  asset?: string;
  plant?: string;
  severity: 'Info' | 'Warning' | 'Critical';
  correlationId: string;
  user?: string;
  message: string;
  aiDecision?: string;
  auditRecord?: string;
  priority: number;
}

type EventListener = (event: EnterpriseEvent) => void;

class EventBusService {
  private listeners: EventListener[] = [];
  private eventHistory: EnterpriseEvent[] = [];
  private deadLetterQueue: EnterpriseEvent[] = [];

  constructor() {
    // Seed initial audit events
    const initialEvents: EnterpriseEvent[] = [
      {
        eventId: 'evt-001',
        timestamp: new Date(Date.now() - 600000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        category: 'System',
        source: 'AWS_ALB',
        plant: 'Plant Alpha (Gulf Coast)',
        severity: 'Info',
        correlationId: 'corr-init-01',
        message: 'Application Load Balancer health check passed — 142ms latency',
        priority: 1,
      },
      {
        eventId: 'evt-002',
        timestamp: new Date(Date.now() - 300000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        category: 'Telemetry',
        source: 'Vibration Probe VP-102',
        asset: 'PUMP-P102',
        plant: 'Plant Alpha (Gulf Coast)',
        severity: 'Warning',
        correlationId: 'corr-p102-vib',
        message: 'Vibration velocity reading 11.8 mm/s — approaching ISO 10816 limit',
        priority: 2,
      },
      {
        eventId: 'evt-003',
        timestamp: new Date(Date.now() - 120000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        category: 'AI',
        source: 'Supervisor Agent',
        asset: 'PUMP-P102',
        plant: 'Plant Alpha (Gulf Coast)',
        severity: 'Warning',
        correlationId: 'corr-p102-vib',
        message: 'Dispatched Risk Agent & Maintenance Agent for 5-Whys RCA',
        aiDecision: 'Calculated 72% bearing race wear probability; RUL updated to 18 days',
        priority: 3,
      },
    ];
    this.eventHistory = initialEvents;
  }

  public subscribe(listener: EventListener): () => void {
    this.listeners.push(listener);
    return () => {
      this.listeners = this.listeners.filter(l => l !== listener);
    };
  }

  public emit(event: Omit<EnterpriseEvent, 'eventId' | 'timestamp'>): EnterpriseEvent {
    const fullEvent: EnterpriseEvent = {
      ...event,
      eventId: `evt-${Math.floor(1000 + Math.random() * 9000)}`,
      timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
    };

    this.eventHistory.unshift(fullEvent);
    if (this.eventHistory.length > 200) {
      this.eventHistory.pop();
    }

    this.listeners.forEach(listener => {
      try {
        listener(fullEvent);
      } catch (err) {
        console.error('Failed to process event in listener:', err);
        this.deadLetterQueue.push(fullEvent);
      }
    });

    return fullEvent;
  }

  public getHistory(): EnterpriseEvent[] {
    return this.eventHistory;
  }

  public getDeadLetters(): EnterpriseEvent[] {
    return this.deadLetterQueue;
  }
}

export const eventBus = new EventBusService();
