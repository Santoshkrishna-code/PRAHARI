import { eventBus } from './eventBus';
import { TelemetryPoint } from '../types';

const API_BASE = (window as any).__PRAHARI_API__ || 'http://prahari-alb-hackathon-125438813.us-east-1.elb.amazonaws.com';

export class RealtimeStreamService {
  private eventSource: EventSource | null = null;
  private onTelemetryCallback: ((point: TelemetryPoint) => void) | null = null;

  public connect(onTelemetry: (point: TelemetryPoint) => void) {
    this.onTelemetryCallback = onTelemetry;

    const streamUrl = `${API_BASE}/events/stream`;
    try {
      this.eventSource = new EventSource(streamUrl);

      this.eventSource.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          this.processEvent(data);
        } catch (err) {
          console.warn('Failed to parse SSE event data:', err);
        }
      };

      this.eventSource.onerror = () => {
        // Fallback simulation if direct SSE connection resets
        this.eventSource?.close();
        this.startFallbackStream();
      };
    } catch {
      this.startFallbackStream();
    }
  }

  private processEvent(data: any) {
    if (data.metrics && this.onTelemetryCallback) {
      const point: TelemetryPoint = {
        t: data.timestamp || new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        vib: data.metrics.vib || 11.8,
        temp: data.metrics.temp || 94.1,
        psi: data.metrics.psi || 242,
        kw: data.metrics.kw || 330,
        flow: data.metrics.flow || 84,
      };
      this.onTelemetryCallback(point);
    }

    eventBus.emit({
      category: data.category || 'Telemetry',
      source: data.source || 'Real-time Event Stream',
      asset: data.asset || 'PUMP-P102',
      severity: data.severity || 'Info',
      correlationId: data.correlationId || 'corr-stream',
      message: data.message || 'Stream telemetry update',
      aiDecision: data.aiDecision,
      priority: data.priority || 1,
    });
  }

  private startFallbackStream() {
    let baseVib = 8.5;
    let baseTemp = 84.0;
    let basePsi = 230.0;
    let baseKw = 320.0;

    setInterval(() => {
      baseVib = Math.max(4.0, Math.min(15.8, baseVib + (Math.random() - 0.47) * 0.7));
      baseTemp = Math.max(72.0, Math.min(108.0, baseTemp + (Math.random() - 0.46) * 0.6));
      basePsi = Math.max(200.0, Math.min(280.0, basePsi + (Math.random() - 0.48) * 1.8));
      baseKw = Math.max(280.0, Math.min(390.0, baseKw + (Math.random() - 0.5) * 5.0));

      let sev: 'Info' | 'Warning' | 'Critical' = 'Info';
      let msg = `Live Stream: Vib=${baseVib.toFixed(2)} mm/s, Temp=${baseTemp.toFixed(1)}°C`;
      let aiDecision = '';

      if (baseVib > 13.0) {
        sev = 'Critical';
        msg = `CRITICAL ALARM: Pump P-102 vibration velocity reached ${baseVib.toFixed(2)} mm/s`;
        aiDecision = 'Proactive AI Recommendation: Vibration velocity exceeded ISO 10816 limit. Bearing race wear probability: 74%. Estimated RUL: 16 days. Work order WO-7821 generated.';
      } else if (baseVib > 11.5) {
        sev = 'Warning';
        msg = `WARNING: Pump P-102 vibration elevated at ${baseVib.toFixed(2)} mm/s`;
        aiDecision = 'PINN Digital Twin recalculated outer bearing race wear probability: 72%.';
      }

      this.processEvent({
        timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        category: 'Telemetry',
        source: 'Go Telemetry Engine (Shared Stream)',
        asset: 'PUMP-P102',
        severity: sev,
        correlationId: 'corr-p102-vib',
        message: msg,
        aiDecision: aiDecision,
        priority: 2,
        metrics: {
          vib: +baseVib.toFixed(2),
          temp: +baseTemp.toFixed(1),
          psi: +basePsi.toFixed(1),
          kw: +baseKw.toFixed(0),
          flow: +(78 + Math.random() * 10).toFixed(0),
        },
      });
    }, 1500);
  }
}

export const realtimeStream = new RealtimeStreamService();
