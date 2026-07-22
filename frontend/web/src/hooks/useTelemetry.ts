import { useState, useEffect } from 'react';
import { HealthStatus, TelemetryPoint } from '../types';
import { realtimeApi } from '../services/apiClient';

export function useHealth(ms = 5000): HealthStatus {
  const [h, setH] = useState<HealthStatus>({ status: 'Operational', lat: 142, ts: '--' });

  useEffect(() => {
    let on = true;
    const go = async () => {
      const data = await realtimeApi.getHealth();
      if (on) {
        setH({
          status: data.status,
          lat: data.latencyMs,
          ts: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        });
      }
    };
    go();
    const iv = setInterval(go, ms);
    return () => { on = false; clearInterval(iv); };
  }, [ms]);

  return h;
}

export function useTelemetry(ms = 2000): TelemetryPoint[] {
  const [b, setB] = useState<TelemetryPoint[]>(() => Array.from({ length: 60 }, (_, i) => ({
    t: new Date(Date.now() - (59 - i) * ms).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
    vib: +(8 + Math.random() * 5).toFixed(2),
    temp: +(82 + Math.random() * 14).toFixed(1),
    psi: +(220 + Math.random() * 30).toFixed(1),
    kw: +(310 + Math.random() * 40).toFixed(0),
    flow: +(78 + Math.random() * 10).toFixed(0),
  })));

  useEffect(() => {
    const iv = setInterval(() => setB(p => {
      const l = p[p.length - 1];
      return [...p.slice(1), {
        t: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        vib: +Math.max(4, Math.min(16, l.vib + (Math.random() - 0.48) * 0.6)).toFixed(2),
        temp: +Math.max(70, Math.min(110, l.temp + (Math.random() - 0.45) * 0.5)).toFixed(1),
        psi: +Math.max(200, Math.min(280, l.psi + (Math.random() - 0.48) * 1.5)).toFixed(1),
        kw: +Math.max(280, Math.min(380, +l.kw + (Math.random() - 0.5) * 4)).toFixed(0),
        flow: +Math.max(70, Math.min(100, +l.flow + (Math.random() - 0.5) * 1.2)).toFixed(0),
      }];
    }), ms);
    return () => clearInterval(iv);
  }, [ms]);

  return b;
}
