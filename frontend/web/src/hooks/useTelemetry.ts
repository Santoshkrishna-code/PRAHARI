import { useState, useEffect } from 'react';
import { HealthStatus, TelemetryPoint } from '../types';
import { realtimeApi } from '../services/apiClient';
import { realtimeStream } from '../services/websocket';

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

export function useTelemetry(): TelemetryPoint[] {
  const [b, setB] = useState<TelemetryPoint[]>(() => Array.from({ length: 60 }, (_, i) => ({
    t: new Date(Date.now() - (59 - i) * 1500).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
    vib: +(8 + Math.random() * 5).toFixed(2),
    temp: +(82 + Math.random() * 14).toFixed(1),
    psi: +(220 + Math.random() * 30).toFixed(1),
    kw: +(310 + Math.random() * 40).toFixed(0),
    flow: +(78 + Math.random() * 10).toFixed(0),
  })));

  useEffect(() => {
    realtimeStream.connect((point: TelemetryPoint) => {
      setB(prev => [...prev.slice(1), point]);
    });
  }, []);

  return b;
}
