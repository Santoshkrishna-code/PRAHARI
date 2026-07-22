import React, { useState } from 'react';
import {
  LayoutDashboard, Activity, BarChart3, FileText, Radio, Box, Package, Wrench,
  Shield, AlertTriangle, ClipboardCheck, Target, Brain, Cpu, Eye, Server,
  Settings, ChevronRight, ChevronDown, Search, Bell, User, Zap, ArrowUpRight,
  ArrowDownRight, Minus, Clock, CheckCircle2, XCircle, AlertCircle, Sparkles,
  MessageSquare, Play, Circle, TrendingUp, ChevronLeft, Command, Filter,
  Download, Plus, RefreshCw, Maximize2, Layers, Thermometer, Gauge, BarChart,
  PieChart, Calendar, MapPin, Hash, ArrowRight, ExternalLink, Wifi, Database,
  Lock, Unlock, FileSearch, GitBranch, MoreHorizontal, LogOut, UserCheck,
  SlidersHorizontal, Rewind, SkipForward
} from 'lucide-react';

import { TelemetryPoint, HealthStatus, Asset, Incident } from './types';
import { UserSession } from './types/auth';
import { Toolbar, ToolBtn, ToolSep, Chart, Metric } from './components/common/CommonUI';

// ═══════════════════════════════════════════════════════════
// WORKSPACE 1: COMMAND CENTER
// ═══════════════════════════════════════════════════════════
export const CommandCenter: React.FC<{ tele: TelemetryPoint[]; onReportIncident?: () => void }> = ({ tele, onReportIncident }) => {
  const l = tele[tele.length - 1] || { vib: 11.8, temp: 94.1, psi: 242, kw: 330, flow: 84, t: '09:47' };
  return (
    <div className="h-full flex flex-col">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">COMMAND CENTER</span>
        <ToolSep />
        <span className="text-[10px] text-zinc-500">PLANT ALPHA REFINERY</span>
        <span className="text-[10px] text-zinc-600 ml-1">GULF COAST SITE</span>
        <div className="flex-1" />
        <ToolBtn onClick={onReportIncident} className="!bg-red-600/80 !text-white hover:!bg-red-500">
          <AlertTriangle size={12} /> Report Incident
        </ToolBtn>
        <ToolBtn><RefreshCw size={12} /> Refresh</ToolBtn>
        <ToolBtn><Maximize2 size={12} /></ToolBtn>
      </Toolbar>

      <div className="flex-1 overflow-y-auto p-5 space-y-5">
        {/* Hero Status Row */}
        <div className="flex gap-5">
          <div className="flex-1 flex items-center gap-5 py-4 px-5 rounded-xl bg-emerald-500/[0.06]">
            <div className="w-14 h-14 rounded-2xl bg-emerald-500/15 flex items-center justify-center">
              <CheckCircle2 size={28} className="text-emerald-400" />
            </div>
            <div>
              <p className="text-[10px] text-zinc-500 uppercase tracking-wider">Plant Operations Status</p>
              <p className="text-2xl font-bold text-emerald-400 tracking-tight">NORMAL</p>
              <p className="text-xs text-zinc-500 mt-0.5">All 1,284 telemetry signals operating within ISO safety thresholds</p>
            </div>
          </div>

          <div className="flex gap-4 items-stretch">
            {[
              { l: 'Live Signals', v: '1,284', c: 'text-white', ic: <Wifi size={14} className="text-zinc-500" /> },
              { l: 'Active Risks', v: '2', c: 'text-amber-400', ic: <AlertTriangle size={14} className="text-amber-400" /> },
              { l: 'AI Actions', v: '3', c: 'text-indigo-400', ic: <Sparkles size={14} className="text-indigo-400" /> },
              { l: 'Open Incidents', v: '0', c: 'text-emerald-400', ic: <AlertCircle size={14} className="text-emerald-400" /> },
              { l: 'Active Permits', v: '28', c: 'text-white', ic: <Shield size={14} className="text-zinc-500" /> },
            ].map(s => (
              <div key={s.l} className="w-28 flex flex-col justify-center py-3 px-3.5 rounded-xl bg-white/[0.02] border border-white/[0.04]">
                <div className="flex items-center gap-1.5 mb-1">{s.ic}<span className="text-[9px] text-zinc-500 uppercase tracking-wider">{s.l}</span></div>
                <span className={`text-xl font-bold ${s.c}`}>{s.v}</span>
              </div>
            ))}
          </div>
        </div>

        <div className="grid grid-cols-5 gap-5" style={{ gridTemplateColumns: '2fr 1fr 2fr' }}>
          {/* AI Recommendations */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2">
                <Sparkles size={14} className="text-indigo-400" />
                <span className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider">AI Recommendations</span>
              </div>
              <span className="text-[10px] text-zinc-600">3 Pending Actions</span>
            </div>
            {[
              { pri: 'high', title: 'Replace bearing assembly — Pump P-102', summary: 'Vibration trending toward ISO 10816 alarm limit. RUL: 18 days. Bearing race wear probability: 72%.', agent: 'Maintenance Agent', confidence: 94, time: '2m' },
              { pri: 'medium', title: 'Audit Zone B contractor badges', summary: 'Contractor C-4412 detected with expired badge near reactor area. Access auto-revoked.', agent: 'Contractor Agent', confidence: 92, time: '8m' },
              { pri: 'low', title: 'Schedule lubrication — Compressor C-03', summary: 'PM interval threshold reached in 6 days. Auto-escalation configured.', agent: 'Maintenance Agent', confidence: 96, time: '14m' },
            ].map((r, i) => (
              <div key={i} className="p-3.5 rounded-xl bg-white/[0.02] hover:bg-white/[0.04] border border-white/[0.04] transition-colors cursor-pointer group">
                <div className="flex items-start gap-2.5">
                  <span className={`w-1.5 h-1.5 rounded-full mt-2 shrink-0 ${r.pri === 'high' ? 'bg-amber-400' : r.pri === 'medium' ? 'bg-indigo-400' : 'bg-zinc-600'}`} />
                  <div className="flex-1 min-w-0">
                    <p className="text-[13px] font-medium text-zinc-200 group-hover:text-white transition-colors">{r.title}</p>
                    <p className="text-[12px] text-zinc-500 mt-1 leading-relaxed">{r.summary}</p>
                    <div className="flex items-center gap-3 mt-2 text-[10px] text-zinc-600">
                      <span className="flex items-center gap-1"><Brain size={10} />{r.agent}</span>
                      <span>{r.confidence}% conf.</span>
                      <span>{r.time} ago</span>
                    </div>
                  </div>
                  <ChevronRight size={14} className="text-zinc-700 group-hover:text-zinc-400 mt-1 transition-colors" />
                </div>
              </div>
            ))}
          </div>

          {/* Real-time Telemetry Gauges */}
          <div className="space-y-3">
            <span className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider flex items-center gap-2">
              <Gauge size={14} className="text-zinc-500" /> Key Indicators
            </span>
            {[
              { label: 'Vibration', value: l.vib, unit: 'mm/s', max: 16, color: l.vib > 12 ? '#f59e0b' : '#22c55e' },
              { label: 'Temperature', value: l.temp, unit: '°C', max: 120, color: l.temp > 95 ? '#f59e0b' : '#818cf8' },
              { label: 'Pressure', value: l.psi, unit: 'PSI', max: 300, color: '#22d3ee' },
              { label: 'Power Draw', value: l.kw, unit: 'kW', max: 400, color: '#a78bfa' },
              { label: 'Flow Rate', value: l.flow, unit: 'L/m', max: 110, color: '#34d399' },
            ].map(g => (
              <div key={g.label} className="p-2.5 rounded-lg bg-white/[0.02] border border-white/[0.04]">
                <div className="flex justify-between items-baseline mb-1.5">
                  <span className="text-[10px] text-zinc-500">{g.label}</span>
                  <span className="text-sm font-semibold text-white">{g.value}<span className="text-[9px] text-zinc-500 ml-0.5">{g.unit}</span></span>
                </div>
                <div className="w-full h-1 rounded-full bg-zinc-800">
                  <div className="h-1 rounded-full transition-all duration-500" style={{ width: `${(+g.value / g.max) * 100}%`, backgroundColor: g.color }} />
                </div>
              </div>
            ))}
          </div>

          {/* Live Activity Feed */}
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <span className="text-[11px] font-semibold text-zinc-400 uppercase tracking-wider flex items-center gap-2">
                <Activity size={14} className="text-zinc-500" /> Live Event Stream
              </span>
              <div className="flex items-center gap-1">
                <span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" />
                <span className="text-[10px] text-emerald-400 font-medium">Live</span>
              </div>
            </div>
            <div className="space-y-1 max-h-[340px] overflow-y-auto pr-1">
              {[
                { t: '09:47:22', src: 'AI Supervisor', msg: 'Dispatched Risk Agent for P-102 vibration anomaly', sev: 'ai' },
                { t: '09:47:18', src: 'Edge Vision', msg: 'CAM-002 PPE violation — no helmet in Zone B', sev: 'err' },
                { t: '09:47:04', src: 'Maintenance', msg: 'RUL recalculated for P-102: 18 days remaining', sev: 'warn' },
                { t: '09:46:51', src: 'Permit Agent', msg: 'PTW-8902 isolation verified on Valve V-88', sev: 'ok' },
                { t: '09:46:38', src: 'Contractor', msg: 'Badge C-4412 expired — access revoked automatically', sev: 'warn' },
                { t: '09:46:22', src: 'Telemetry', msg: 'Vibration probe VP-102 reading 11.8 mm/s', sev: 'warn' },
                { t: '09:46:11', src: 'Edge Vision', msg: 'CAM-001 PPE scan complete — 3 workers verified', sev: 'ok' },
                { t: '09:45:58', src: 'Telemetry', msg: 'Temperature TS-102 nominal at 88°C', sev: 'ok' },
              ].map((e, i) => (
                <div key={i} className="flex items-start gap-2 py-2 px-2 rounded-lg hover:bg-white/[0.02] transition-colors text-[12px] cursor-pointer group">
                  <span className={`w-1.5 h-1.5 rounded-full mt-1.5 shrink-0 ${e.sev === 'err' ? 'bg-red-400' : e.sev === 'warn' ? 'bg-amber-400' : e.sev === 'ai' ? 'bg-indigo-400' : 'bg-emerald-400/60'}`} />
                  <div className="flex-1 min-w-0">
                    <span className="text-zinc-400 group-hover:text-zinc-200 transition-colors">{e.msg}</span>
                    <div className="flex gap-2 mt-0.5 text-[10px] text-zinc-600">
                      <span>{e.src}</span><span>{e.t}</span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 2: OPERATIONS CENTER
// ═══════════════════════════════════════════════════════════
export const OperationsCenter: React.FC<{ tele: TelemetryPoint[] }> = ({ tele }) => {
  const l = tele[tele.length - 1] || { vib: 11.8, temp: 94.1, psi: 242, kw: 330, flow: 84 };
  const [chartMode, setChartMode] = useState<'vib' | 'temp' | 'psi' | 'kw'>('vib');
  const chartConfig = {
    vib: { data: tele.map(p => p.vib), color: '#818cf8', threshold: 14, label: 'Vibration Velocity', unit: 'mm/s', alarm: 'ISO 10816 Limit @ 14.0' },
    temp: { data: tele.map(p => p.temp), color: '#f59e0b', threshold: 100, label: 'Bearing Temperature', unit: '°C', alarm: 'Alarm @ 100°C' },
    psi: { data: tele.map(p => p.psi), color: '#22d3ee', threshold: 260, label: 'Line Pressure', unit: 'PSI', alarm: 'High @ 260 PSI' },
    kw: { data: tele.map(p => +p.kw), color: '#a78bfa', threshold: 370, label: 'Power Draw', unit: 'kW', alarm: 'Peak @ 370 kW' },
  }[chartMode];

  return (
    <div className="h-full flex flex-col">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">OPERATIONS CENTER</span>
        <ToolSep />
        <ToolBtn active={chartMode === 'vib'} onClick={() => setChartMode('vib')}>Vibration</ToolBtn>
        <ToolBtn active={chartMode === 'temp'} onClick={() => setChartMode('temp')}>Temperature</ToolBtn>
        <ToolBtn active={chartMode === 'psi'} onClick={() => setChartMode('psi')}>Pressure</ToolBtn>
        <ToolBtn active={chartMode === 'kw'} onClick={() => setChartMode('kw')}>Power</ToolBtn>
        <div className="flex-1" />
        <span className="text-[10px] text-emerald-400 font-medium flex items-center gap-1"><span className="w-1.5 h-1.5 rounded-full bg-emerald-400 animate-pulse" /> Live Telemetry</span>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 flex flex-col overflow-hidden">
          <div className="h-16 flex items-center gap-8 px-5 border-b border-white/[0.04] bg-white/[0.01]">
            <Metric label="Vibration" value={l.vib} unit="mm/s" accent={l.vib > 12 ? 'text-amber-400' : undefined} small />
            <Metric label="Temperature" value={l.temp} unit="°C" accent={l.temp > 95 ? 'text-amber-400' : undefined} small />
            <Metric label="Pressure" value={l.psi} unit="PSI" small />
            <Metric label="Power" value={l.kw} unit="kW" small />
            <Metric label="Flow" value={l.flow} unit="L/m" small />
          </div>

          <div className="flex-1 p-4 flex flex-col">
            <div className="flex items-baseline justify-between mb-2">
              <div>
                <span className="text-sm font-medium text-white">{chartConfig.label} — Pump P-102</span>
                <span className="text-xs text-zinc-600 ml-2">{chartConfig.alarm}</span>
              </div>
              <span className="text-lg font-semibold text-white">{chartConfig.data[chartConfig.data.length - 1]} <span className="text-xs text-zinc-500 font-normal">{chartConfig.unit}</span></span>
            </div>
            <div className="flex-1 rounded-xl bg-white/[0.015] border border-white/[0.04] p-3">
              <Chart data={chartConfig.data} color={chartConfig.color} threshold={chartConfig.threshold} h={220} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 3: INDUSTRIAL TWIN
// ═══════════════════════════════════════════════════════════
export const IndustrialTwin: React.FC<{ tele: TelemetryPoint[] }> = ({ tele }) => {
  const l = tele[tele.length - 1] || { temp: 94.1, vib: 11.8 };
  const [selected, setSelected] = useState('pump');
  const assets = [
    { id: 'reactor', label: 'Reactor B', x: 180, y: 130, health: 97, temp: '340°C', st: 'Running', rul: '94d' },
    { id: 'pump', label: 'Pump P-102', x: 400, y: 200, health: 74, temp: `${l.temp}°C`, st: 'Warning', rul: '18d' },
    { id: 'hx', label: 'HX-04', x: 600, y: 120, health: 91, temp: '88°C', st: 'Running', rul: '62d' },
    { id: 'comp', label: 'Compressor C-03', x: 360, y: 340, health: 96, temp: '72°C', st: 'Running', rul: '84d' },
    { id: 'valve', label: 'Valve V-88', x: 140, y: 310, health: 98, temp: 'N/A', st: 'Locked', rul: 'N/A' },
    { id: 'boiler', label: 'Boiler A', x: 620, y: 310, health: 93, temp: '210°C', st: 'Running', rul: '71d' },
  ];
  const sel = assets.find(a => a.id === selected);

  return (
    <div className="h-full flex flex-col">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">INDUSTRIAL TWIN</span>
        <ToolSep />
        <span className="text-[10px] text-zinc-500">Reactor Complex B Schematic</span>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 relative bg-[#0b0b0f]">
          <svg viewBox="0 0 800 460" className="w-full h-full" preserveAspectRatio="xMidYMid meet">
            {[[240,150,380,200],[460,200,580,140],[400,240,380,320],[180,190,160,290],[430,340,580,320]].map(([x1,y1,x2,y2], i) => (
              <line key={i} x1={x1} y1={y1} x2={x2} y2={y2} stroke="rgba(99,102,241,0.1)" strokeWidth="2" strokeDasharray="8,6" />
            ))}
            {assets.map(a => {
              const s = selected === a.id;
              const w = a.st === 'Warning';
              return (
                <g key={a.id} onClick={() => setSelected(a.id)} className="cursor-pointer">
                  <circle cx={a.x} cy={a.y} r={s ? 40 : 34} fill={s ? 'rgba(99,102,241,0.15)' : w ? 'rgba(245,158,11,0.08)' : 'rgba(255,255,255,0.02)'} stroke={s ? '#6366f1' : w ? '#f59e0b' : 'rgba(255,255,255,0.08)'} strokeWidth={s ? 2 : 1} />
                  <circle cx={a.x} cy={a.y} r={5} fill={a.health > 90 ? '#22c55e' : a.health > 80 ? '#f59e0b' : '#ef4444'} />
                  <text x={a.x} y={a.y + 54} textAnchor="middle" fill={s ? '#c7d2fe' : 'rgba(255,255,255,0.4)'} fontSize="11" fontFamily="Inter" fontWeight="500">{a.label}</text>
                </g>
              );
            })}
          </svg>
        </div>

        {sel && (
          <div className="w-72 border-l border-white/[0.04] p-4 bg-white/[0.01] space-y-4">
            <div>
              <p className="text-[10px] text-zinc-500 uppercase">Selected Equipment</p>
              <h2 className="text-base font-bold text-white mt-0.5">{sel.label}</h2>
              <span className={`inline-block mt-1 text-[10px] px-2 py-0.5 rounded-full ${sel.st === 'Warning' ? 'bg-amber-500/15 text-amber-400' : 'bg-emerald-500/10 text-emerald-400'}`}>{sel.st}</span>
            </div>
            <div className="grid grid-cols-2 gap-3">
              <Metric label="Health" value={`${sel.health}%`} small />
              <Metric label="RUL" value={sel.rul} small />
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 4: AI COMMAND CENTER
// ═══════════════════════════════════════════════════════════
export const AICommandCenter: React.FC = () => {
  const steps = [
    { text: 'Detected vibration anomaly on Pump P-102', status: 'done', detail: 'Reading 11.8 mm/s, trending toward ISO limit.' },
    { text: 'Queried CMMS maintenance history', status: 'done', detail: 'PM overdue by 14 days. Work order WO-7821 un-escalated.' },
    { text: 'Calculated bearing wear probability', status: 'done', detail: 'PINN digital twin predicts 72% wear probability. RUL: 18 days.' },
    { text: 'Verified PTW-8902 isolation lock', status: 'done', detail: 'Valve V-88 isolation locked. No conflicts.' },
    { text: 'Synthesizing corrective action plan (CAPA)', status: 'active', detail: '4 actions generated for maintenance crew.' },
  ];

  return (
    <div className="h-full flex flex-col">
      <Toolbar>
        <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">AI COMMAND CENTER</span>
        <ToolSep />
        <span className="text-[10px] text-indigo-400 font-medium">10 Autonomous Agents Active</span>
      </Toolbar>

      <div className="flex-1 flex overflow-hidden">
        <div className="flex-1 p-5 space-y-4 overflow-y-auto">
          <div className="p-4 rounded-xl bg-indigo-600/10 border border-indigo-500/20">
            <h2 className="text-sm font-bold text-white">Active Reasoning: Pump P-102 Anomaly</h2>
            <p className="text-xs text-zinc-400 mt-1">Multi-agent supervisor is conducting 5-Whys RCA and evidence correlation.</p>
          </div>
          <div className="space-y-2">
            {steps.map((s, i) => (
              <div key={i} className="flex gap-3 py-2 px-3 rounded-lg bg-white/[0.02]">
                <CheckCircle2 size={16} className={s.status === 'done' ? 'text-emerald-400' : 'text-indigo-400'} />
                <div>
                  <p className="text-[12px] font-medium text-zinc-200">{s.text}</p>
                  <p className="text-[11px] text-zinc-500 mt-0.5">{s.detail}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

// ═══════════════════════════════════════════════════════════
// WORKSPACE 5: INCIDENTS WORKSPACE
// ═══════════════════════════════════════════════════════════
export const IncidentsWorkspace: React.FC<{ onReportIncident?: () => void }> = ({ onReportIncident }) => (
  <div className="h-full flex flex-col">
    <Toolbar>
      <span className="text-[11px] font-semibold text-zinc-300 tracking-wider">INCIDENTS & HAZARDS</span>
      <div className="flex-1" />
      <ToolBtn onClick={onReportIncident} className="!bg-red-600/80 !text-white">
        <Plus size={12} /> Report Incident
      </ToolBtn>
    </Toolbar>
    <div className="flex-1 p-5 space-y-4 overflow-y-auto">
      <div className="p-4 rounded-xl bg-white/[0.02] border border-white/[0.04]">
        <div className="flex justify-between">
          <span className="text-xs font-mono text-zinc-500">INC-2026-0447</span>
          <span className="text-xs text-amber-400 font-medium">Under Investigation</span>
        </div>
        <h2 className="text-sm font-bold text-white mt-1">Pump P-102 Vibration Excursion</h2>
        <p className="text-xs text-zinc-400 mt-1">5-Whys Root Cause: PM overdue by 14 days led to bearing lubrication failure.</p>
      </div>
    </div>
  </div>
);

// ═══════════════════════════════════════════════════════════
// OTHER WORKSPACES (EXPORTS)
// ═══════════════════════════════════════════════════════════
export const VisionIntelligence: React.FC = () => (
  <div className="h-full flex flex-col">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">VISION INTELLIGENCE</span></Toolbar>
    <div className="p-5 text-xs text-zinc-400">4-Camera YOLOv8 Inference Grid Operational</div>
  </div>
);

export const AgentOrchestration: React.FC = () => (
  <div className="h-full flex flex-col">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">AGENT ORCHESTRATION</span></Toolbar>
    <div className="p-5 text-xs text-zinc-400">10-Agent Execution Trace & Memory DAG</div>
  </div>
);

export const AssetsWorkspace: React.FC<{ tele: TelemetryPoint[]; onAddAsset?: () => void }> = ({ onAddAsset }) => (
  <div className="h-full flex flex-col">
    <Toolbar>
      <span className="text-[11px] font-semibold text-zinc-300">ASSETS WORKSPACE</span>
      <div className="flex-1" />
      <ToolBtn onClick={onAddAsset} className="!bg-indigo-600 !text-white"><Plus size={12} /> Add Asset</ToolBtn>
    </Toolbar>
    <div className="p-5 text-xs text-zinc-400">IBM Maximo Asset Management Grid (47 Monitored Assets)</div>
  </div>
);

export const PlatformOps: React.FC<{ health: HealthStatus }> = ({ health }) => (
  <div className="h-full flex flex-col">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">PLATFORM OPERATIONS</span></Toolbar>
    <div className="p-5 text-xs text-zinc-400">AWS Infrastructure Observability: {health.status} ({health.lat}ms)</div>
  </div>
);

export const OpsIntelligence: React.FC<{ tele: TelemetryPoint[] }> = () => (
  <div className="h-full flex flex-col"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">OPERATIONS INTELLIGENCE</span></Toolbar><div className="p-5 text-xs text-zinc-400">Safety Index: 94.2/100</div></div>
);

export const PermitsWorkspace: React.FC = () => (
  <div className="h-full flex flex-col"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">SAFE WORK PERMITS</span></Toolbar><div className="p-5 text-xs text-zinc-400">28 Active Permits to Work</div></div>
);

export const MaintenanceWorkspace: React.FC = () => (
  <div className="h-full flex flex-col"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">MAINTENANCE WORKFLOW</span></Toolbar><div className="p-5 text-xs text-zinc-400">Predictive Work Order Queue</div></div>
);

export const RiskWorkspace: React.FC = () => (
  <div className="h-full flex flex-col"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">RISK ASSESSMENT</span></Toolbar><div className="p-5 text-xs text-zinc-400">5×5 Industrial Risk Matrix</div></div>
);

export const InspectionsWorkspace: React.FC = () => (
  <div className="h-full flex flex-col"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">INSPECTIONS</span></Toolbar><div className="p-5 text-xs text-zinc-400">Compliance Audit Checklists</div></div>
);

export const ExecutiveInsights: React.FC = () => (
  <div className="h-full flex flex-col"><Toolbar><span className="text-[11px] font-semibold text-zinc-300">EXECUTIVE INSIGHTS</span></Toolbar><div className="p-5 text-xs text-zinc-400">ISO 45001 & OSHA Audit Reports</div></div>
);

export const SettingsWorkspace: React.FC<{ session?: UserSession }> = ({ session }) => (
  <div className="h-full flex flex-col">
    <Toolbar><span className="text-[11px] font-semibold text-zinc-300">SETTINGS</span></Toolbar>
    <div className="p-5 text-xs text-zinc-400">Organization: {session?.orgName || 'Alpha Chemical Refinery Inc.'} ({session?.role})</div>
  </div>
);
