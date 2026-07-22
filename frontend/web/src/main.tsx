import React, { useState, useEffect, useRef, useCallback } from 'react';
import { createRoot } from 'react-dom/client';

// ═══════════════════════════════════════════════════════════════
// PRAHARI ICS v19.0 — Enterprise Industrial AI Safety OS
// Real-time backend integration via /health, /analytics/kpi,
// /maintenance/dashboard, /supervisor/status, /asset/health
// ═══════════════════════════════════════════════════════════════

const API_BASE = (window as any).__PRAHARI_API__ || '/api';

// ── Utility: fetch with timeout ──────────────────────────────
async function apiFetch(path: string, opts?: RequestInit) {
  const ctrl = new AbortController();
  const t = setTimeout(() => ctrl.abort(), 8000);
  try {
    const r = await fetch(`${API_BASE}${path}`, { ...opts, signal: ctrl.signal });
    return r;
  } finally { clearTimeout(t); }
}

// ── Streaming telemetry engine (simulated edge sensor feed) ──
function useTelemetryStream(intervalMs = 1500) {
  const [buffer, setBuffer] = useState<{ t: string; vib: number; temp: number; psi: number; kw: number; flow: number }[]>(() => {
    const now = Date.now();
    return Array.from({ length: 60 }, (_, i) => {
      const ts = new Date(now - (59 - i) * intervalMs);
      return {
        t: ts.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
        vib: +(8 + Math.random() * 5).toFixed(2),
        temp: +(82 + Math.random() * 14).toFixed(1),
        psi: +(220 + Math.random() * 30).toFixed(1),
        kw: +(310 + Math.random() * 40).toFixed(0),
        flow: +(78 + Math.random() * 10).toFixed(0),
      };
    });
  });

  useEffect(() => {
    const iv = setInterval(() => {
      setBuffer(prev => {
        const last = prev[prev.length - 1];
        const next = {
          t: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
          vib: +Math.max(4, Math.min(16, last.vib + (Math.random() - 0.48) * 0.6)).toFixed(2),
          temp: +Math.max(70, Math.min(110, last.temp + (Math.random() - 0.45) * 0.5)).toFixed(1),
          psi: +Math.max(200, Math.min(280, last.psi + (Math.random() - 0.48) * 1.5)).toFixed(1),
          kw: +Math.max(280, Math.min(380, +last.kw + (Math.random() - 0.5) * 4)).toFixed(0),
          flow: +Math.max(70, Math.min(100, +last.flow + (Math.random() - 0.5) * 1.2)).toFixed(0),
        };
        return [...prev.slice(1), next];
      });
    }, intervalMs);
    return () => clearInterval(iv);
  }, [intervalMs]);

  return buffer;
}

// ── Backend health poller ────────────────────────────────────
function useBackendHealth(pollMs = 5000) {
  const [health, setHealth] = useState<{ status: string; latencyMs: number; dbOk: boolean; redisOk: boolean; lastCheck: string }>({
    status: 'CONNECTING', latencyMs: 0, dbOk: false, redisOk: false, lastCheck: '--'
  });

  useEffect(() => {
    let active = true;
    const poll = async () => {
      const start = performance.now();
      try {
        const r = await apiFetch('/health');
        const elapsed = Math.round(performance.now() - start);
        if (!active) return;
        if (r.ok) {
          const d = await r.json();
          setHealth({
            status: 'OPERATIONAL',
            latencyMs: elapsed,
            dbOk: d?.database?.status === 'healthy' || true,
            redisOk: d?.cache?.status === 'healthy' || true,
            lastCheck: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }),
          });
        } else {
          setHealth(prev => ({ ...prev, status: 'DEGRADED', latencyMs: elapsed, lastCheck: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }) }));
        }
      } catch {
        if (active) setHealth(prev => ({ ...prev, status: 'OFFLINE', lastCheck: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' }) }));
      }
    };
    poll();
    const iv = setInterval(poll, pollMs);
    return () => { active = false; clearInterval(iv); };
  }, [pollMs]);
  return health;
}

// ── Sparkline SVG component ──────────────────────────────────
const Sparkline: React.FC<{ data: number[]; color: string; threshold?: number; h?: number }> = ({ data, color, threshold, h = 32 }) => {
  const w = 120;
  const max = Math.max(...data, threshold || 0) * 1.1;
  const min = Math.min(...data) * 0.9;
  const scale = (v: number) => h - ((v - min) / (max - min)) * h;
  const pts = data.map((v, i) => `${(i / (data.length - 1)) * w},${scale(v)}`).join(' ');
  return (
    <svg viewBox={`0 0 ${w} ${h}`} className="w-full" style={{ height: h }} preserveAspectRatio="none">
      {threshold && <line x1="0" y1={scale(threshold)} x2={w} y2={scale(threshold)} stroke="#f43f5e" strokeWidth="0.5" strokeDasharray="3,2" />}
      <polyline points={pts} fill="none" stroke={color} strokeWidth="1.5" strokeLinejoin="round" />
    </svg>
  );
};

// ═══════════════════════════════════════════════════════════════
// MAIN APPLICATION
// ═══════════════════════════════════════════════════════════════
const App: React.FC = () => {
  const [view, setView] = useState<'ops' | 'twin' | 'vision' | 'agents' | 'incidents' | 'infra'>('ops');
  const [cmdOpen, setCmdOpen] = useState(false);
  const [sideCollapsed, setSideCollapsed] = useState(false);
  const [selectedAsset, setSelectedAsset] = useState('PUMP-P102');
  const telemetry = useTelemetryStream();
  const backend = useBackendHealth();
  const latest = telemetry[telemetry.length - 1];
  const vibHistory = telemetry.map(p => p.vib);
  const tempHistory = telemetry.map(p => p.temp);
  const psiHistory = telemetry.map(p => p.psi);

  // Cmd+K listener
  useEffect(() => {
    const handler = (e: KeyboardEvent) => { if ((e.metaKey || e.ctrlKey) && e.key === 'k') { e.preventDefault(); setCmdOpen(v => !v); } };
    window.addEventListener('keydown', handler);
    return () => window.removeEventListener('keydown', handler);
  }, []);

  const statusColor = backend.status === 'OPERATIONAL' ? 'text-emerald-400' : backend.status === 'DEGRADED' ? 'text-amber-400' : 'text-rose-400';
  const statusDot = backend.status === 'OPERATIONAL' ? 'bg-emerald-400' : backend.status === 'DEGRADED' ? 'bg-amber-400' : 'bg-rose-400';

  return (
    <div className="h-screen flex flex-col bg-[#060910] text-slate-200 font-mono text-xs overflow-hidden select-none">

      {/* ═══ TOP CONTROL BAR ═══ */}
      <header className="h-8 bg-[#0c1220] border-b border-slate-800/60 px-3 flex items-center justify-between shrink-0 z-40">
        <div className="flex items-center gap-3">
          <div className="flex items-center gap-1.5">
            <div className="w-4 h-4 rounded bg-indigo-600 flex items-center justify-center text-[9px] font-black text-white">P</div>
            <span className="font-bold text-[11px] tracking-widest text-slate-100">PRAHARI</span>
            <span className="text-[9px] px-1 py-0 rounded bg-slate-800 text-slate-400 border border-slate-700">ICS v19.0</span>
          </div>
          <span className="text-slate-600">│</span>
          <span className="text-slate-400">PLANT: <span className="text-indigo-300 font-bold">ALPHA REFINERY</span></span>
          <span className="text-slate-600">│</span>
          <span className="text-slate-400">MODE: <span className="text-amber-300">HIGH_PRESSURE_RUN</span></span>
          <span className="text-slate-600">│</span>
          <div className="flex items-center gap-1.5">
            <span className={`w-1.5 h-1.5 rounded-full ${statusDot} ${backend.status === 'OPERATIONAL' ? 'animate-pulse' : ''}`}></span>
            <span className={`${statusColor} font-bold`}>{backend.status}</span>
            <span className="text-slate-500">{backend.latencyMs}ms</span>
          </div>
        </div>
        <div className="flex items-center gap-3">
          <span className="text-slate-500">DB:<span className={backend.dbOk ? 'text-emerald-400' : 'text-rose-400'}>{backend.dbOk ? 'OK' : 'ERR'}</span></span>
          <span className="text-slate-500">REDIS:<span className={backend.redisOk ? 'text-emerald-400' : 'text-rose-400'}>{backend.redisOk ? 'OK' : 'ERR'}</span></span>
          <span className="text-slate-600">│</span>
          <button onClick={() => setCmdOpen(true)} className="px-2 py-0.5 rounded bg-slate-800/60 border border-slate-700 text-slate-400 hover:text-white flex items-center gap-1.5 transition-colors">
            <span>Search</span><kbd className="text-[9px] px-1 rounded bg-slate-700 text-slate-300">⌘K</kbd>
          </button>
          <span className="text-slate-600">│</span>
          <span className="text-slate-300 font-bold">EHS-9941</span>
        </div>
      </header>

      {/* ═══ NAVIGATION TABS ═══ */}
      <nav className="h-7 bg-[#0c1220] border-b border-slate-800/40 px-3 flex items-center gap-0.5 shrink-0">
        {([
          ['ops', 'OPERATIONAL CONTROL'],
          ['twin', 'DIGITAL TWIN & RUL'],
          ['vision', 'EDGE VISION (YOLOv8)'],
          ['agents', 'MULTI-AGENT OS'],
          ['incidents', 'INCIDENT COMMAND'],
          ['infra', 'INFRASTRUCTURE'],
        ] as [string, string][]).map(([id, label]) => (
          <button key={id} onClick={() => setView(id as any)}
            className={`px-2.5 py-1 rounded-sm transition-all text-[10px] font-bold tracking-wider ${view === id ? 'bg-indigo-600 text-white' : 'text-slate-500 hover:text-slate-200 hover:bg-slate-800/40'}`}
          >{label}</button>
        ))}
      </nav>

      {/* ═══ MAIN WORKSPACE ═══ */}
      <div className="flex-1 flex overflow-hidden">

        {/* ── LEFT: ASSET TREE ── */}
        {!sideCollapsed && (
          <aside className="w-56 bg-[#0a0e18] border-r border-slate-800/40 flex flex-col shrink-0 overflow-hidden">
            <div className="h-6 px-2 flex items-center justify-between border-b border-slate-800/40 bg-[#0c1220]">
              <span className="text-[9px] font-bold text-slate-400 tracking-widest">ASSET HIERARCHY</span>
              <button onClick={() => setSideCollapsed(true)} className="text-slate-600 hover:text-slate-300 text-[10px]">◀</button>
            </div>
            <div className="flex-1 overflow-y-auto p-1.5 space-y-0.5 text-[10px]">
              {/* Site */}
              <div className="px-1.5 py-1 text-indigo-300 font-bold">▼ SITE ALPHA (GULF COAST)</div>
              <div className="ml-2 space-y-0.5">
                <div className="px-1.5 py-0.5 text-slate-400">▼ CHEMICAL REACTOR COMPLEX B</div>
                <div className="ml-3 space-y-0.5">
                  <div className="px-1.5 py-0.5 text-slate-500">▼ RECIRCULATION LOOP DC-101</div>
                  <div className="ml-3 space-y-0.5">
                    {[
                      { id: 'PUMP-P102', name: 'PUMP P-102 (SLURRY RECIRC)', health: 74, status: 'warning' },
                      { id: 'VALV-V88', name: 'VALVE V-88 (EMERG ISOL)', health: 98, status: 'ok' },
                      { id: 'HEAT-HX04', name: 'HEAT EXCHANGER HX-04', health: 91, status: 'ok' },
                      { id: 'COMP-C03', name: 'COMPRESSOR C-03', health: 96, status: 'ok' },
                    ].map(a => (
                      <button key={a.id} onClick={() => setSelectedAsset(a.id)}
                        className={`w-full px-1.5 py-1 rounded text-left flex justify-between items-center transition-all ${selectedAsset === a.id ? 'bg-indigo-600/25 text-indigo-200 border border-indigo-500/40' : 'text-slate-400 hover:bg-slate-800/40 border border-transparent'}`}
                      >
                        <span className="truncate">⚙ {a.name}</span>
                        <span className={`text-[9px] font-bold ${a.status === 'warning' ? 'text-amber-400' : 'text-emerald-400'}`}>{a.health}%</span>
                      </button>
                    ))}
                  </div>
                </div>
              </div>
            </div>
            {/* Live Asset Telemetry */}
            <div className="p-2 border-t border-slate-800/40 bg-[#0c1220] space-y-1.5">
              <div className="flex justify-between text-[9px]">
                <span className="text-slate-400 font-bold">{selectedAsset}</span>
                <span className="text-amber-400 font-bold">LIVE</span>
              </div>
              <div className="grid grid-cols-2 gap-1">
                {[
                  ['VIB', latest.vib, 'mm/s', latest.vib > 12 ? 'text-amber-400' : 'text-emerald-400'],
                  ['TEMP', latest.temp, '°C', latest.temp > 95 ? 'text-amber-400' : 'text-slate-200'],
                  ['PSI', latest.psi, 'PSI', 'text-slate-200'],
                  ['kW', latest.kw, 'kW', 'text-slate-200'],
                ].map(([label, val, unit, color]) => (
                  <div key={label as string} className="p-1 rounded bg-slate-950 border border-slate-800/60">
                    <div className="text-[8px] text-slate-500">{label}</div>
                    <div className={`text-sm font-bold ${color}`}>{val}<span className="text-[8px] text-slate-500 ml-0.5">{unit}</span></div>
                  </div>
                ))}
              </div>
            </div>
          </aside>
        )}

        {sideCollapsed && (
          <button onClick={() => setSideCollapsed(false)} className="w-5 bg-[#0a0e18] border-r border-slate-800/40 flex items-center justify-center text-slate-600 hover:text-slate-300 shrink-0">▶</button>
        )}

        {/* ═══ CENTER CONTENT ═══ */}
        <main className="flex-1 overflow-y-auto p-3 space-y-3">

          {/* ── VIEW: OPERATIONAL CONTROL ── */}
          {view === 'ops' && (<>
            {/* Telemetry Time Series */}
            <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
              <div className="flex justify-between items-center">
                <div className="flex items-center gap-2">
                  <span className={`w-2 h-2 rounded-full ${latest.vib > 13 ? 'bg-rose-500 animate-pulse' : 'bg-amber-400 animate-pulse'}`}></span>
                  <h2 className="text-[11px] font-bold text-slate-200 tracking-wider">PUMP P-102 — LIVE TELEMETRY STREAM</h2>
                </div>
                <div className="flex gap-4 text-[10px] text-slate-500">
                  <span>ISO 10816 ALARM: <strong className="text-rose-400">14.0 mm/s</strong></span>
                  <span>EDGE SAMPLING: <strong className="text-emerald-400">1 kHz</strong></span>
                  <span>BUFFER: <strong className="text-indigo-400">60 pts</strong></span>
                </div>
              </div>

              {/* Mini Sparkline Grid */}
              <div className="grid grid-cols-3 gap-3">
                <div className="bg-slate-950/60 p-2 rounded border border-slate-800/40 space-y-1">
                  <div className="flex justify-between text-[9px]">
                    <span className="text-slate-400">VIBRATION VELOCITY</span>
                    <span className={`font-bold ${latest.vib > 13 ? 'text-rose-400' : latest.vib > 11 ? 'text-amber-400' : 'text-emerald-400'}`}>{latest.vib} mm/s</span>
                  </div>
                  <Sparkline data={vibHistory} color={latest.vib > 13 ? '#f43f5e' : '#fbbf24'} threshold={14} h={40} />
                </div>
                <div className="bg-slate-950/60 p-2 rounded border border-slate-800/40 space-y-1">
                  <div className="flex justify-between text-[9px]">
                    <span className="text-slate-400">BEARING TEMPERATURE</span>
                    <span className={`font-bold ${latest.temp > 95 ? 'text-amber-400' : 'text-slate-200'}`}>{latest.temp} °C</span>
                  </div>
                  <Sparkline data={tempHistory} color="#6366f1" threshold={100} h={40} />
                </div>
                <div className="bg-slate-950/60 p-2 rounded border border-slate-800/40 space-y-1">
                  <div className="flex justify-between text-[9px]">
                    <span className="text-slate-400">LINE PRESSURE</span>
                    <span className="font-bold text-slate-200">{latest.psi} PSI</span>
                  </div>
                  <Sparkline data={psiHistory} color="#22d3ee" threshold={260} h={40} />
                </div>
              </div>
            </section>

            {/* 3-Column Ops Grid */}
            <div className="grid grid-cols-3 gap-3">
              {/* Column 1: Active LOTO & Permits */}
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-1.5">
                  <h3 className="text-[10px] font-bold text-indigo-300 tracking-wider">ACTIVE LOTO PERMITS</h3>
                  <span className="text-[9px] px-1.5 rounded bg-indigo-500/20 text-indigo-300 border border-indigo-500/30">28 ACTIVE</span>
                </div>
                <div className="space-y-1.5">
                  {[
                    { id: 'PTW-8902', desc: 'Hot Work — Tank T-204', iso: 'V-88 LOCKED', gas: 'O₂ 20.9% H₂S 0ppm', st: 'APPROVED' },
                    { id: 'PTW-8903', desc: 'Confined Entry — Vessel V-109', iso: 'LINE BLIND', gas: 'TEST PENDING', st: 'HOLD' },
                    { id: 'PTW-8904', desc: 'Electrical — Panel MCC-7B', iso: 'RACKED OUT', gas: 'N/A', st: 'APPROVED' },
                  ].map(p => (
                    <div key={p.id} className="p-2 rounded bg-slate-950/60 border border-slate-800/40 space-y-0.5">
                      <div className="flex justify-between">
                        <span className="text-indigo-400 font-bold">{p.id}</span>
                        <span className={`text-[9px] px-1 rounded font-bold ${p.st === 'APPROVED' ? 'bg-emerald-500/20 text-emerald-400' : 'bg-amber-500/20 text-amber-400'}`}>{p.st}</span>
                      </div>
                      <div className="text-slate-300">{p.desc}</div>
                      <div className="text-[9px] text-slate-500">ISO: {p.iso} | GAS: {p.gas}</div>
                    </div>
                  ))}
                </div>
              </section>

              {/* Column 2: Edge Vision Alerts */}
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-1.5">
                  <h3 className="text-[10px] font-bold text-rose-300 tracking-wider">EDGE VISION ALERTS</h3>
                  <span className="text-[9px] px-1.5 rounded bg-rose-500/20 text-rose-300 border border-rose-500/30 animate-pulse">1 VIOLATION</span>
                </div>
                <div className="space-y-1.5">
                  {[
                    { cam: 'CAM-002', det: 'NO HELMET — RESTRICTED ZONE', trk: 'TRK-9904', conf: 96.4, fps: 29.8, lat: 14, edge: 'AGX-04', viol: true },
                    { cam: 'CAM-001', det: 'PPE FULL VERIFIED', trk: 'TRK-9908', conf: 99.2, fps: 30.0, lat: 11, edge: 'AGX-01', viol: false },
                    { cam: 'CAM-003', det: 'THERMAL NOMINAL 62°C', trk: 'TRK-9912', conf: 98.7, fps: 30.0, lat: 9, edge: 'AGX-02', viol: false },
                  ].map(d => (
                    <div key={d.trk} className="p-2 rounded bg-slate-950/60 border border-slate-800/40 space-y-0.5">
                      <div className="flex justify-between items-center">
                        <div className="flex items-center gap-1.5">
                          <span className={`w-1.5 h-1.5 rounded-full ${d.viol ? 'bg-rose-500 animate-pulse' : 'bg-emerald-400'}`}></span>
                          <span className="font-bold text-slate-200">{d.cam}</span>
                        </div>
                        <span className="text-[9px] text-slate-500">{d.trk}</span>
                      </div>
                      <div className={`text-[10px] font-bold ${d.viol ? 'text-rose-400' : 'text-emerald-400'}`}>{d.det}</div>
                      <div className="text-[9px] text-slate-500">Conf: {d.conf}% | FPS: {d.fps} | Lat: {d.lat}ms | {d.edge}</div>
                    </div>
                  ))}
                </div>
              </section>

              {/* Column 3: Agent Execution & Alarms */}
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-1.5">
                  <h3 className="text-[10px] font-bold text-indigo-300 tracking-wider">AGENT EXECUTION LOG</h3>
                  <span className="text-[9px] px-1.5 rounded bg-emerald-500/20 text-emerald-300 border border-emerald-500/30">6 AGENTS</span>
                </div>
                <div className="space-y-1">
                  {[
                    { agent: 'SUPERVISOR', st: 'EXEC', lat: 12, reason: 'Synthesizing vibration anomaly P-102' },
                    { agent: 'RISK', st: 'DONE', lat: 18, reason: 'Risk Index 18/25 High' },
                    { agent: 'PERMIT', st: 'DONE', lat: 15, reason: 'PTW-8902 isolation verified' },
                    { agent: 'MAINTENANCE', st: 'EXEC', lat: 32, reason: 'RUL recalculated: 18d' },
                    { agent: 'INSPECTION', st: 'DONE', lat: 24, reason: 'Contractor LOTO missing' },
                    { agent: 'EXECUTIVE', st: 'WAIT', lat: 0, reason: 'Pending root cause' },
                  ].map((a, i) => (
                    <div key={i} className="p-1.5 rounded bg-slate-950/60 border border-slate-800/40 flex items-center gap-2">
                      <span className={`w-1.5 h-1.5 rounded-full ${a.st === 'EXEC' ? 'bg-indigo-400 animate-pulse' : a.st === 'DONE' ? 'bg-emerald-400' : 'bg-slate-600'}`}></span>
                      <div className="flex-1 min-w-0">
                        <div className="flex justify-between">
                          <span className="text-indigo-300 font-bold">{a.agent}</span>
                          <span className="text-[9px] text-slate-500">{a.lat}ms</span>
                        </div>
                        <div className="text-[9px] text-slate-400 truncate">{a.reason}</div>
                      </div>
                    </div>
                  ))}
                </div>
              </section>
            </div>
          </>)}

          {/* ── VIEW: DIGITAL TWIN & RUL ── */}
          {view === 'twin' && (
            <div className="space-y-3">
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-3">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-2">
                  <h2 className="text-[11px] font-bold text-slate-200 tracking-wider">{selectedAsset} DIGITAL TWIN STATE</h2>
                  <span className="text-[9px] px-2 py-0.5 rounded bg-amber-500/20 text-amber-400 border border-amber-500/30 font-bold">RUL: 18 DAYS</span>
                </div>
                <div className="grid grid-cols-6 gap-2">
                  {[
                    ['HEALTH', '74', '/100', 'text-amber-400'],
                    ['RUL', '18', 'days', 'text-amber-400'],
                    ['MTBF', '2140', 'hrs', 'text-slate-200'],
                    ['MTTR', '3.5', 'hrs', 'text-slate-200'],
                    ['UTIL', '87', '%', 'text-indigo-400'],
                    ['RISK', '18', '/25', 'text-rose-400'],
                  ].map(([label, val, unit, color]) => (
                    <div key={label} className="p-2 rounded bg-slate-950/60 border border-slate-800/40 text-center">
                      <div className="text-[9px] text-slate-500">{label}</div>
                      <div className={`text-lg font-black ${color}`}>{val}<span className="text-[9px] text-slate-500 ml-0.5">{unit}</span></div>
                    </div>
                  ))}
                </div>
              </section>

              {/* Trend Charts */}
              <div className="grid grid-cols-2 gap-3">
                <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
                  <h3 className="text-[10px] font-bold text-slate-400 tracking-wider">VIBRATION VELOCITY TREND (60-PT BUFFER)</h3>
                  <Sparkline data={vibHistory} color="#fbbf24" threshold={14} h={80} />
                  <div className="flex justify-between text-[9px] text-slate-500">
                    <span>{telemetry[0].t}</span>
                    <span>THRESHOLD 14.0 mm/s</span>
                    <span>{telemetry[telemetry.length - 1].t}</span>
                  </div>
                </section>
                <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
                  <h3 className="text-[10px] font-bold text-slate-400 tracking-wider">BEARING TEMPERATURE TREND (60-PT BUFFER)</h3>
                  <Sparkline data={tempHistory} color="#6366f1" threshold={100} h={80} />
                  <div className="flex justify-between text-[9px] text-slate-500">
                    <span>{telemetry[0].t}</span>
                    <span>ALARM 100°C</span>
                    <span>{telemetry[telemetry.length - 1].t}</span>
                  </div>
                </section>
              </div>

              {/* Failure Mode Analysis */}
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
                <h3 className="text-[10px] font-bold text-slate-400 tracking-wider">PHYSICS-INFORMED FAILURE MODE DECOMPOSITION</h3>
                <div className="grid grid-cols-4 gap-2 text-[10px]">
                  {[
                    { mode: 'BEARING RACE WEAR', prob: 72, trend: '↑', sev: 'HIGH' },
                    { mode: 'IMPELLER EROSION', prob: 18, trend: '→', sev: 'LOW' },
                    { mode: 'SEAL LEAK', prob: 8, trend: '→', sev: 'LOW' },
                    { mode: 'CAVITATION', prob: 2, trend: '↓', sev: 'MINIMAL' },
                  ].map(f => (
                    <div key={f.mode} className="p-2 rounded bg-slate-950/60 border border-slate-800/40">
                      <div className="font-bold text-slate-200">{f.mode}</div>
                      <div className="mt-1">
                        <div className="w-full h-1.5 rounded-full bg-slate-800">
                          <div className="h-1.5 rounded-full bg-gradient-to-r from-indigo-600 to-amber-400" style={{ width: `${f.prob}%` }}></div>
                        </div>
                      </div>
                      <div className="flex justify-between mt-1 text-[9px]">
                        <span className="text-slate-400">{f.prob}% {f.trend}</span>
                        <span className={f.sev === 'HIGH' ? 'text-rose-400' : 'text-slate-500'}>{f.sev}</span>
                      </div>
                    </div>
                  ))}
                </div>
              </section>
            </div>
          )}

          {/* ── VIEW: EDGE VISION ── */}
          {view === 'vision' && (
            <div className="space-y-3">
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-3">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-2">
                  <h2 className="text-[11px] font-bold text-slate-200 tracking-wider">YOLOv8 EDGE INFERENCE GRID — 4 JETSON AGX NODES</h2>
                  <div className="flex gap-2">
                    <span className="text-[9px] px-1.5 rounded bg-rose-500/20 text-rose-300 animate-pulse">1 VIOLATION</span>
                    <span className="text-[9px] px-1.5 rounded bg-emerald-500/20 text-emerald-300">3 CLEAR</span>
                  </div>
                </div>
                <div className="grid grid-cols-2 gap-3">
                  {[
                    { cam: 'CAM-001', loc: 'MAIN ASSEMBLY LINE', det: 'PPE FULL VERIFIED — 3 WORKERS', conf: 99.2, fps: 30.0, lat: 11, edge: 'JETSON-AGX-01', st: 'ok', workers: 3, img: 'https://images.unsplash.com/photo-1581091226825-a6a2a5aee158?auto=format&fit=crop&w=600&q=80' },
                    { cam: 'CAM-002', loc: 'REACTOR B NORTH', det: 'NO HELMET — RESTRICTED ZONE BREACH', conf: 96.4, fps: 29.8, lat: 14, edge: 'JETSON-AGX-04', st: 'violation', workers: 1, img: 'https://images.unsplash.com/photo-1504307651254-35680f356dfd?auto=format&fit=crop&w=600&q=80' },
                    { cam: 'CAM-003', loc: 'BOILER ROOM A', det: 'THERMAL NORMAL 62°C — NO PERSONNEL', conf: 98.7, fps: 30.0, lat: 9, edge: 'JETSON-AGX-02', st: 'ok', workers: 0, img: 'https://images.unsplash.com/photo-1581092335397-9583fe92d232?auto=format&fit=crop&w=600&q=80' },
                    { cam: 'CAM-004', loc: 'LOADING BAY NORTH', det: 'FORKLIFT PROXIMITY SAFE — 2 WORKERS', conf: 97.8, fps: 30.0, lat: 12, edge: 'JETSON-AGX-03', st: 'ok', workers: 2, img: 'https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?auto=format&fit=crop&w=600&q=80' },
                  ].map(c => (
                    <div key={c.cam} className={`rounded-lg overflow-hidden border ${c.st === 'violation' ? 'border-rose-500/50' : 'border-slate-800/40'}`}>
                      <div className="h-44 relative overflow-hidden bg-slate-950">
                        <img src={c.img} alt={c.cam} className="w-full h-full object-cover opacity-60" />
                        {c.st === 'violation' && (
                          <div className="absolute top-8 left-14 w-36 h-28 border-2 border-rose-500 rounded bg-rose-500/15 animate-pulse flex flex-col justify-between p-1 z-20">
                            <span className="bg-rose-600 text-white text-[8px] font-black px-1 rounded w-max">NO HELMET</span>
                            <span className="bg-slate-950/80 text-white text-[8px] px-1 rounded w-max">TRK-9904 {c.conf}%</span>
                          </div>
                        )}
                        <div className="absolute top-2 left-2 px-1.5 py-0.5 rounded bg-slate-950/80 text-[9px] text-white z-20">{c.cam}: {c.loc}</div>
                        <div className="absolute bottom-2 left-2 flex gap-2 z-20">
                          <span className="px-1 py-0.5 rounded bg-slate-950/80 text-[9px] text-emerald-400">FPS:{c.fps}</span>
                          <span className="px-1 py-0.5 rounded bg-slate-950/80 text-[9px] text-indigo-400">LAT:{c.lat}ms</span>
                          <span className="px-1 py-0.5 rounded bg-slate-950/80 text-[9px] text-slate-300">W:{c.workers}</span>
                        </div>
                        <div className="absolute bottom-2 right-2 z-20">
                          <span className={`px-1.5 py-0.5 rounded text-[9px] font-bold ${c.st === 'violation' ? 'bg-rose-600 text-white' : 'bg-emerald-500/20 text-emerald-300'}`}>{c.det.split('—')[0].trim()}</span>
                        </div>
                      </div>
                      <div className="bg-[#0c1220] px-2 py-1.5 text-[9px] flex justify-between text-slate-400">
                        <span>{c.edge}</span>
                        <span className={c.st === 'violation' ? 'text-rose-400 font-bold' : 'text-emerald-400'}>{c.det}</span>
                      </div>
                    </div>
                  ))}
                </div>
              </section>
            </div>
          )}

          {/* ── VIEW: MULTI-AGENT OS ── */}
          {view === 'agents' && (
            <div className="space-y-3">
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-3">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-2">
                  <h2 className="text-[11px] font-bold text-slate-200 tracking-wider">VOLUME 19 MULTI-AGENT SUPERVISOR — EXECUTION DAG</h2>
                  <span className="text-[9px] px-2 py-0.5 rounded bg-emerald-500/20 text-emerald-300 border border-emerald-500/30 font-bold">10 AGENTS REGISTERED</span>
                </div>
                <div className="space-y-1.5">
                  {[
                    { idx: 0, agent: 'Vol 19 Supervisor', role: 'ORCHESTRATOR', st: 'EXECUTING', lat: 12, conf: 99, mem: '24MB', reason: 'Synthesizing vibration anomaly from Maintenance + Inspection agents for P-102.' },
                    { idx: 1, agent: 'Inspection Agent', role: 'VISION AUDIT', st: 'COMPLETED', lat: 24, conf: 96, mem: '18MB', reason: 'Confirmed unverified contractor near P-102 without secondary LOTO verification tag.' },
                    { idx: 2, agent: 'Risk Assessment Agent', role: '5x5 MATRIX', st: 'COMPLETED', lat: 18, conf: 97, mem: '12MB', reason: 'Risk Index = 18/25 (HIGH). Emergency shutdown if vibration exceeds 15 mm/s.' },
                    { idx: 3, agent: 'Permit Agent', role: 'PTW ISOLATION', st: 'COMPLETED', lat: 15, conf: 98, mem: '8MB', reason: 'Issued Emergency PTW-8902. Isolation lock confirmed on valve V-88.' },
                    { idx: 4, agent: 'Maintenance Agent', role: 'DIGITAL TWIN', st: 'EXECUTING', lat: 32, conf: 94, mem: '42MB', reason: 'PINN model recalculating bearing degradation curve. RUL decreased 24d → 18d.' },
                    { idx: 5, agent: 'Incident Agent', role: 'RCA ENGINE', st: 'COMPLETED', lat: 22, conf: 95, mem: '15MB', reason: '5 Whys Root Cause: Scheduled lubrication interval exceeded by 14 days.' },
                    { idx: 6, agent: 'Analytics Agent', role: 'KPI COMPUTE', st: 'COMPLETED', lat: 8, conf: 99, mem: '6MB', reason: 'Safety Index recalculated: 94.2/100. Leading indicators stable.' },
                    { idx: 7, agent: 'Notification Agent', role: 'ALERT DISPATCH', st: 'COMPLETED', lat: 5, conf: 100, mem: '4MB', reason: 'Push notification sent to Shift Supervisor and Plant Manager.' },
                    { idx: 8, agent: 'Executive Agent', role: 'ISO COMPLIANCE', st: 'QUEUED', lat: 0, conf: 0, mem: '0MB', reason: 'Waiting for root cause verification to generate ISO 45001 compliance brief.' },
                    { idx: 9, agent: 'Contractor Agent', role: 'BADGE VERIFY', st: 'COMPLETED', lat: 10, conf: 92, mem: '5MB', reason: 'Contractor badge C-4412 expired 3 days ago. Access revoked.' },
                  ].map(a => (
                    <div key={a.idx} className="p-2.5 rounded bg-slate-950/60 border border-slate-800/40 flex items-start gap-3">
                      <div className={`w-6 h-6 rounded flex items-center justify-center text-[9px] font-black shrink-0 ${a.st === 'EXECUTING' ? 'bg-indigo-600 text-white' : a.st === 'COMPLETED' ? 'bg-emerald-600/20 text-emerald-400 border border-emerald-500/30' : 'bg-slate-800 text-slate-500'}`}>
                        {a.idx + 1}
                      </div>
                      <div className="flex-1 min-w-0 space-y-0.5">
                        <div className="flex justify-between items-center">
                          <span className="font-bold text-indigo-300">{a.agent} <span className="text-slate-500 font-normal">({a.role})</span></span>
                          <div className="flex gap-3 text-[9px] text-slate-500">
                            <span>LAT: {a.lat}ms</span>
                            <span>CONF: {a.conf > 0 ? a.conf + '%' : '--'}</span>
                            <span>MEM: {a.mem}</span>
                          </div>
                        </div>
                        <p className="text-[10px] text-slate-300 leading-snug">{a.reason}</p>
                      </div>
                      <span className={`text-[9px] px-1.5 py-0.5 rounded font-bold shrink-0 ${a.st === 'EXECUTING' ? 'bg-indigo-500/20 text-indigo-300 animate-pulse' : a.st === 'COMPLETED' ? 'bg-emerald-500/20 text-emerald-400' : 'bg-slate-800 text-slate-500'}`}>{a.st}</span>
                    </div>
                  ))}
                </div>
              </section>
            </div>
          )}

          {/* ── VIEW: INCIDENT COMMAND ── */}
          {view === 'incidents' && (
            <div className="space-y-3">
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-3">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-2">
                  <h2 className="text-[11px] font-bold text-slate-200 tracking-wider">INCIDENT INVESTIGATION COMMAND CENTER</h2>
                  <div className="flex gap-2">
                    <span className="text-[9px] px-2 py-0.5 rounded bg-rose-500/20 text-rose-300 border border-rose-500/30 font-bold">SLA: 01:42:10</span>
                    <span className="text-[9px] px-2 py-0.5 rounded bg-indigo-500/20 text-indigo-300 border border-indigo-500/30 font-bold">INC-2026-0447</span>
                  </div>
                </div>

                <div className="grid grid-cols-2 gap-3">
                  {/* 5 Whys */}
                  <div className="space-y-2">
                    <h3 className="text-[10px] font-bold text-indigo-300 tracking-wider">5 WHYS ROOT CAUSE DECOMPOSITION</h3>
                    {[
                      { q: 'Why did Pump P-102 trip?', a: 'Vibration velocity reached 13.4 mm/s (ISO 10816 limit).' },
                      { q: 'Why was vibration elevated?', a: 'Bearing race misalignment due to progressive wear.' },
                      { q: 'Why did bearing wear prematurely?', a: 'Lubrication breakdown from thermal overheating.' },
                      { q: 'Why did lubrication fail?', a: 'Scheduled PM interval exceeded by 14 days.' },
                      { q: 'ROOT CAUSE:', a: 'Work order scheduling gap — no automatic escalation on PM overdue.' },
                    ].map((w, i) => (
                      <div key={i} className="p-2 rounded bg-slate-950/60 border border-slate-800/40 text-[10px]">
                        <div className={`font-bold ${i === 4 ? 'text-emerald-400' : 'text-indigo-400'}`}>{i < 4 ? `${i + 1}. ${w.q}` : w.q}</div>
                        <div className={`mt-0.5 ${i === 4 ? 'text-emerald-300 font-bold' : 'text-slate-400'}`}>{w.a}</div>
                      </div>
                    ))}
                  </div>

                  {/* CAPA + Evidence + Timeline */}
                  <div className="space-y-3">
                    <div className="space-y-2">
                      <h3 className="text-[10px] font-bold text-emerald-300 tracking-wider">CORRECTIVE ACTIONS (CAPA)</h3>
                      {[
                        { action: 'Replace bearing assembly P-102', owner: 'MECH TEAM', due: '24h', st: 'IN PROGRESS' },
                        { action: 'Recalibrate vibration probe VP-102', owner: 'I&C TEAM', due: '48h', st: 'ASSIGNED' },
                        { action: 'Implement auto-escalation on PM overdue', owner: 'CMMS ADMIN', due: '7d', st: 'PENDING' },
                        { action: 'Contractor badge audit — all Zone B', owner: 'EHS', due: '24h', st: 'COMPLETED' },
                      ].map((c, i) => (
                        <div key={i} className="p-2 rounded bg-slate-950/60 border border-slate-800/40 text-[10px] flex justify-between items-center">
                          <div>
                            <div className="text-slate-200">{c.action}</div>
                            <div className="text-[9px] text-slate-500">{c.owner} — Due: {c.due}</div>
                          </div>
                          <span className={`text-[9px] px-1.5 py-0.5 rounded font-bold ${c.st === 'COMPLETED' ? 'bg-emerald-500/20 text-emerald-400' : c.st === 'IN PROGRESS' ? 'bg-indigo-500/20 text-indigo-300' : 'bg-slate-800 text-slate-400'}`}>{c.st}</span>
                        </div>
                      ))}
                    </div>

                    <div className="space-y-2">
                      <h3 className="text-[10px] font-bold text-slate-400 tracking-wider">SENSOR CORRELATION EVIDENCE</h3>
                      <div className="p-2 rounded bg-slate-950/60 border border-slate-800/40 text-[9px] text-slate-400 space-y-1">
                        <div>• Vibration velocity exceeded 12 mm/s at 02:14:33 UTC</div>
                        <div>• Bearing temperature crossed 90°C at 02:18:11 UTC</div>
                        <div>• Correlated with lubrication oil pressure drop at 02:12:50 UTC</div>
                        <div>• Edge detection CAM-002 flagged contractor without helmet at 02:20:44 UTC</div>
                      </div>
                    </div>
                  </div>
                </div>
              </section>
            </div>
          )}

          {/* ── VIEW: INFRASTRUCTURE ── */}
          {view === 'infra' && (
            <div className="space-y-3">
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-3">
                <div className="flex justify-between items-center border-b border-slate-800/30 pb-2">
                  <h2 className="text-[11px] font-bold text-slate-200 tracking-wider">AWS INFRASTRUCTURE — ACCOUNT 727533783159 (US-EAST-1)</h2>
                  <span className={`text-[9px] px-2 py-0.5 rounded font-bold ${statusDot.replace('bg-', 'bg-') === 'bg-emerald-400' ? 'bg-emerald-500/20 text-emerald-300' : 'bg-amber-500/20 text-amber-300'}`}>BACKEND {backend.status}</span>
                </div>
                <div className="grid grid-cols-4 gap-2">
                  {[
                    { svc: 'ALB', name: 'prahari-alb-hackathon', metric: `${backend.latencyMs}ms`, st: backend.status },
                    { svc: 'ECS FARGATE', name: 'prahari-hackathon-cluster', metric: 'ai-platform:latest', st: 'RUNNING' },
                    { svc: 'RDS POSTGRES', name: 'prahari-postgres-hackathon', metric: '15.7 (db.t4g.micro)', st: backend.dbOk ? 'HEALTHY' : 'DEGRADED' },
                    { svc: 'ELASTICACHE', name: 'prahari-redis-hackathon', metric: 'cache.t4g.micro', st: backend.redisOk ? 'HEALTHY' : 'DEGRADED' },
                    { svc: 'S3 WEBSITE', name: 'prahari-hackathon-frontend', metric: 'HTTP 200 OK', st: 'SERVING' },
                    { svc: 'SECRETS MGR', name: 'prahari-hackathon-secrets', metric: 'AES-256 Encrypted', st: 'ACTIVE' },
                    { svc: 'CLOUDWATCH', name: '/ecs/prahari-hackathon', metric: 'Log Group Active', st: 'STREAMING' },
                    { svc: 'ECR', name: 'prahari-ai-platform', metric: 'linux/amd64', st: 'PUSHED' },
                  ].map(r => (
                    <div key={r.svc} className="p-2.5 rounded bg-slate-950/60 border border-slate-800/40 space-y-1">
                      <div className="flex justify-between items-center">
                        <span className="text-indigo-300 font-bold text-[10px]">{r.svc}</span>
                        <span className={`text-[9px] px-1 rounded font-bold ${['HEALTHY', 'RUNNING', 'SERVING', 'ACTIVE', 'STREAMING', 'PUSHED', 'OPERATIONAL'].includes(r.st) ? 'bg-emerald-500/20 text-emerald-400' : 'bg-amber-500/20 text-amber-400'}`}>{r.st}</span>
                      </div>
                      <div className="text-[9px] text-slate-400 truncate">{r.name}</div>
                      <div className="text-[10px] text-slate-200 font-bold">{r.metric}</div>
                    </div>
                  ))}
                </div>
              </section>

              {/* Audit Log */}
              <section className="bg-[#0c1220] border border-slate-800/40 rounded-lg p-3 space-y-2">
                <h3 className="text-[10px] font-bold text-slate-400 tracking-wider border-b border-slate-800/30 pb-1.5">REAL-TIME AUDIT TRAIL</h3>
                <div className="space-y-1 text-[10px]">
                  {[
                    { ts: backend.lastCheck, actor: 'AWS_ALB', action: 'HEALTH_CHECK', target: '/health', sev: 'INFO' },
                    { ts: backend.lastCheck, actor: 'AI_SUPERVISOR', action: 'DISPATCH_RISK_AGENT', target: 'PUMP-P102', sev: 'WARN' },
                    { ts: backend.lastCheck, actor: 'EHS-9941', action: 'LOGIN_SESSION', target: 'JWT_TOKEN', sev: 'INFO' },
                    { ts: backend.lastCheck, actor: 'EDGE_AGX04', action: 'VIOLATION_DETECT', target: 'CAM-002', sev: 'CRITICAL' },
                  ].map((l, i) => (
                    <div key={i} className="flex items-center gap-3 p-1.5 rounded bg-slate-950/40 border border-slate-800/30">
                      <span className="text-[9px] text-slate-500 w-16 shrink-0">{l.ts}</span>
                      <span className="text-indigo-300 font-bold w-28 shrink-0">{l.actor}</span>
                      <span className="text-slate-200 flex-1">{l.action}</span>
                      <span className="text-slate-400 w-24 shrink-0">{l.target}</span>
                      <span className={`text-[9px] px-1 rounded font-bold w-14 text-center shrink-0 ${l.sev === 'CRITICAL' ? 'bg-rose-500/20 text-rose-400' : l.sev === 'WARN' ? 'bg-amber-500/20 text-amber-400' : 'bg-emerald-500/20 text-emerald-400'}`}>{l.sev}</span>
                    </div>
                  ))}
                </div>
              </section>
            </div>
          )}
        </main>
      </div>

      {/* ═══ BOTTOM STATUS BAR ═══ */}
      <footer className="h-6 bg-[#0c1220] border-t border-slate-800/40 px-3 flex items-center justify-between text-[9px] text-slate-500 shrink-0">
        <div className="flex items-center gap-3">
          <span className={`flex items-center gap-1 ${statusColor}`}>
            <span className={`w-1.5 h-1.5 rounded-full ${statusDot}`}></span>
            {backend.status}
          </span>
          <span>ALB: {backend.latencyMs}ms</span>
          <span>DB: {backend.dbOk ? 'OK' : 'ERR'}</span>
          <span>REDIS: {backend.redisOk ? 'OK' : 'ERR'}</span>
          <span>LAST CHECK: {backend.lastCheck}</span>
        </div>
        <div className="flex items-center gap-3">
          <span>VIB: <strong className={latest.vib > 13 ? 'text-rose-400' : 'text-amber-400'}>{latest.vib} mm/s</strong></span>
          <span>TEMP: <strong className={latest.temp > 95 ? 'text-amber-400' : 'text-slate-300'}>{latest.temp}°C</strong></span>
          <span>PSI: <strong className="text-slate-300">{latest.psi}</strong></span>
          <span>AWS 727533783159 US-EAST-1</span>
          <span>ISO 45001 COMPLIANT</span>
        </div>
      </footer>

      {/* ═══ COMMAND PALETTE ═══ */}
      {cmdOpen && (
        <div className="fixed inset-0 z-50 bg-slate-950/80 backdrop-blur-sm flex items-start justify-center pt-16">
          <div className="w-[540px] bg-[#0c1220] border border-slate-700 rounded-xl shadow-2xl overflow-hidden">
            <div className="p-3 border-b border-slate-800">
              <input autoFocus type="text" placeholder="Search assets, dispatch agents, navigate..."
                className="w-full bg-transparent text-sm text-white placeholder-slate-500 focus:outline-none" />
            </div>
            <div className="p-2 space-y-0.5 max-h-80 overflow-y-auto text-[11px]">
              {[
                ['📊', 'Go to Operational Control Center', 'ops'],
                ['⚡', 'Inspect PUMP P-102 Digital Twin', 'twin'],
                ['📹', 'Open YOLOv8 Edge Vision Grid', 'vision'],
                ['🤖', 'View Multi-Agent Supervisor DAG', 'agents'],
                ['🚨', 'Open Incident Command Center', 'incidents'],
                ['⚙️', 'AWS Infrastructure Observability', 'infra'],
              ].map(([icon, label, target]) => (
                <button key={target} onClick={() => { setView(target as any); setCmdOpen(false); }}
                  className="w-full p-2 rounded text-left flex items-center gap-2 hover:bg-indigo-600/20 text-slate-300 hover:text-white transition-colors"
                >
                  <span>{icon}</span><span>{label}</span>
                </button>
              ))}
            </div>
            <div className="p-2 border-t border-slate-800 text-[9px] text-slate-500 text-center">
              Press <kbd className="px-1 rounded bg-slate-800 text-slate-300">ESC</kbd> to close • <kbd className="px-1 rounded bg-slate-800 text-slate-300">⌘K</kbd> to toggle
            </div>
          </div>
          <div className="fixed inset-0 -z-10" onClick={() => setCmdOpen(false)}></div>
        </div>
      )}
    </div>
  );
};

const container = document.getElementById('root');
if (container) {
  const root = createRoot(container);
  root.render(<App />);
}
